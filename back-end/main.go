package main

import (
	"context"

	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/open-backend/api/media"
	"github.com/open-backend/api/player"
	"github.com/open-backend/api/server"
	"github.com/open-backend/database"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

func main() {
	if os.Getenv("ENV") != "PROD" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		fmt.Println(".env file loaded")
	}

	fx.New(
		fx.Options(
			fx.Provide(
				initApp,      // start http server, cors settings, and logging.
				database.New, // Initialise database connection & create /ping route
				player.New,   // Initialise endpoints for players
				media.New,    // Initialise endpoints for media
				server.New,
			),
			fx.Invoke(
				func(
					router chi.Router,
					db *database.DBService,
					ply *player.PlayerService,
					media *media.MediaService,
					srv *server.ServerService) {

					// conection check
					router.Mount("/ping", db.Routes)

					// /user/ -> POST (login)
					// /user/ -> DELETE (logout)
					// /user/ -> GET (get data from session uid)
					router.Mount("/user", ply.Routes)

					// server statistics displayed in
					// homepage of the website
					router.Mount("/server", srv.Routes)

					// /media/ -> POST (add new media )
					// /media/ -> GET (get all media)
					// /media/{id} -> GET (get media by specified id)
					// /media/trending -> GET (retrieves newest and most viewed media)
					// /media/add_views -> POST (increments view each page refresh)

					// media -> comment
					// /comment/ -> POST (creates a new comment entry)
					// /comment/{mediaid} -> GET (retrieves all comments of given media id)
					router.Mount("/media", media.Routes)
				}),
		),
	).Run()
}

func initApp(lc fx.Lifecycle) (chi.Router, error) {

	router := chi.NewRouter()

	router.Use(
		middleware.RequestID,
		middleware.RealIP,
		middleware.Logger,
		middleware.Recoverer,
		cors.Handler(cors.Options{
			AllowedOrigins: []string{
				"http://localhost:3000", // localhost, development mainly in // npm run dev
				"https://gta-open.ga",   // live demo, preview
			},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: true,
			MaxAge:           300,
		}))

	server := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go server.ListenAndServe()
			fmt.Println("Web server is running on port", server.Addr)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			server.Shutdown(ctx)

			return nil
		},
	})

	return router, nil
}
