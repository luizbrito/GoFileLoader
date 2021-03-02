package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

//Postgres conection data
const (
	//host     = "localhost"
	host     = "database"
	port     = 5432
	user     = "go_user"
	password = "go_password"
	dbname   = "go_database"
)

// function return a database connection
func postgresConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	//Open the comunication with database and return a connection or a error
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	return db
}

// function to close database connection
// Call this function when finalize the insert
func closeDB(db *sql.DB) {
	defer db.Close()
}
