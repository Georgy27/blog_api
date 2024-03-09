package handler

import (
	"context"
	blogService "github.com/Georgy27/blogger_api/internal/blog/service"
	bloggerV1 "github.com/Georgy27/blogger_api/pkg/blogger_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type BlogHandler struct {
	bloggerV1.UnimplementedBloggerV1Server
	blogService blogService.BlogService
}

func NewBlogHandler(blogService blogService.BlogService) *BlogHandler {
	return &BlogHandler{
		blogService: blogService,
	}
}

func (h *BlogHandler) CreateBlog(ctx context.Context, req *bloggerV1.CreateBlogRequest) (*bloggerV1.CreateBlogResponse, error) {
	return nil, nil
}

func (h *BlogHandler) GetBlog(ctx context.Context, req *bloggerV1.GetBlogRequest) (*bloggerV1.GetBlogResponse, error) {
	return nil, nil
}

func (h *BlogHandler) UpdateBlog(ctx context.Context, req *bloggerV1.UpdateBlogRequest) (*emptypb.Empty, error) {
	return nil, nil
}

func (h *BlogHandler) DeleteBlog(ctx context.Context, req *bloggerV1.DeleteBlogRequest) (*emptypb.Empty, error) {
	return nil, nil
}

func (h *BlogHandler) ListBlogs(ctx context.Context, req *bloggerV1.ListBlogsRequest) (*bloggerV1.ListBlogsResponse, error) {
	return nil, nil
}
