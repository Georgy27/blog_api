package handler

import (
	"context"
	"github.com/Georgy27/blogger_api/internal/post/converter"
	postService "github.com/Georgy27/blogger_api/internal/post/service"
	"github.com/Georgy27/blogger_api/internal/post/validation"
	"github.com/Georgy27/blogger_api/pkg/errors"
	commonErrors "github.com/Georgy27/blogger_api/pkg/errors"
	"github.com/Georgy27/blogger_api/pkg/errors/codes"
	"github.com/Georgy27/blogger_api/pkg/errors/validate"
	postV1 "github.com/Georgy27/blogger_api/pkg/proto/post_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type PostHandler struct {
	postV1.UnimplementedPostV1Server
	postService postService.PostService
}

func NewPostHandler(postService postService.PostService) *PostHandler {
	return &PostHandler{
		postService: postService,
	}
}

func (h *PostHandler) CreatePost(ctx context.Context, req *postV1.CreatePostRequest) (*postV1.CreatePostResponse, error) {
	err := validate.Validate(
		ctx,
		validation.ValidateTitle(req.GetTitle()),
		validation.ValidateContent(req.GetContent()),
		validation.ValidateShortDescription(req.GetShortDescription()),
	)

	if err != nil {
		return nil, err
	}

	postInfo := &postV1.PostInfo{
		Title:            req.GetTitle(),
		Content:          req.GetContent(),
		ShortDescription: req.GetShortDescription(),
		BlogId:           req.GetBlogId(),
	}

	post, err := h.postService.CreatePost(ctx, converter.ToPostInfoFromProto(postInfo))

	if err != nil {
		if errors.IsCommonError(err) {
			return nil, commonErrors.NewCommonError(err.Error(), codes.NotFound)
		}
		return nil, commonErrors.NewCommonError(err.Error(), codes.Internal)
	}

	return &postV1.CreatePostResponse{
		Post: converter.ToPostProtoFromService(post),
	}, nil

}

func (h *PostHandler) GetPost(ctx context.Context, req *postV1.GetPostRequest) (*postV1.GetPostResponse, error) {
	return nil, nil
}

func (h *PostHandler) UpdatePost(ctx context.Context, req *postV1.UpdatePostRequest) (*emptypb.Empty, error) {
	return nil, nil
}

func (h *PostHandler) DeletePost(ctx context.Context, req *postV1.DeletePostRequest) (*emptypb.Empty, error) {
	return nil, nil
}

func (h *PostHandler) ListPosts(ctx context.Context, req *postV1.ListPostsRequest) (*postV1.ListPostsResponse, error) {
	return nil, nil
}
