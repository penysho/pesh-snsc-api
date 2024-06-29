package post

import (
	"context"
	"database/sql"
	"post-app/entity/post"
	"post-app/infrastructure/db/repository/models"
	postRepo "post-app/usecase/repository/post"
)

type postRepositoryImpl struct {
	ctx context.Context
	tx  *sql.Tx
}

func NewPostRepositoryImpl(
	ctx context.Context,
	tx *sql.Tx,
) postRepo.PostRepository {
	return &postRepositoryImpl{
		ctx: ctx,
		tx:  tx,
	}
}

func (r *postRepositoryImpl) FindByID(id int) (*post.Post, error) {
	postModel, err := models.Posts(
		models.PostWhere.ID.EQ(int64(id)),
	).One(r.ctx, r.tx)
	if err != nil {
		return nil, err
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
