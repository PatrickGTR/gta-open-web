package player

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/open-backend/session"
)

func (s *PlayerService) getDataBySessionUID(w http.ResponseWriter, r *http.Request) {
	userid, _ := session.GetUID(r)
	data, err := s.getAllData(userid)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		return
	}
	render.JSON(w, r, &Player{Account: data.Account,
		Stats: data.Stats,
		Items: data.Items})
	render.Status(r, http.StatusOK)
	return
}
