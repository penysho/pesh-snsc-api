package main

import (
	"log"
	"net/http"
	"post-app/infrastructure/router"
)

func main() {
	r := router.NewGinRouter()

	s := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:8081",
	}

	log.Fatal(s.ListenAndServe())
}
