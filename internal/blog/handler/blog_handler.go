package handler

import (
	"context"
	"errors"
	"fmt"
	"github.com/Georgy27/blogger_api/internal/blog/converter"
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
	blogInfo := &bloggerV1.BlogInfo{
		Name:        req.GetName(),
		Description: req.GetDescription(),
		WebsiteUrl:  req.GetWebsiteUrl(),
	}

	blog, err := h.blogService.CreateBlog(ctx, converter.ToBlogInfoFromProto(blogInfo))

	if err != nil {
		return nil, fmt.Errorf("failed to create blog: %v", err)
	}

	if blog == nil {
		return nil, errors.New("blog should not be nil")

	}

	return &bloggerV1.CreateBlogResponse{
		Blog: converter.ToBlogProtoFromService(blog),
	}, nil
}

func (h *BlogHandler) GetBlog(ctx context.Context, req *bloggerV1.GetBlogRequest) (*bloggerV1.GetBlogResponse, error) {
	blog, err := h.blogService.GetBlog(ctx, req.GetId())

	if err != nil {
		return nil, fmt.Errorf("failed to get blog: %v", err)
	}

	return &bloggerV1.GetBlogResponse{
		Blog: converter.ToBlogProtoFromService(blog),
	}, nil
}

func (h *BlogHandler) UpdateBlog(ctx context.Context, req *bloggerV1.UpdateBlogRequest) (*emptypb.Empty, error) {

	err := h.blogService.UpdateBlog(ctx, req.GetId(), converter.ToBlogInfoFromProtoUpdate(req))

	if err != nil {
		return nil, fmt.Errorf("failed to update blog: %v", err)
	}

	return &emptypb.Empty{}, nil
}

func (h *BlogHandler) DeleteBlog(ctx context.Context, req *bloggerV1.DeleteBlogRequest) (*emptypb.Empty, error) {

	err := h.blogService.DeleteBlog(ctx, req.GetId())

	if err != nil {
		return nil, fmt.Errorf("failed to delete blog: %v", err)
	}

	return &emptypb.Empty{}, nil
}

func (h *BlogHandler) ListBlogs(ctx context.Context, req *bloggerV1.ListBlogsRequest) (*bloggerV1.ListBlogsResponse, error) {
	blogs, err := h.blogService.ListBlogs(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to list blogs: %v", err)
	}

	return &bloggerV1.ListBlogsResponse{
		Blog: converter.ToBlogProtoListFromService(blogs),
	}, nil
}
