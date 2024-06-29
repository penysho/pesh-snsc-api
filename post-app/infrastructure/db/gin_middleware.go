package db

import (
	"database/sql"
	"net/http"

	postRepo "post-app/infrastructure/db/repository/post"
	postInteractor "post-app/usecase/interactor/post"

	"github.com/gin-gonic/gin"
)

func DBTxMiddleware(dbManeger *DBManeger) gin.HandlerFunc {
	return func(c *gin.Context) {
		pool := dbManeger.GetPool()
		dbTxManeger, err := NewDBTxManegerWithPool(
			pool,
			c,
			sql.TxOptions{},
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to begin transaction"})
			c.Abort()
			return
		}
		tx := dbTxManeger.GetTx()

		defer func() {
			if p := recover(); p != nil {
				dbTxManeger.RollbackTx()
				panic(p)
			} else if c.Writer.Status() != http.StatusOK {
				dbTxManeger.RollbackTx()
			} else {
				err = dbTxManeger.CommitTx()
				if err != nil {
					dbTxManeger.RollbackTx()
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit transaction"})
					c.Abort()
				}
			}
		}()

		postRepo := postRepo.NewPostRepository(c, pool, tx)
		postInteractor := postInteractor.NewPostInteractor(postRepo)
		c.Set("postInteractor", postInteractor)
		c.Next()
	}
}