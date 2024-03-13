package service

import (
	"context"
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
			return errTx
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return blog, nil
}

func (s *blogServ) GetBlog(ctx context.Context, id int64) (*model.Blog, error) {
	blog, err := s.blogRepository.GetBlog(ctx, id)
	if err != nil {
		return nil, err
	}

	return blog, nil
}

func (s *blogServ) UpdateBlog(ctx context.Context, id int64, info *model.UpdateBlogInfo) error {
	_, err := s.blogRepository.GetBlog(ctx, id)
	if err != nil {
		return err
	}

	err = s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		errTx = s.blogRepository.UpdateBlog(ctx, id, info)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *blogServ) DeleteBlog(ctx context.Context, id int64) error {
	_, err := s.blogRepository.GetBlog(ctx, id)
	if err != nil {
		return err
	}

	err = s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		errTx = s.blogRepository.DeleteBlog(ctx, id)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *blogServ) ListBlogs(ctx context.Context) ([]*model.Blog, error) {
	blogs, err := s.blogRepository.ListBlogs(ctx)
	if err != nil {
		return nil, err
	}

	return blogs, nil
}
