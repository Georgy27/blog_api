package repository

import (
	"context"
	"github.com/Georgy27/blogger_api/internal/blog/model"
	"github.com/Georgy27/blogger_api/pkg/db"
)

type blogRepo struct {
	db db.Client
}

func NewBlogRepository(db db.Client) BlogRepository {
	return &blogRepo{
		db: db,
	}
}

func (r *blogRepo) CreateBlog(ctx context.Context, info *model.BlogInfo) (int64, error) {
	return 0, nil
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
