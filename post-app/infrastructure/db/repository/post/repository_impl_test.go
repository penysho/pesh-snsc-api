package post_test

import (
	"context"
	"database/sql"
	"post-app/entity/post"
	"post-app/infrastructure/config"
	"post-app/infrastructure/db"
	"post-app/infrastructure/db/repository/models"
	postRepoImpl "post-app/infrastructure/db/repository/post"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var (
	ctx         context.Context
	dbManeger   *db.DBManeger
	dbTxManeger *db.DBTxManeger
)

func TestMain(t *testing.M) {
	dbConfig, _ := config.NewDBConfig()

	database, _ := db.NewDB()
	dbManeger, _ = db.NewDBManeger(database, dbConfig)
	defer dbManeger.Close()

	ctx = context.Background()
	dbTxManeger, _ = db.NewDBTxManegerWithPool(
		dbManeger.GetPool(),
		ctx,
		sql.TxOptions{},
	)
	defer dbTxManeger.RollbackTx()

	t.Run()
}

func Test_postRepositoryImpl_FindByID(t *testing.T) {
	t.Parallel()

	posts := createPostEntities(1)
	bulkInsertPost(t, ctx, dbTxManeger, posts)

	type fields struct {
		ctx         context.Context
		dbManeger   *db.DBManeger
		dbTxManeger *db.DBTxManeger
	}
	type args struct {
		id int
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      *post.Post
		assertion assert.ErrorAssertionFunc
	}{
		// 正常系
		{
			name: "投稿情報が存在/Entityが返却",
			fields: fields{
				ctx:         ctx,
				dbManeger:   dbManeger,
				dbTxManeger: dbTxManeger,
			},
			args: args{
				id: posts[0].GetId(),
			},
			want:      posts[0],
			assertion: assert.NoError,
		},
		// 異常系
		{
			name: "投稿情報が存在しない/エラーが返却",
			fields: fields{
				ctx:         ctx,
				dbManeger:   dbManeger,
				dbTxManeger: dbTxManeger,
			},
			args: args{
				id: 0,
			},
			want:      nil,
			assertion: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			r := postRepoImpl.NewPostRepository(tt.fields.ctx, tt.fields.dbManeger, tt.fields.dbTxManeger)
			defer tt.fields.dbTxManeger.RollbackTx()

			got, err := r.FindByID(tt.args.id)

			tt.assertion(t, err)
			if tt.want != nil {
				assert.Equal(t, tt.want.GetId(), got.GetId())
			}
		})
	}
}

func createPostEntities(size int) []*post.Post {
	postEntities := make([]*post.Post, 0, size)
	for i := 0; i < size; i++ {
		postEntity := post.NewPost(
			int(10000+i),
			"title",
			1,
			1,
			"caption",
			"https://example.com",
			time.Now(),
		)
		postEntities = append(postEntities, postEntity)
	}

	return postEntities
}

func bulkInsertPost(t *testing.T, ctx context.Context, dbTxManeger *db.DBTxManeger, postEntities []*post.Post) {
	t.Helper()

	tran := dbTxManeger.GetTx()

	for _, postEntity := range postEntities {
		postModel := models.Post{
			ID:    int64(postEntity.GetId()),
			Title: null.StringFrom(postEntity.GetTitle()),
			LikeCount: null.IntFrom(
				postEntity.GetLikeCount(),
			),
			CommentsCount: null.IntFrom(
				postEntity.GetCommentsCount(),
			),
			Caption:   null.StringFrom(postEntity.GetCaption()),
			Permalink: postEntity.GetPermalink(),
			PostedAt:  postEntity.GetPostedAt(),
		}

		err := postModel.Insert(ctx, tran, boil.Infer())
		if err != nil {
			t.Fatal("前準備のデータ登録に失敗しました", err)
		}
	}
}