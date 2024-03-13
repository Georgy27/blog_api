package converter

import (
	"database/sql"
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

func ToBlogProtoListFromService(blogs []*model.Blog) []*bloggerV1.Blog {
	var result []*bloggerV1.Blog
	for _, blog := range blogs {
		result = append(result, ToBlogProtoFromService(blog))
	}
	return result
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

func ToBlogInfoFromProtoUpdate(req *bloggerV1.UpdateBlogRequest) *model.UpdateBlogInfo {
	name := sql.NullString{String: req.GetName().GetValue()}
	if req.GetName() != nil {
		name.Valid = true
	}

	description := sql.NullString{String: req.GetDescription().GetValue()}
	if req.GetDescription() != nil {
		description.Valid = true
	}

	websiteUrl := sql.NullString{String: req.GetWebsiteUrl().GetValue()}
	if req.GetWebsiteUrl() != nil {
		websiteUrl.Valid = true
	}

	return &model.UpdateBlogInfo{
		Name:        name,
		Description: description,
		WebsiteUrl:  websiteUrl,
	}
}
