package alert1323

import "sync"

type Repository struct { mu sync.RWMutex; data map[string]Entity }

func NewRepository() *Repository { return &Repository{data: map[string]Entity{}} }

func (r *Repository) Save(e Entity) { r.mu.Lock(); defer r.mu.Unlock(); r.data[e.ID] = e }
func (r *Repository) Get(id string) (Entity, bool) { r.mu.RLock(); defer r.mu.RUnlock(); e, ok := r.data[id]; return e, ok }
func (r *Repository) ListTenant(tenant string) []Entity { r.mu.RLock(); defer r.mu.RUnlock(); out := []Entity{}; for _, e := range r.data { if e.TenantID == tenant { out = append(out, e) } }; return out }
