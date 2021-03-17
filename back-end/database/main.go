package database

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type DBService struct {
	Routes chi.Router
	db     *sql.DB
}

func New() (*DBService, *sql.DB) {

	db, err := connect()
	if err != nil {
		fmt.Println("An error connecting to MySQL server", err)
	}
	router := chi.NewRouter()

	service := &DBService{
		Routes: router,
		db:     db,
	}

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		err := db.Ping()
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, map[string]string{"status": "conection failed"})
			return
		}

		render.Status(r, http.StatusOK)
		render.JSON(w, r, map[string]string{"status": "conection alive"})
		return
	})

	return service, db
}

func connect() (db *sql.DB, err error) {
	username := os.Getenv("MYSQL_USERNAME")
	password := os.Getenv("MYSQL_PASSWORD")
	database := os.Getenv("MYSQL_DATABASE")
	server := os.Getenv("MYSQL_SERVER")

	db, err = sql.Open("mysql", username+":"+password+"@tcp("+server+")/"+database+"")
	if err != nil {
		return
	}
	fmt.Println("Successfully connected to mysql server ->", server, "| db ->", database)

	return db, err

}
