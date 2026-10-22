package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

// TestRateLimiter_IgnoresSpoofableXForwardedFor pins the fix for a rate-limit
// bypass: the limiter used to key on the first X-Forwarded-For entry, which the
// client controls. Two requests from the same socket that carry different
// forged X-Forwarded-For values must share one bucket, and the bucket must be
// the socket address, not the forged value.
func TestRateLimiter_IgnoresSpoofableXForwardedFor(t *testing.T) {
	rl := newRateLimiter(100, 100, time.Minute, "Fly-Client-IP")
	h := rl.middleware(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
	for _, forged := range []string{"203.0.113.7, 70.41.3.18", "198.51.100.9", ""} {
		req := httptest.NewRequest("GET", "/", nil)
		req.RemoteAddr = "10.0.0.1:1234"
		if forged != "" {
			req.Header.Set("X-Forwarded-For", forged)
		}
		rec := httptest.NewRecorder()
		h(rec, req)
		if rec.Code != 200 {
			t.Fatalf("expected pass-through, got %d", rec.Code)
		}
	}
	if len(rl.visitors) != 1 {
		t.Fatalf("forged X-Forwarded-For sharded the limiter: %d buckets, want 1", len(rl.visitors))
	}
	for ip := range rl.visitors {
		if ip != "10.0.0.1" {
			t.Errorf("bucket = %q, want the socket address 10.0.0.1", ip)
		}
	}
}

// TestRateLimiter_PrefersFlyClientIP checks the trusted header: behind Fly's
// edge every request arrives from the proxy, so the real client is read from
// Fly-Client-IP, which the edge sets and overwrites.
func TestRateLimiter_PrefersFlyClientIP(t *testing.T) {
	rl := newRateLimiter(100, 100, time.Minute, "Fly-Client-IP")
	h := rl.middleware(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
	req := httptest.NewRequest("GET", "/", nil)
	req.RemoteAddr = "172.16.0.5:443"
	req.Header.Set("Fly-Client-IP", "203.0.113.7")
	req.Header.Set("X-Forwarded-For", "198.51.100.9")
	rec := httptest.NewRecorder()
	h(rec, req)
	if rec.Code != 200 {
		t.Fatalf("expected pass-through, got %d", rec.Code)
	}
	for ip := range rl.visitors {
		if ip != "203.0.113.7" {
			t.Errorf("bucket = %q, want Fly-Client-IP 203.0.113.7", ip)
		}
	}
}

// TestRateLimiter_UntrustedHeaderIsIgnored is the deployment-shape guard: the
// same request, but with no trusted header configured, so a client-supplied
// Fly-Client-IP must not shard the limiter. A process that is not behind Fly
// trusts that header only if it is told to.
func TestRateLimiter_UntrustedHeaderIsIgnored(t *testing.T) {
	rl := newRateLimiter(100, 100, time.Minute, "")
	h := rl.middleware(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
	for _, forged := range []string{"203.0.113.7", "203.0.113.8"} {
		req := httptest.NewRequest("GET", "/", nil)
		req.RemoteAddr = "10.0.0.1:1234"
		req.Header.Set("Fly-Client-IP", forged)
		rec := httptest.NewRecorder()
		h(rec, req)
		if rec.Code != 200 {
			t.Fatalf("expected pass-through, got %d", rec.Code)
		}
	}
	if len(rl.visitors) != 1 {
		t.Fatalf("an untrusted header sharded the limiter: %d buckets, want 1", len(rl.visitors))
	}
	for ip := range rl.visitors {
		if ip != "10.0.0.1" {
			t.Errorf("bucket = %q, want the socket address 10.0.0.1", ip)
		}
	}
}

// TestResolveTrustedIPHeader pins the rule: an explicit override wins, Fly's
// header is trusted only when Fly set the deployment's environment, and a bare
// process on the internet trusts nothing.
func TestResolveTrustedIPHeader(t *testing.T) {
	t.Setenv("AI_TRUSTED_IP_HEADER", "")
	t.Setenv("FLY_APP_NAME", "")
	if got := resolveTrustedIPHeader(); got != "" {
		t.Errorf("bare process: header = %q, want none", got)
	}
	t.Setenv("FLY_APP_NAME", "orchid")
	if got := resolveTrustedIPHeader(); got != "Fly-Client-IP" {
		t.Errorf("on Fly: header = %q, want Fly-Client-IP", got)
	}
	t.Setenv("AI_TRUSTED_IP_HEADER", "CF-Connecting-IP")
	if got := resolveTrustedIPHeader(); got != "CF-Connecting-IP" {
		t.Errorf("override: header = %q, want CF-Connecting-IP", got)
	}
}

func TestClientIPFallsBackToRemoteAddr(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	req.RemoteAddr = "192.0.2.10:5000"
	if got := clientIP(req, "Fly-Client-IP"); got != "192.0.2.10" {
		t.Errorf("clientIP = %q, want 192.0.2.10", got)
	}
	// A RemoteAddr without a port is returned unchanged rather than emptied.
	req.RemoteAddr = "192.0.2.10"
	if got := clientIP(req, "Fly-Client-IP"); got != "192.0.2.10" {
		t.Errorf("clientIP = %q, want 192.0.2.10", got)
	}
}

func TestRateLimiter_Friendly429(t *testing.T) {
	rl := newRateLimiter(100, 1, time.Hour, "Fly-Client-IP")
	h := rl.middleware(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
	mkReq := func() *http.Request {
		req := httptest.NewRequest("GET", "/", nil)
		req.RemoteAddr = "10.9.9.9:1"
		return req
	}
	rec := httptest.NewRecorder()
	h(rec, mkReq())
	if rec.Code != 200 {
		t.Fatalf("first request should pass, got %d", rec.Code)
	}
	rec = httptest.NewRecorder()
	h(rec, mkReq())
	if rec.Code != 429 {
		t.Fatalf("expected 429, got %d", rec.Code)
	}
	body := rec.Body.String()
	if body == "" || strings.Contains(body, "rate limit exceeded") {
		t.Errorf("expected friendly 429 copy, got %q", body)
	}
}
