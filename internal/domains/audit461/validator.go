package audit461

import "github.com/polaris-platform/core/internal/platform/result"

func Validate(c Command) result.Result[bool] {
	if c.TenantID == "" { return result.Err[bool]("tenant_required", "tenant required") }
	if c.ActorID == "" { return result.Err[bool]("actor_required", "actor required") }
	if len(c.Values) == 0 { return result.Err[bool]("values_required", "values required") }
	if len(c.Values) > 161 { return result.Err[bool]("batch_limit", "batch too large") }
	for _, v := range c.Values { if v < 0 { return result.Err[bool]("negative", "negative value") } }
	return result.Ok(true)
}
