package auth1536

import "testing"

func TestValidateOk(t *testing.T) {
	if v := Validate(Command{TenantID: "t", ActorID: "a", Values: []float64{0.5}}); !v.OK { t.Fatal(v) }
}

func TestValidateBatchLimit(t *testing.T) {
	vals := make([]float64, 166)
	for i := range vals { vals[i] = 0.5 }
	if v := Validate(Command{TenantID: "t", ActorID: "a", Values: vals}); v.OK { t.Fatal("expected batch limit") }
}
