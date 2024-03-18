package converter

import (
	"database/sql"
	"github.com/Georgy27/blogger_api/internal/post/model"
	postV1 "github.com/Georgy27/blogger_api/pkg/proto/post_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToPostProtoFromService(post *model.Post) *postV1.Post {
	var updatedAt *timestamppb.Timestamp

	if post.UpdatedAt.Valid {
		updatedAt = timestamppb.New(post.UpdatedAt.Time)
	}

	return &postV1.Post{
		Id:        post.ID,
		Info:      ToPostInfoProtoFromService(post.Info),
		CreatedAt: timestamppb.New(post.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

func ToPostProtoListFromService(posts []*model.Post) []*postV1.Post {
	var result []*postV1.Post
	for _, post := range posts {
		result = append(result, ToPostProtoFromService(post))
	}
	return result
}

func ToPostInfoProtoFromService(postInfo model.PostInfo) *postV1.PostInfo {
	return &postV1.PostInfo{
		Title:            postInfo.Title,
		Content:          postInfo.Content,
		ShortDescription: postInfo.ShortDescription,
		BlogId:           postInfo.BlogID,
		BlogName:         postInfo.BlogName,
	}
}

func ToPostInfoFromProto(postInfo *postV1.PostInfo) *model.PostInfo {
	return &model.PostInfo{
		Title:            postInfo.Title,
		Content:          postInfo.Content,
		ShortDescription: postInfo.ShortDescription,
		BlogID:           postInfo.BlogId,
		BlogName:         "",
	}
}

func ToPostInfoFromProtoUpdate(req *postV1.UpdatePostRequest) *model.UpdatePostInfo {
	title := sql.NullString{String: req.GetTitle().GetValue()}
	if req.GetTitle() != nil {
		title.Valid = true
	}

	content := sql.NullString{String: req.GetContent().GetValue()}
	if req.GetContent() != nil {
		content.Valid = true
	}

	shortDescription := sql.NullString{String: req.GetShortDescription().GetValue()}
	if req.GetShortDescription() != nil {
		shortDescription.Valid = true
	}

	return &model.UpdatePostInfo{
		BlogID:           req.GetBlogId(),
		Title:            title,
		ShortDescription: shortDescription,
		Content:          content,
	}
}
