package service

import (
	"context"
	"fmt"
	"github.com/Georgy27/blogger_api/internal/blog/model"
	blogRepository "github.com/Georgy27/blogger_api/internal/blog/repository"
	"github.com/Georgy27/blogger_api/pkg/db"
)

type blogServ struct {
	blogRepository blogRepository.BlogRepository
	txManager      db.TxManager
}

func NewBlogService(repo blogRepository.BlogRepository, txManager db.TxManager) BlogService {
	return &blogServ{
		blogRepository: repo,
		txManager:      txManager,
	}
}

func (s *blogServ) CreateBlog(ctx context.Context, info *model.BlogInfo) (*model.Blog, error) {
	var blog *model.Blog

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		blog, errTx = s.blogRepository.CreateBlog(ctx, info)
		if errTx != nil {
			fmt.Println("inside service error")
			return errTx
		}

		fmt.Println("hello there from service")
		fmt.Println(blog)

		return nil
	})

	if err != nil {
		return nil, err
	}

	return blog, nil
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
