package repository

import (
	"fmt"
	"github.com/Georgy27/blogger_api/internal/blog/repository/converter"
	sq "github.com/Masterminds/squirrel"
	"time"

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
	builder := sq.Select("*").
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{columnID: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "blog_repo_get_blog",
		QueryRaw: query,
	}

	var blog modelRepo.Blog

	err = r.db.DB().ScanOneContext(ctx, &blog, q, args...)

	if err != nil {
		return nil, err
	}

	return converter.ToBlogFromRepo(&blog), nil
}

func (r *blogRepo) UpdateBlog(ctx context.Context, id int64, blogInfo *model.UpdateBlogInfo) error {
	builder := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{columnID: id})

	if blogInfo.Name.Valid {
		builder = builder.Set(columnName, blogInfo.Name)
	}

	if blogInfo.Description.Valid {
		builder = builder.Set(columnDescription, blogInfo.Description)
	}

	if blogInfo.WebsiteUrl.Valid {
		builder = builder.Set(columnWebsiteUrl, blogInfo.WebsiteUrl)
	}

	builder = builder.Set(columnUpdatedAt, time.Now().UTC().Format(time.RFC3339))

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "blog_repo_update_blog",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *blogRepo) DeleteBlog(ctx context.Context, id int64) error {

	builder := sq.Delete(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{columnID: id})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "blog_repo_delete_blog",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *blogRepo) ListBlogs(ctx context.Context) ([]*model.Blog, error) {
	builder := sq.Select("*").
		PlaceholderFormat(sq.Dollar).
		From(tableName)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "blog_repo_list_blogs",
		QueryRaw: query,
	}

	var blogs []*modelRepo.Blog
	err = r.db.DB().ScanAllContext(ctx, &blogs, q, args...)
	if err != nil {
		return nil, err
	}

	return converter.ToBlogsFromRepo(blogs), nil
}
