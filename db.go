package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql",
		"root:root@tcp(127.0.0.1:8889)/shades")
	if err != nil {
		log.Fatal(err)
	}

	return db, err
}
