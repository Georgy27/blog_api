package service

import (
	"context"
	"github.com/Georgy27/blogger_api/internal/blog/model"
)

type BlogService interface {
	CreateBlog(ctx context.Context, info *model.BlogInfo) (int64, error)
	GetBlog(ctx context.Context, id int64) (*model.Blog, error)
	UpdateBlog(ctx context.Context, id int64, info *model.BlogInfo) error
	DeleteBlog(ctx context.Context, id int64) error
	ListBlogs(ctx context.Context) ([]*model.Blog, error)
}
