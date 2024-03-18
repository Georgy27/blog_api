package converter

import (
	"github.com/Georgy27/blogger_api/internal/post/model"
	modelRepo "github.com/Georgy27/blogger_api/internal/post/repository/model"
)

func ToPostFromRepo(post *modelRepo.Post) *model.Post {
	return &model.Post{
		ID:        post.ID,
		Info:      *ToPostInfoFromRepo(post.Info),
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}
}

func ToPostsFromRepo(posts []*modelRepo.Post) []*model.Post {
	var result []*model.Post
	for _, post := range posts {
		result = append(result, ToPostFromRepo(post))
	}
	return result
}

func ToPostInfoFromRepo(postInfo modelRepo.PostInfo) *model.PostInfo {
	return &model.PostInfo{
		Title:            postInfo.Title,
		Content:          postInfo.Content,
		ShortDescription: postInfo.ShortDescription,
		BlogID:           postInfo.BlogID,
		BlogName:         postInfo.BlogName,
	}
}
