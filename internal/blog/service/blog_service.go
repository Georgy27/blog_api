package service

import (
	"context"
	"github.com/Georgy27/blogger_api/internal/blog/model"
	blogRepository "github.com/Georgy27/blogger_api/internal/blog/repository"
)

type blogServ struct {
	blogRepository blogRepository.BlogRepository
}

func NewBlogService(repo blogRepository.BlogRepository) BlogService {
	return &blogServ{
		blogRepository: repo,
	}
}

func (s *blogServ) CreateBlog(ctx context.Context, info *model.BlogInfo) (int64, error) {
	return s.blogRepository.CreateBlog(ctx, info)
}

func (s *blogServ) GetBlog(ctx context.Context, id int64) (*model.Blog, error) {
	return s.blogRepository.GetBlog(ctx, id)
}

func (s *blogServ) UpdateBlog(ctx context.Context, id int64, info *model.BlogInfo) error {
	return s.blogRepository.UpdateBlog(ctx, id, info)
}

func (s *blogServ) DeleteBlog(ctx context.Context, id int64) error {
	return s.blogRepository.DeleteBlog(ctx, id)
}

func (s *blogServ) ListBlogs(ctx context.Context) ([]*model.Blog, error) {
	return s.blogRepository.ListBlogs(ctx)
}
