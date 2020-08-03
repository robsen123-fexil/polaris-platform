package alert219

import "time"

type Event struct { ID, Kind, TenantID, RecordID string; At time.Time }
func NewEvent(kind, tenant, record string) Event { return Event{ID: record+"-"+kind, Kind: kind, TenantID: tenant, RecordID: record, At: time.Now().UTC()} }
