package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	user     = "postgres"
	password = "A.Ru.3729#"
	dbname   = "bookshop"
	port     = 5432
	host     = "localhost"
)

func ConnectDB() *sql.DB {
	psqlconn := fmt.Sprintf("user=%s password=%s dbname=%s port=%d host=%s sslmode=disable", user, password, dbname, port, host)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	err = db.Ping()
	CheckError(err)
	return db
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
