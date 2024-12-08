package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/penysho/pesh-snsc-api/post-app/infrastructure/config"
	"github.com/penysho/pesh-snsc-api/post-app/infrastructure/db"
	"github.com/penysho/pesh-snsc-api/post-app/infrastructure/logger"
	"github.com/penysho/pesh-snsc-api/post-app/infrastructure/router"
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

	logger.InitLoggerWithLevel(logger.LevelDebug)

	r := router.NewGinRouter(dbManeger)
	s := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("%s:%s", appConfig.AppHost, strconv.Itoa(int(appConfig.AppPort))),
	}

	log.Fatal(s.ListenAndServe())
}
