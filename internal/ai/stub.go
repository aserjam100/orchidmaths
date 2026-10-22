package ai

import (
	"context"
	"sync"
)

// Stub is a deterministic in-process Provider for tests. It opens no socket, so
// the suite exercises the registry, redaction, quota and failover logic without
// spending a token or depending on the network being reachable. Production code
// never constructs one.
type Stub struct {
	// IDValue is returned by ID.
	IDValue string
	// Pol is returned by Policy.
	Pol Policy
	// Caps is returned by Capabilities.
	Caps Capabilities
	// Reply is returned verbatim by Chat when Err is nil.
	Reply string
	// Err, when set, is returned by Chat instead of Reply.
	Err error

	mu      sync.Mutex
	calls   int
	lastReq Request
}

func (s *Stub) ID() string                 { return s.IDValue }
func (s *Stub) Policy() Policy             { return s.Pol }
func (s *Stub) Capabilities() Capabilities { return s.Caps }

func (s *Stub) Chat(_ context.Context, req Request) (Response, error) {
	s.mu.Lock()
	s.calls++
	s.lastReq = req
	s.mu.Unlock()
	if s.Err != nil {
		return Response{}, s.Err
	}
	return Response{Text: s.Reply, Model: req.Model}, nil
}

// Calls is how many times Chat has been entered.
func (s *Stub) Calls() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.calls
}

// LastRequest is the most recent Request passed to Chat.
func (s *Stub) LastRequest() Request {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.lastReq
}
