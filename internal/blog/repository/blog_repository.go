package repository

import (
	"fmt"
	"github.com/Georgy27/blogger_api/internal/blog/repository/converter"
	sq "github.com/Masterminds/squirrel"

	"context"
	"github.com/Georgy27/blogger_api/internal/blog/model"
	modelRepo "github.com/Georgy27/blogger_api/internal/blog/repository/model"
	"github.com/Georgy27/blogger_api/pkg/db"
)

const (
	tableName = "blog"

	columnName        = "name"
	columnDescription = "description"
	columnWebsiteUrl  = "website_url"
	columnCreatedAt   = "created_at"
	columnUpdatedAt   = "updated_at"
	columnID          = "id"
)

type blogRepo struct {
	db db.Client
}

func NewBlogRepository(db db.Client) BlogRepository {
	return &blogRepo{
		db: db,
	}
}

func (r *blogRepo) CreateBlog(ctx context.Context, info *model.BlogInfo) (*model.Blog, error) {
	fmt.Println(info.Name, info.Description, info.WebsiteUrl)
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(columnName, columnDescription, columnWebsiteUrl).
		Values(info.Name, info.Description, info.WebsiteUrl).
		Suffix("RETURNING *")

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "blog_repo_create_blog",
		QueryRaw: query,
	}

	var blog modelRepo.Blog

	err = r.db.DB().ScanOneContext(ctx, &blog, q, args...)

	if err != nil {
		return nil, err
	}

	return converter.ToBlogFromRepo(&blog), nil

}

func (r *blogRepo) GetBlog(ctx context.Context, id int64) (*model.Blog, error) {
	return nil, nil
}

func (r *blogRepo) UpdateBlog(ctx context.Context, id int64, info *model.BlogInfo) error {
	return nil
}

func (r *blogRepo) DeleteBlog(ctx context.Context, id int64) error {
	return nil
}

func (r *blogRepo) ListBlogs(ctx context.Context) ([]*model.Blog, error) {
	return nil, nil
}
