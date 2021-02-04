package routes

import (
	"net/http"
	"strconv"

	"github.com/go-chi/render"
	"github.com/open-backend/server"
)

func ServerStats(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	option, _ := strconv.Atoi(params.Get("type"))

	if option == 0 {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, &Exception{
			Code:    "type.param.empty",
			Message: "Invalid param type",
		})
		return
	}

	username, err := server.GetHighest(uint8(option))

	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, &Exception{
			Code:    "invalid.type",
			Message: "Invalid type value",
		})
		return
	}

	render.JSON(w, r, struct {
		Username string `json:"username"`
	}{
		Username: username,
	})
	render.Status(r, http.StatusOK)
	return
}
