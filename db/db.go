package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConectDataBase() *sql.DB {
	str := "host=localhost port=25432 user=admin dbname=books sslmode=disable password=123456"

	db, err := sql.Open("postgres", str)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
