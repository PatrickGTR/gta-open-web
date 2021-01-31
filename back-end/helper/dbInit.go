package helper

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var mysqlHandle *sql.DB

func init() {
	// Development purposes only, remove on production.
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mysqlHandle = InitDatabase()
}

func InitDatabase() (db *sql.DB) {

	username := os.Getenv("MYSQL_USERNAME")
	password := os.Getenv("MYSQL_PASSWORD")
	database := os.Getenv("MYSQL_DATABASE")
	server := os.Getenv("MYSQL_SERVER")
	db, err := sql.Open("mysql", username+":"+password+"@tcp("+server+")/"+database+"")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Successfully connected to mysql server ->", server, "| db ->", database)
	return db
}

func ExecuteQuery(query string, arg ...interface{}) (rows *sql.Rows, err error) {
	rows, err = mysqlHandle.Query(query, arg...)
	return
}
