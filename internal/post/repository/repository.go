package repository

import (
	"context"
	"github.com/Georgy27/blogger_api/internal/post/model"
)

type PostRepository interface {
	CreatePost(ctx context.Context, info *model.PostInfo) (*model.Post, error)
	GetPost(ctx context.Context, id int64) (*model.Post, error)
	UpdatePost(ctx context.Context, id int64, info *model.UpdatePostInfo) error
	DeletePost(ctx context.Context, id int64) error
	ListPosts(ctx context.Context, limit int64, offset int64) ([]*model.Post, error)
	CountPosts(ctx context.Context) (int64, error)
}
