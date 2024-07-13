package post

import (
	"context"

	domainError "github.com/penysho/pesh-snsc-api/post-app/entity/error"
	"github.com/penysho/pesh-snsc-api/post-app/entity/post"
	"github.com/penysho/pesh-snsc-api/post-app/infrastructure/db"
	"github.com/penysho/pesh-snsc-api/post-app/infrastructure/db/repository/models"
	"github.com/penysho/pesh-snsc-api/post-app/infrastructure/logger"
	postRepo "github.com/penysho/pesh-snsc-api/post-app/usecase/repository/post"
)

type postRepositoryImpl struct {
	ctx         context.Context
	dbManeger   *db.DBManeger
	dbTxManeger *db.DBTxManeger
}

func NewPostRepository(
	ctx context.Context,
	dbManeger *db.DBManeger,
	dbTxManeger *db.DBTxManeger,
) postRepo.PostRepository {
	return &postRepositoryImpl{
		ctx:         ctx,
		dbManeger:   dbManeger,
		dbTxManeger: dbTxManeger,
	}
}

func (r *postRepositoryImpl) FindByID(id int) (*post.Post, error) {
	tran := r.dbTxManeger.GetTx()

	postModel, err := models.Posts(
		models.PostWhere.ID.EQ(int64(id)),
	).One(r.ctx, tran)
	if err != nil {
		logger.Error("投稿情報が見つかりません", "id", id)
		return nil, domainError.NotFound
	}

	return post.NewPost(
		int(postModel.ID),
		postModel.Title.String,
		postModel.LikeCount.Int,
		postModel.CommentsCount.Int,
		postModel.Caption.String,
		postModel.Permalink,
		postModel.PostedAt,
	), nil
}
