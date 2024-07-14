package middleware

import (
	"database/sql"
	"net/http"

	"github.com/penysho/pesh-snsc-api/post-app/infrastructure/db"
	postRepo "github.com/penysho/pesh-snsc-api/post-app/infrastructure/db/repository/post"
	"github.com/penysho/pesh-snsc-api/post-app/infrastructure/logger"
	postInteractor "github.com/penysho/pesh-snsc-api/post-app/usecase/interactor/post"

	"github.com/gin-gonic/gin"
)

// DBTxMiddleware リクエスト毎にトランザクションマネージャーを生成してInteractorを定義する
func DBTxMiddleware(dbManeger *db.DBManeger) gin.HandlerFunc {
	return func(c *gin.Context) {
		pool := dbManeger.GetPool()
		dbTxManeger, err := db.NewDBTxManegerWithPool(
			pool,
			c,
			sql.TxOptions{},
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to begin transaction"})
			c.Abort()
			return
		}

		defer func() {
			if p := recover(); p != nil {
				logger.Error("システムエラーのためDBトランザクションをロールバックします")
				dbTxManeger.RollbackTx()
				panic(p)
			} else if c.Writer.Status() != http.StatusOK {
				logger.Error("ステータスコードが異常なためDBトランザクションをロールバックします")
				dbTxManeger.RollbackTx()
			} else {
				err = dbTxManeger.CommitTx()
				if err != nil {
					logger.Error("DBトランザクションのコミットに失敗したためロールバックします")
					dbTxManeger.RollbackTx()
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit transaction"})
					c.Abort()
				}
			}
		}()

		postRepo := postRepo.NewPostRepository(c, dbManeger, dbTxManeger)
		postInteractor := postInteractor.NewPostInteractor(postRepo)
		c.Set("postInteractor", postInteractor)
		c.Next()
	}
}
