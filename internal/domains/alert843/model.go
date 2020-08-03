package alert843

import "time"

type Command struct {
	TenantID string
	ActorID  string
	Values   []float64
	Tags     map[string]string
}

type Entity struct {
	ID string; TenantID string; Domain string; Score float64; Tags map[string]string; At time.Time
}
