package router

import (
	server "post-app/apidef/generated"
	"post-app/infrastructure/db"
	"post-app/interface/controller"

	"github.com/gin-gonic/gin"
)

func NewGinRouter(dbManeger *db.DBManeger) *gin.Engine {
	r := gin.Default()

	r.Use(db.DBTxMiddleware(dbManeger))

	interactProvider := controller.NewInteractProvider()
	serverImpl := controller.NewServer(interactProvider)
	server.RegisterHandlers(r, serverImpl)
	return r
}
