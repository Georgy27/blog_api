package repository

import (
	"context"
	"github.com/Georgy27/blogger_api/internal/post/model"
	"github.com/Georgy27/blogger_api/internal/post/repository/converter"
	modelRepo "github.com/Georgy27/blogger_api/internal/post/repository/model"
	"github.com/Georgy27/blogger_api/pkg/db"
	sq "github.com/Masterminds/squirrel"
)

const (
	tableName = "post"

	columnID               = "id"
	columnTitle            = "title"
	columnContent          = "content"
	columnShortDescription = "short_description"
	columnBlogID           = "blog_id"
	columnBlogName         = "blog_name"
	columnCreatedAt        = "created_at"
	columnUpdatedAt        = "updated_at"
)

type postRepo struct {
	db db.Client
}

func NewPostRepository(db db.Client) PostRepository {
	return &postRepo{
		db: db,
	}
}

func (r *postRepo) CreatePost(ctx context.Context, info *model.PostInfo) (*model.Post, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(columnTitle, columnContent, columnShortDescription, columnBlogID).
		Values(info.Title, info.Content, info.ShortDescription, info.BlogID).
		Suffix("RETURNING *")

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "post_repo_create_post",
		QueryRaw: query,
	}

	var post modelRepo.Post

	err = r.db.DB().ScanOneContext(ctx, &post, q, args...)

	if err != nil {
		return nil, err
	}

	return converter.ToPostFromRepo(&post), nil

}
func (r *postRepo) GetPost(ctx context.Context, id int64) (*model.Post, error) {
	return nil, nil
}

func (r *postRepo) UpdatePost(ctx context.Context, id int64, info *model.UpdatePostInfo) error {
	return nil
}

func (r *postRepo) DeletePost(ctx context.Context, id int64) error {
	return nil
}

func (r *postRepo) ListPosts(ctx context.Context) ([]*model.Post, error) {
	return nil, nil
}
