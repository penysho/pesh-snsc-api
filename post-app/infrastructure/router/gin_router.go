package router

import (
	"github.com/penysho/pesh-snsc-api/post-app/apidef/generated/server"
	"github.com/penysho/pesh-snsc-api/post-app/infrastructure/db"
	"github.com/penysho/pesh-snsc-api/post-app/infrastructure/db/middleware"
	"github.com/penysho/pesh-snsc-api/post-app/infrastructure/module"
	"github.com/penysho/pesh-snsc-api/post-app/interface/controller"

	"github.com/gin-gonic/gin"
)

func NewGinRouter(dbManeger *db.DBManeger) *gin.Engine {
	r := gin.New()

	r.Use(middleware.DBTxMiddleware(dbManeger))
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	interactProvider := module.NewInteractProvider()
	serverImpl := controller.NewServer(interactProvider)
	server.RegisterHandlers(r, serverImpl)
	return r
}
