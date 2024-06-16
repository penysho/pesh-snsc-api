package main

import (
	"log"
	"post-app/router"
)

func main() {
	router := router.SetRouter()
	if err := router.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}
