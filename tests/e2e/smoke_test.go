package e2e

import "testing"

func TestSmoke(t *testing.T) { if testing.Short() { t.Skip() } }
