package service

import (
	"fmt"
	"socialmediaplatform/model"
	"time"

	"github.com/google/uuid"
)

type SocialMediaService struct {
	PostData map[uuid.UUID]model.PostData
}

func NewSocialMediaService() *SocialMediaService {
	return &SocialMediaService{
		PostData: make(map[uuid.UUID]model.PostData),
	}
}

func (sms *SocialMediaService) CreatePostService(content string) (uuid.UUID, error) {
	id := uuid.New()

	postData := model.PostData{
		PostId:   id,
		Context:  content,
		PostedAt: time.Now(),
		Comments: []model.CommentsData{},
	}
	sms.PostData[postData.PostId] = postData
	return id, nil
}

func (sms *SocialMediaService) AddCommentService(postId uuid.UUID, comment string) error{

	post, exist := sms.PostData[postId]
	if !exist{
		fmt.Println("No posts found")
		return fmt.Errorf("No posts Found")
	}
	commentData := model.CommentsData{
		CommentId : uuid.New(),
		Comment: comment,
		CommentedAt: time.Now(),
	}
	post.Comments = append(post.Comments, commentData)
	sms.PostData[postId] = post
	return nil
}

func (sms *SocialMediaService) LikePostService(postId uuid.UUID) error{
	post, exist := sms.PostData[postId]
	if !exist{
		return fmt.Errorf("No posts Found")
	}
	post.Likes++
	sms.PostData[postId] = post
	return nil
}

func (sms *SocialMediaService) DislikePostService(postId uuid.UUID) error{
	post, exist := sms.PostData[postId]
	if !exist{
		return fmt.Errorf("No posts Found")
	}
	post.Dislikes++
	sms.PostData[postId] = post
	return nil
}

func (sms *SocialMediaService) GetPostService(postId uuid.UUID) (model.PostData, error){
	post, exist := sms.PostData[postId]
	if !exist{
		return post, fmt.Errorf("No posts Found")
	} 
	return post, nil
}
