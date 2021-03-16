package server

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type ServerService struct {
	Routes chi.Router
	db     *sql.DB
}

func New(db *sql.DB) *ServerService {

	router := chi.NewRouter()

	service := &ServerService{
		Routes: router,
		db:     db,
	}

	router.Get("/stats", service.ServerStats)
	router.Get("/banlist", service.BanList)

	return service
}

// ServerStats (/server/stats)
func (s *ServerService) ServerStats(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	dataType, _ := strconv.Atoi(params.Get("type"))
	option, _ := strconv.Atoi(params.Get("option"))
	data, err := s.getServerData(uint8(dataType), uint8(option))

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

func (s *ServerService) BanList(w http.ResponseWriter, r *http.Request) {

	banList, err := s.getAllBans()

	if err != nil {
		// an error occured
		return
	}

	render.JSON(w, r, banList)
	render.Status(r, http.StatusOK)
	return
}
