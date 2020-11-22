package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "reblog"
	password = "reblog"
	dbname   = "reblog"
)

func main() {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname))

	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Successfully connected! DB Type: %T \n", db)

	insertUser(db)
}

func insertUser(db *sql.DB) {
	var lastInsertID int
	err := db.QueryRow("INSERT INTO public.user (username,password) VALUES($1,$2) returning id;", "nguyendaigia", "daigialaso1").Scan(&lastInsertID)
	if err != nil {
		panic(err)
	}
	fmt.Println("Last inserted id =", lastInsertID)
}
