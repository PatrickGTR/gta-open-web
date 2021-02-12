package main

// TODO
// Review code & refactor

import (
	"fmt"
	"net/http"

	"github.com/open-backend/api"
	"github.com/open-backend/user"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

func main() {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins:  []string{"Origin"},
		AllowOriginFunc: func(r *http.Request, origin string) bool { return true },
		AllowedMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:  []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		//ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		//MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// A good base middleware stack
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// nested route
	router.Route("/user", func(r chi.Router) {
		r.Post("/", api.VerifyUser)                      // Login
		r.Delete("/", isLoggedIn(api.Logout))            // Logout
		r.Get("/{userid}", isLoggedIn(api.GetDataByUID)) // Dashboard
	})

	router.Route("/server", func(r chi.Router) {
		r.Get("/stats", api.ServerStats)
		r.Get("/banlist", api.BanList)
	})
	PORT := 8000
	fmt.Println("Web server is running on port", PORT)
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), router)
}

func isLoggedIn(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// grab the content of the session.
		session, _ := user.Cookie.Get(r, "sessionid")
		sessionContent := session.Values["accountID"]
		if sessionContent == nil {
			// no session id set
			render.Status(r, http.StatusUnauthorized)
			return
		}

		sessionUID := session.Values["accountID"].(int)
		if sessionUID <= 0 {
			// account ID starts at 1
			render.Status(r, http.StatusUnauthorized)
			return
		}

		// proceed to the next route
		next(w, r)
	})
}
