package rbac

import "testing"

func TestAllowed(t *testing.T) { e := New(map[string][]string{"admin": {"write"}}); if !e.Allowed([]string{"admin"}, "write") { t.Fatal() } }
