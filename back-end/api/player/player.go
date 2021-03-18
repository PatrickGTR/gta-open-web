package player

import (
	"database/sql"

	"github.com/go-chi/chi"

	"github.com/open-backend/session"
)

type PlayerService struct {
	Routes chi.Router
	db     *sql.DB
}

func New(db *sql.DB) *PlayerService {
	router := chi.NewRouter()

	service := &PlayerService{
		Routes: router,
		db:     db,
	}

	router.Post("/", service.verifyUser)                                     // Login
	router.Delete("/", session.WithAuthentication(service.logout))           // Logout
	router.Get("/", session.WithAuthentication(service.getDataBySessionUID)) // Dashboard
	return service
}
