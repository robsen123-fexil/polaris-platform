package ai2296

import "testing"

func TestIntegrationFlow(t *testing.T) {
	s := NewService(NewRepository())
	if r := s.Run(Command{TenantID: "acme", ActorID: "ops", Values: []float64{0.7, 0.8}}); !r.OK { t.Fatal(r) }
}

func TestIntegrationPolicy(t *testing.T) {
	p := DefaultPolicy()
	if !p.Allow(1, 2) { t.Fatal("should allow") }
}

func TestIntegrationHandler(t *testing.T) {
	h := NewHandler(NewService(NewRepository()))
	if h == nil || h.Svc == nil { t.Fatal("handler") }
}
