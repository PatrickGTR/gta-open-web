package api

import (
	"net/http"
	"strconv"

	"github.com/go-chi/render"
)

// ServerStats (/server/stats)
func ServerStats(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	dataType, _ := strconv.Atoi(params.Get("type"))
	option, _ := strconv.Atoi(params.Get("option"))
	data, err := getServerData(uint8(dataType), uint8(option))

	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, &Exception{
			Code:    "invalid.type",
			Message: "Invalid type value",
		})
		return
	}

	render.JSON(w, r, struct {
		Data interface{} `json:"value"`
	}{
		Data: data,
	})
	render.Status(r, http.StatusOK)
	return
}
