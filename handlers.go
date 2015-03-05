package main

import (
	//	"database/sql"
	//	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	//	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	// index
	fmt.Fprintln(w, "Welcome to the shades!")
}

func ShadesIndex(w http.ResponseWriter, r *http.Request) {
	// index of all shades
	db, err := Connect()
	if err != nil {
		log.Fatal(err)
	}
	var (
		Id  int
		Hex string
	)
	// db query
	rows, err := db.Query("SELECT Id, Hex FROM shades WHERE Grey = ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&Id, &Hex)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintln(w, Id, Hex)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	//	db.Close()
	Close(db)
}

func ShadeShow(w http.ResponseWriter, r *http.Request) {
	// show a particular shade
}

func ShadeRand(w http.ResponseWriter, r *http.Request) {
	// show a random shade
}

func ShadesRand(w http.ResponseWriter, r *http.Request) {
	// show a number of random shades
}

func ShadeByHEX(w http.ResponseWriter, r *http.Request) {
	// show a shade by certain category
}
