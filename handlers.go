package main

import (
	//	"database/sql"
	"encoding/json"
	"time"
	//	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	// index route
	msg := Msg{Msg: "Welcome to the shades!", Time: time.Now()}
	// http-headers
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(msg); err != nil {
		panic(err)
	}
}

func ShadesIndex(w http.ResponseWriter, r *http.Request) {
	// index of all shades
	db, err := Connect()
	if err != nil {
		log.Fatal(err)
	}
	// db query
	rows, err := db.Query("SELECT Id, Grey, Hex FROM shades")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	shades := []Shade{}
	for rows.Next() {
		var r Shade
		err = rows.Scan(&r.Id, &r.Grey, &r.Hex)
		if err != nil {
			log.Fatal(err)
		}
		shades = append(shades, r)
	}

	// http-headers
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(shades); err != nil {
		panic(err)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	db.Close()
}

func ShadeShow(w http.ResponseWriter, r *http.Request) {
	// show a particular shade
	vars := mux.Vars(r)
	Pattern := vars["Pattern"]
	FilterType := vars["FilterType"]
	db, err := Connect()
	if err != nil {
		log.Fatal(err)
	}

	var res Shade

	if FilterType == "id" {
		err = db.QueryRow("SELECT Id, Grey, Hex FROM shades WHERE Id = ?", Pattern).Scan(&res.Id, &res.Grey, &res.Hex)
	} else if FilterType == "hex" {
		Pattern = "#" + Pattern
		err = db.QueryRow("SELECT Id, Grey, Hex FROM shades WHERE Hex = ?", Pattern).Scan(&res.Id, &res.Grey, &res.Hex)
	}

	// http-headers
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err != nil {
		//		log.Fatal("Hello", err)
		errMsg := Error{Error: "Query was not successful", Time: time.Now()}

		if err := json.NewEncoder(w).Encode(errMsg); err != nil {
			panic(err)
		}
	} else {
		if err := json.NewEncoder(w).Encode(res); err != nil {
			panic(err)
		}
	}
}

func ShadeRand(w http.ResponseWriter, r *http.Request) {
	// show a random shade
}

func ShadesRand(w http.ResponseWriter, r *http.Request) {
	// show a number of random shades
}
