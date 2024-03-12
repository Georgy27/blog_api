package converter

import (
	"github.com/Georgy27/blogger_api/internal/blog/model"
	modelRepo "github.com/Georgy27/blogger_api/internal/blog/repository/model"
)

func ToBlogFromRepo(blog *modelRepo.Blog) *model.Blog {
	return &model.Blog{
		ID:        blog.ID,
		Info:      *ToBlogInfoFromRepo(blog.Info),
		CreatedAt: blog.CreatedAt,
		UpdatedAt: blog.UpdatedAt,
	}
}

func ToBlogInfoFromRepo(blogInfo modelRepo.BlogInfo) *model.BlogInfo {
	return &model.BlogInfo{
		Name:        blogInfo.Name,
		Description: blogInfo.Description,
		WebsiteUrl:  blogInfo.WebsiteUrl,
	}
}
