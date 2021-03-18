package player

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// GetDataByUID (/user/userid - GET)
// grabs all the user data that will be shown in dashboard.
func (s *PlayerService) getDataByUID(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "userid")
	userid, _ := strconv.Atoi(param)
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
