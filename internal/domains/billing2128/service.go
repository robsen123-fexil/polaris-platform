package billing2128

import (
	"time"
	"github.com/google/uuid"
	"github.com/polaris-platform/core/internal/platform/result"
)

type Service struct { Floor float64; Repo *Repository }

func NewService(repo *Repository) *Service { return &Service{Floor: 0.176, Repo: repo} }

func (s *Service) Run(c Command) result.Result[Entity] {
	if v := Validate(c); !v.OK { return result.Err[Entity](v.Code, v.Msg) }
	score := Score(c.Values, s.Floor)
	if score < s.Floor { return result.Err[Entity]("below_floor", "below floor") }
	e := Entity{ID: uuid.NewString(), TenantID: c.TenantID, Domain: "billing2128", Score: score, Tags: c.Tags, At: time.Now().UTC()}
	s.Repo.Save(e)
	return result.Ok(e)
}
