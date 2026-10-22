package server

import (
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

type rateLimiter struct {
	mu       sync.Mutex
	visitors map[string]*visitor
	rate     int
	burst    int
	interval time.Duration
	// ipHeader is the header that carries the real client IP, or "" when no
	// header is trusted and the socket address is used. See
	// resolveTrustedIPHeader.
	ipHeader string
}

type visitor struct {
	tokens    int
	lastCheck time.Time
}

func newRateLimiter(rate, burst int, interval time.Duration, ipHeader string) *rateLimiter {
	rl := &rateLimiter{
		visitors: make(map[string]*visitor),
		rate:     rate,
		burst:    burst,
		interval: interval,
		ipHeader: ipHeader,
	}
	go rl.cleanup()
	return rl
}

func (rl *rateLimiter) cleanup() {
	ticker := time.NewTicker(10 * time.Minute)
	for range ticker.C {
		rl.mu.Lock()
		cutoff := time.Now().Add(-30 * time.Minute)
		for ip, v := range rl.visitors {
			if v.lastCheck.Before(cutoff) {
				delete(rl.visitors, ip)
			}
		}
		rl.mu.Unlock()
	}
}

func (rl *rateLimiter) allow(ip string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	v, ok := rl.visitors[ip]
	now := time.Now()

	if !ok {
		rl.visitors[ip] = &visitor{tokens: rl.burst - 1, lastCheck: now}
		return true
	}

	elapsed := now.Sub(v.lastCheck)
	v.lastCheck = now
	v.tokens += int(elapsed / rl.interval)
	if v.tokens > rl.burst {
		v.tokens = rl.burst
	}
	if v.tokens <= 0 {
		return false
	}
	v.tokens--
	return true
}

func (rl *rateLimiter) middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !rl.allow(clientIP(r, rl.ipHeader)) {
			http.Error(w, `{"error":"too many attempts — try again in a minute"}`, 429)
			return
		}
		next(w, r)
	}
}

// resolveTrustedIPHeader decides which request header, if any, carries the true
// client IP.
//
// A header is only trustworthy when the edge that sets it also overwrites any
// value the client supplied. Otherwise it is just another spoofable input, and
// both the rate limiter and the guest AI quota are bypassed by setting it: a
// caller sends `Fly-Client-IP: 1.2.3.4` and is limited under a key of their own
// choosing, or, worse, spends somebody else's AI budget.
//
// Fly's edge sets Fly-Client-IP and overwrites it, and Fly puts FLY_APP_NAME in
// the machine environment, so that header is trusted exactly when FLY_APP_NAME
// is present. A deployment behind a different proxy sets AI_TRUSTED_IP_HEADER to
// the header that proxy controls. With neither, the socket address is the only
// untamperable signal, which is the correct answer for a directly exposed
// process.
func resolveTrustedIPHeader() string {
	if h := strings.TrimSpace(os.Getenv("AI_TRUSTED_IP_HEADER")); h != "" {
		return h
	}
	if os.Getenv("FLY_APP_NAME") != "" {
		return "Fly-Client-IP"
	}
	return ""
}

// clientIP returns the caller's address for rate limiting and guest AI quotas.
//
// It reads trustedHeader only, never X-Forwarded-For: that header's first entry
// is client-controlled, so a parse that trusted it was a rate-limit bypass for
// anyone willing to set a header. When no header is trusted, or the trusted one
// is absent, the socket address is used, so a request that never traversed the
// edge is keyed by the address that actually connected.
func clientIP(r *http.Request, trustedHeader string) string {
	if trustedHeader != "" {
		if ip := strings.TrimSpace(r.Header.Get(trustedHeader)); ip != "" {
			return ip
		}
	}
	if host, _, err := net.SplitHostPort(r.RemoteAddr); err == nil {
		return host
	}
	return r.RemoteAddr
}
