package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"post-app/infrastructure/config"
	"post-app/infrastructure/db"
	"strconv"

	"post-app/infrastructure/router"
)

func main() {
	appConfig, err := config.NewAppConfig()
	if err != nil {
		os.Exit(1)
	}
	dbConfig, err := config.NewDBConfig()
	if err != nil {
		os.Exit(1)
	}

	database, err := db.NewDB()
	if err != nil {
		os.Exit(1)
	}

	dbManeger, err := db.NewDBManeger(database, dbConfig)
	if err != nil {
		os.Exit(1)
	}
	defer dbManeger.Close()

	r := router.NewGinRouter(dbManeger)
	s := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("%s:%s", appConfig.AppHost, strconv.Itoa(int(appConfig.AppPort))),
	}

	log.Fatal(s.ListenAndServe())
}
