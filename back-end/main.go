package main

// TODO
// Modularize
// Create Helper Function
// Add Cors
// Add Rate limit (to non server caller)
// env for MySQL credentials

import (
	"fmt"
	"net/http"

	"github.com/open-backend/helper"
	"github.com/open-backend/routes"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	helper.SqlHandle = helper.InitDatabase()
	// close connection when function ends
	defer helper.SqlHandle.Close()

	router := chi.NewRouter()

	// A good base middleware stack
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// nested route
	router.Route("/user", func(r chi.Router) {
		r.Get("/{userid}", routes.GetDataByUID)
		r.Post("/", routes.Login)
	})

	PORT := 8000
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), router)
	fmt.Printf("Web server running on port %d", PORT)
}
