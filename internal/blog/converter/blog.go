package converter

import (
	"github.com/Georgy27/blogger_api/internal/blog/model"
	bloggerV1 "github.com/Georgy27/blogger_api/pkg/blogger_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToBlogProtoFromService(blog *model.Blog) *bloggerV1.Blog {
	var updatedAt *timestamppb.Timestamp

	if blog.UpdatedAt.Valid {
		updatedAt = timestamppb.New(blog.UpdatedAt.Time)
	}

	return &bloggerV1.Blog{
		Id:        blog.ID,
		Info:      ToBlogInfoProtoFromService(blog.Info),
		CreatedAt: timestamppb.New(blog.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

func ToBlogInfoProtoFromService(blogInfo model.BlogInfo) *bloggerV1.BlogInfo {
	return &bloggerV1.BlogInfo{
		Name:        blogInfo.Name,
		Description: blogInfo.Description,
		WebsiteUrl:  blogInfo.WebsiteUrl,
	}
}

func ToBlogInfoFromProto(blogInfo *bloggerV1.BlogInfo) *model.BlogInfo {
	return &model.BlogInfo{
		Name:        blogInfo.Name,
		Description: blogInfo.Description,
		WebsiteUrl:  blogInfo.WebsiteUrl,
	}
}
