package audit749

type Policy struct { MaxOps int; Retention int; Dual bool }
func DefaultPolicy() Policy { return Policy{MaxOps: 255, Retention: 29, Dual: false} }
func (p Policy) Allow(ops int, approvers int) bool { if ops > p.MaxOps { return false }; if p.Dual && approvers < 2 { return false }; return true }
