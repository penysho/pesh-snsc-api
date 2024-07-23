package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"github.com/penysho/pesh-snsc-api/post-app/infrastructure/config"
	"github.com/penysho/pesh-snsc-api/post-app/infrastructure/db"
	"github.com/penysho/pesh-snsc-api/post-app/infrastructure/router"
)

// https://github.com/awslabs/aws-lambda-go-api-proxy
var ginLambda *ginadapter.GinLambda

func init() {
	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Gin cold start")

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
	r.GET("/post-app", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "post-app",
		})
	})

	ginLambda = ginadapter.New(r)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
