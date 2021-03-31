package media

import (
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/open-backend/session"
)

type MediaService struct {
	Routes chi.Router
	db     *sql.DB
}

func New(db *sql.DB) *MediaService {
	router := chi.NewRouter()

	service := &MediaService{
		Routes: router,
		db:     db,
	}

	router.Post("/", session.WithAuthentication(service.addMedia))
	router.Get("/", service.getAll)
	router.Get("/{id}", service.getOne)
	router.Get("/trending", service.getTrending)
	router.Post("/add_views", service.incrementViews)

	router.Route("/comment", func(r chi.Router) {
		r.Post("/", session.WithAuthentication(service.postComment))
		r.Get("/{mediaid}", service.getComments)
	})

	return service
}
