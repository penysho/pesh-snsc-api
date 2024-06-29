package router

import (
	"context"
	"net/http"
	server "post-app/apidef/generated"
	"post-app/infrastructure/db"
	postRepoImpl "post-app/infrastructure/db/repository/post"
	"post-app/interface/controller"

	"github.com/gin-gonic/gin"
)

func NewGinRouter(dbManeger *db.DBManeger) *gin.Engine {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		ctx := context.Background()
		dbTxManeger, err := db.NewDBTxManeger(dbManeger)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to begin transaction"})
			c.Abort()
			return
		}
		tx := dbTxManeger.GetTx()

		defer func() {
			if r := recover(); r != nil {
				dbTxManeger.RollbackTx()
				panic(r)
			} else if c.Writer.Status() != http.StatusOK {
				tx.Rollback()
			} else {
				err = dbTxManeger.CommitTx()
				if err != nil {
					dbTxManeger.RollbackTx()
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit transaction"})
					c.Abort()
				}
			}
		}()

		postRepo := postRepoImpl.NewPostRepositoryImpl(ctx, tx)
		c.Set("postRepo", postRepo)
		c.Next()
	})

	serverImpl := controller.NewServer()
	server.RegisterHandlers(r, serverImpl)
	return r
}
