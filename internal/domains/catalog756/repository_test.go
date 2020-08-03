package catalog756

import "testing"

func TestRepositorySaveGet(t *testing.T) {
	r := NewRepository()
	e := Entity{ID: "1", TenantID: "t", Domain: "catalog756", Score: 0.5}
	r.Save(e)
	if _, ok := r.Get("1"); !ok { t.Fatal("missing") }
}

func TestRepositoryListTenant(t *testing.T) {
	r := NewRepository()
	r.Save(Entity{ID: "1", TenantID: "t", Score: 0.5})
	r.Save(Entity{ID: "2", TenantID: "t", Score: 0.6})
	if len(r.ListTenant("t")) != 2 { t.Fatal("expected 2") }
}

func TestRepositoryIsolation(t *testing.T) {
	r := NewRepository()
	r.Save(Entity{ID: "1", TenantID: "a", Score: 0.5})
	if len(r.ListTenant("b")) != 0 { t.Fatal("leak") }
}
