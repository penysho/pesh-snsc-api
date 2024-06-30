package post

import (
	"context"
	"database/sql"
	domainError "post-app/entity/error"
	"post-app/entity/post"
	"post-app/infrastructure/db/repository/models"
	"post-app/infrastructure/logger"
	postRepo "post-app/usecase/repository/post"
)

type postRepositoryImpl struct {
	ctx  context.Context
	pool *sql.DB
	tx   *sql.Tx
}

func NewPostRepository(
	ctx context.Context,
	pool *sql.DB,
	tx *sql.Tx,
) postRepo.PostRepository {
	return &postRepositoryImpl{
		ctx:  ctx,
		pool: pool,
		tx:   tx,
	}
}

func (r *postRepositoryImpl) FindByID(id int) (*post.Post, error) {
	postModel, err := models.Posts(
		models.PostWhere.ID.EQ(int64(id)),
	).One(r.ctx, r.tx)
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
