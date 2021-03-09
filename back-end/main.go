package main

// TODO
// Review code & refactor

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/open-backend/api"
	"github.com/open-backend/session"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

func main() {
	router := chi.NewRouter()

	// A good base middleware stack
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	setupCORS(router)

	setupRoutes(router)
	srv := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}
	setupServer(srv)

}

func setupCORS(router *chi.Mux) {
	router.Use(cors.Handler(cors.Options{
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
}

func setupRoutes(router *chi.Mux) {
	// nested route
	router.Route("/user", func(r chi.Router) {
		r.Post("/", api.VerifyUser)             // Login
		r.Delete("/", isLoggedIn(api.Logout))   // Logout
		r.Get("/", isLoggedIn(api.GetSelfData)) // Dashboard
	})

	router.Route("/server", func(r chi.Router) {
		r.Get("/stats", api.ServerStats)
		r.Get("/banlist", api.BanList)
	})

	router.Route("/media", func(r chi.Router) {
		r.Post("/", isLoggedIn(api.MediaPost))
		r.Get("/", api.MediaGetAll)
		r.Get("/{id}", api.MediaGetOne)
		r.Get("/trending", api.MediaGet)
		r.Post("/add_views", api.MediaIncrementViews)

		// comment API
		r.Post("/comment", isLoggedIn(api.MediaPostComment))
		r.Get("/comment/{mediaid}", api.MediaGetComments)

	})
}

func setupServer(server *http.Server) {
	// Graceful server shutdown
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to initialize server: %v\n", err)
		}
	}()

	fmt.Println("Web server is running on port", server.Addr)

	// Wait for kill signal of channel
	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// This blocks until a signal is passed into the quit channel
	<-quit

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown server
	fmt.Println("Shutting down server...")
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	}
}

func isLoggedIn(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// grab the content of the session.
		// if there are none set, return unauthorized http status
		_, err := session.GetUID(r)
		if err != nil {
			render.Status(r, http.StatusUnauthorized)
			return
		}

		// proceed to the next route
		next(w, r)
	})
}
