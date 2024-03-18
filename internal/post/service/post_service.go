package service

import (
	"context"
	blogService "github.com/Georgy27/blogger_api/internal/blog/service"
	"github.com/Georgy27/blogger_api/internal/post/model"
	postModel "github.com/Georgy27/blogger_api/internal/post/model"
	postRepository "github.com/Georgy27/blogger_api/internal/post/repository"
	commonErrors "github.com/Georgy27/blogger_api/pkg/errors"
	"github.com/Georgy27/blogger_api/pkg/errors/codes"

	"github.com/Georgy27/blogger_api/pkg/db"
)

type postServ struct {
	postRepository postRepository.PostRepository
	blogService    blogService.BlogService
	txManager      db.TxManager
}

func NewPostService(repo postRepository.PostRepository, blogService blogService.BlogService, txManager db.TxManager) *postServ {
	return &postServ{
		postRepository: repo,
		blogService:    blogService,
		txManager:      txManager,
	}
}

func (s *postServ) CreatePost(ctx context.Context, info *model.PostInfo) (*model.Post, error) {
	var post *postModel.Post

	blog, err := s.blogService.GetBlog(ctx, info.BlogID)
	if err != nil {
		return nil, commonErrors.NewCommonError("blog not found", codes.NotFound)
	}

	err = s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		post, errTx = s.postRepository.CreatePost(ctx, info)
		if errTx != nil {
			return errTx
		}

		post.Info.BlogName = blog.Info.Name

		return nil
	})

	if err != nil {
		return nil, err
	}

	return post, nil
}

func (s *postServ) GetPost(ctx context.Context, id int64) (*model.Post, error) {
	return nil, nil
}

func (s *postServ) UpdatePost(ctx context.Context, id int64, info *model.UpdatePostInfo) error {
	return nil
}

func (s *postServ) DeletePost(ctx context.Context, id int64) error {
	return nil
}

func (s *postServ) ListPosts(ctx context.Context) ([]*model.Post, error) {
	return nil, nil
}
