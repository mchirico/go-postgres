package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

const (
	port   = 5432
	user   = "postgres"
	dbname = "postgres"
)

func main() {

	password := os.Getenv("POSTGRESQL_M_PASSWORD")
	host := os.Getenv("POSTGRESQL_M_HOST")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=require",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}
