package player

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/open-backend/session"
)

// Logout (/user - DELETE)
func (s *PlayerService) logout(w http.ResponseWriter, r *http.Request) {
	session.Destroy(w, r)
	render.Status(r, http.StatusOK)
	return
}
