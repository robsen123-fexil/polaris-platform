package rbac

type Evaluator struct{ m map[string]map[string]struct{} }
func New(rolePerms map[string][]string) *Evaluator { m := map[string]map[string]struct{}{}; for role, perms := range rolePerms { s := map[string]struct{}{}; for _, p := range perms { s[p] = struct{}{} }; m[role] = s }; return &Evaluator{m: m} }
func (e *Evaluator) Allowed(roles []string, perm string) bool { for _, r := range roles { if perms, ok := e.m[r]; ok { if _, ok := perms[perm]; ok { return true } } }; return false }
