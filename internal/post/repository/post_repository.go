package repository

import (
	"context"
	"github.com/Georgy27/blogger_api/internal/post/model"
	"github.com/Georgy27/blogger_api/internal/post/repository/converter"
	modelRepo "github.com/Georgy27/blogger_api/internal/post/repository/model"
	"github.com/Georgy27/blogger_api/pkg/db"
	sq "github.com/Masterminds/squirrel"
	"time"
)

const (
	tableName = "post"

	columnID               = "id"
	columnTitle            = "title"
	columnContent          = "content"
	columnShortDescription = "short_description"
	columnBlogID           = "blog_id"
	columnBlogName         = "name"
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
	builder := sq.Select("post.id", columnBlogID, columnTitle, columnContent,
		columnShortDescription, "blog.name as blog_name", "post.created_at", "post.updated_at").
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Join("blog ON blog.id = post.blog_id").
		Where(sq.Eq{"post.id": id})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "post_repo_get_post",
		QueryRaw: query,
	}

	var post modelRepo.Post

	err = r.db.DB().ScanOneContext(ctx, &post, q, args...)

	if err != nil {
		return nil, err
	}

	return converter.ToPostFromRepo(&post), nil
}

func (r *postRepo) UpdatePost(ctx context.Context, id int64, info *model.UpdatePostInfo) error {
	builder := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{columnID: id})

	if info.Title.Valid {
		builder = builder.Set(columnTitle, info.Title.String)
	}

	if info.ShortDescription.Valid {
		builder = builder.Set(columnShortDescription, info.ShortDescription.String)
	}

	if info.Content.Valid {
		builder = builder.Set(columnContent, info.Content.String)
	}

	builder = builder.Set(columnUpdatedAt, time.Now().UTC().Format(time.RFC3339))

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "post_repo_update_post",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)

	if err != nil {
		return err
	}

	return nil

}

func (r *postRepo) DeletePost(ctx context.Context, id int64) error {
	return nil
}

func (r *postRepo) ListPosts(ctx context.Context, limit int64, offset int64) ([]*model.Post, error) {
	builder := sq.Select("post.id", columnBlogID, columnTitle, columnContent, columnShortDescription,
		"blog.name as blog_name", "post.created_at", "post.updated_at").
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Join("blog ON blog.id = post.blog_id")

	if limit != 0 {
		builder = builder.Limit(uint64(limit))
	} else {
		builder = builder.Limit(uint64(10))
	}

	builder = builder.Offset(uint64(offset))

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "post_repo_list_posts",
		QueryRaw: query,
	}

	var posts []*modelRepo.Post
	err = r.db.DB().ScanAllContext(ctx, &posts, q, args...)
	if err != nil {
		return nil, err
	}

	return converter.ToPostsFromRepo(posts), nil
}

func (r *postRepo) CountPosts(ctx context.Context) (int64, error) {
	builder := sq.Select("count(*)").
		PlaceholderFormat(sq.Dollar).
		From(tableName)

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "post_repo_count_posts",
		QueryRaw: query,
	}

	var count int64
	err = r.db.DB().ScanOneContext(ctx, &count, q, args...)
	if err != nil {
		return 0, err
	}

	return count, nil

}
