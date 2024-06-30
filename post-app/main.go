package main

import (
	"log"
	"net/http"
	"os"
	"post-app/infrastructure/db"

	"post-app/infrastructure/router"
)

func main() {
	database, err := db.NewDB()
	if err != nil {
		os.Exit(1)
	}

	dbManeger, err := db.NewDBManeger(database)
	if err != nil {
		os.Exit(1)
	}
	defer dbManeger.Close()

	r := router.NewGinRouter(dbManeger)
	s := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:8081",
	}

	log.Fatal(s.ListenAndServe())
}
