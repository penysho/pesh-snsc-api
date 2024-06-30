package router

import (
	server "post-app/apidef/generated"
	"post-app/infrastructure/db"
	"post-app/infrastructure/module"
	"post-app/interface/controller"

	"github.com/gin-gonic/gin"
)

func NewGinRouter(dbManeger *db.DBManeger) *gin.Engine {
	r := gin.New()

	r.Use(db.DBTxMiddleware(dbManeger))
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	interactProvider := module.NewInteractProvider()
	serverImpl := controller.NewServer(interactProvider)
	server.RegisterHandlers(r, serverImpl)
	return r
}
