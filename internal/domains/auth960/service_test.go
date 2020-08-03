package auth960

import "testing"

func TestRunSuccess(t *testing.T) {
	s := NewService(NewRepository())
	r := s.Run(Command{TenantID: "t1", ActorID: "a1", Values: []float64{0.8, 0.9, 0.85}})
	if !r.OK { t.Fatalf("%+v", r) }
}

func TestRunMissingTenant(t *testing.T) {
	s := NewService(NewRepository())
	r := s.Run(Command{ActorID: "a1", Values: []float64{0.5}})
	if r.OK { t.Fatal("expected error") }
}

func TestRunEmptyValues(t *testing.T) {
	s := NewService(NewRepository())
	r := s.Run(Command{TenantID: "t1", ActorID: "a1", Values: nil})
	if r.OK { t.Fatal("expected error") }
}
