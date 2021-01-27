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
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/open-backend/helper"
	"github.com/open-backend/routes"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	helper.SqlHandle = helper.InitDatabase()
	// close connection when function ends
	defer helper.SqlHandle.Close()

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		//ExposedHeaders:   []string{"Link"},
		//AllowCredentials: false,
		//MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// A good base middleware stack
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// nested route
	router.Route("/user", func(r chi.Router) {
		r.Get("/{userid}", ValidateMiddleware(routes.GetDataByUID))
		r.Post("/", routes.Login)
	})

	PORT := 8000
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), router)
	fmt.Printf("Web server running on port %d", PORT)
}

// Exception is an alias
type Exception helper.MessageData

func ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// No authorization header found
		authHeader := r.Header.Get("authorization")
		if authHeader == "" {
			render.JSON(w, r, &Exception{
				Code:    "no.token",
				Message: "Authorization header required.",
			})
			return
		}

		// Split [Bearer Token]
		// index[0] -> Bearer
		// index[1] -> Token
		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 {
			render.JSON(w, r, &Exception{
				Code:    "invalid.bearer.len",
				Message: "Bearer length invalid, must be 2.",
			})
			return
		}

		token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("SECRET_KEY")), nil
		})

		// An error occured during the process of retrieving token,
		// token can be expired, unverified etc..
		if err != nil {
			render.JSON(w, r, &Exception{
				Code:    "token.error",
				Message: err.Error(),
			})
			return
		}

		if !token.Valid {
			render.JSON(w, r, &Exception{
				Code:    "invalid.auth.token",
				Message: "Invalid authorization token received",
			})
			return
		}

		next(w, r)
	})
}
