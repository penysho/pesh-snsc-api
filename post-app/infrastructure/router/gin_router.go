package router

import (
	server "post-app/apidef/generated"
	"post-app/interface/controller"

	"github.com/gin-gonic/gin"
)

func NewGinRouter() *gin.Engine {
	serverImpl := controller.NewServer()

	r := gin.Default()

	server.RegisterHandlers(r, serverImpl)

	return r
}
