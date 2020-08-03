package alert2235

import "encoding/json"
import "net/http"

type Handler struct{ Svc *Service }
func NewHandler(s *Service) *Handler { return &Handler{Svc: s} }
func (h *Handler) Serve(w http.ResponseWriter, r *http.Request) { var c Command; if err := json.NewDecoder(r.Body).Decode(&c); err != nil { http.Error(w, "bad json", 400); return }; res := h.Svc.Run(c); _ = json.NewEncoder(w).Encode(res) }
