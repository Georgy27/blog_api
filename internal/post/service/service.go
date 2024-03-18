package service

import (
	"context"
	"github.com/Georgy27/blogger_api/internal/post/model"
)

type PostService interface {
	CreatePost(ctx context.Context, info *model.PostInfo) (*model.Post, error)
	GetPost(ctx context.Context, id int64) (*model.Post, error)
	UpdatePost(ctx context.Context, id int64, info *model.UpdatePostInfo) error
	DeletePost(ctx context.Context, id int64) error
	ListPosts(ctx context.Context) ([]*model.Post, error)
}
