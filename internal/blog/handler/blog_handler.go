package handler

import (
	"context"
	"github.com/Georgy27/blogger_api/internal/blog/converter"
	blogService "github.com/Georgy27/blogger_api/internal/blog/service"
	"github.com/Georgy27/blogger_api/internal/blog/validation"
	commonErrors "github.com/Georgy27/blogger_api/pkg/errors"
	"github.com/Georgy27/blogger_api/pkg/errors/codes"
	"github.com/Georgy27/blogger_api/pkg/errors/validate"
	bloggerV1 "github.com/Georgy27/blogger_api/pkg/proto/blogger_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	create = "create"
	update = "update"
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
	err := validate.Validate(
		ctx,
		validation.ValidateName(req.GetName(), create),
		validation.ValidateDescription(req.GetDescription(), create),
		validation.ValidateWebsiteUrl(req.GetWebsiteUrl(), create),
	)

	if err != nil {
		return nil, err
	}

	blogInfo := &bloggerV1.BlogInfo{
		Name:        req.GetName(),
		Description: req.GetDescription(),
		WebsiteUrl:  req.GetWebsiteUrl(),
	}

	blog, err := h.blogService.CreateBlog(ctx, converter.ToBlogInfoFromProto(blogInfo))

	if err != nil {
		return nil, commonErrors.NewCommonError("failed to create blog", codes.Internal)
	}

	return &bloggerV1.CreateBlogResponse{
		Blog: converter.ToBlogProtoFromService(blog),
	}, nil
}

func (h *BlogHandler) GetBlog(ctx context.Context, req *bloggerV1.GetBlogRequest) (*bloggerV1.GetBlogResponse, error) {
	err := validate.Validate(
		ctx,
		validation.ValidateID(req.GetId()),
	)

	if err != nil {
		return nil, err
	}

	blog, err := h.blogService.GetBlog(ctx, req.GetId())

	if err != nil {
		return nil, commonErrors.NewCommonError("failed to get blog; blog not found", codes.NotFound)
	}

	return &bloggerV1.GetBlogResponse{
		Blog: converter.ToBlogProtoFromService(blog),
	}, nil
}

func (h *BlogHandler) UpdateBlog(ctx context.Context, req *bloggerV1.UpdateBlogRequest) (*emptypb.Empty, error) {
	err := validate.Validate(
		ctx,
		validation.ValidateName(req.GetName().GetValue(), update),
		validation.ValidateDescription(req.GetDescription().GetValue(), update),
		validation.ValidateWebsiteUrl(req.GetWebsiteUrl().GetValue(), update),
	)

	if err != nil {
		return nil, err
	}

	err = h.blogService.UpdateBlog(ctx, req.GetId(), converter.ToBlogInfoFromProtoUpdate(req))

	if err != nil {
		return nil, commonErrors.NewCommonError("failed to update blog", codes.NotFound)
	}

	return &emptypb.Empty{}, nil
}

func (h *BlogHandler) DeleteBlog(ctx context.Context, req *bloggerV1.DeleteBlogRequest) (*emptypb.Empty, error) {
	err := validate.Validate(
		ctx,
		validation.ValidateID(req.GetId()),
	)

	if err != nil {
		return nil, err
	}

	err = h.blogService.DeleteBlog(ctx, req.GetId())

	if err != nil {
		return nil, commonErrors.NewCommonError("failed to delete blog; blog not found", codes.NotFound)
	}

	return &emptypb.Empty{}, nil
}

func (h *BlogHandler) ListBlogs(ctx context.Context, req *bloggerV1.ListBlogsRequest) (*bloggerV1.ListBlogsResponse, error) {
	blogs, err := h.blogService.ListBlogs(ctx)

	if err != nil {
		return nil, commonErrors.NewCommonError("failed to list blogs", codes.Internal)
	}

	return &bloggerV1.ListBlogsResponse{
		Blog: converter.ToBlogProtoListFromService(blogs),
	}, nil
}
