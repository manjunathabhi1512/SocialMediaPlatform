package service

import (
	"fmt"
	"socialmediaplatform/model"
	"socialmediaplatform/repository"
	"time"

	"github.com/google/uuid"
)

type SocialMediaService struct {
	PostData map[uuid.UUID]model.PostData
	Repo     repository.SocialMediaRepository
}

// DB Flag for switching between db or local storage using maps & structs
var DbFlag bool = true

func NewSocialMediaService(Repo repository.SocialMediaRepository) *SocialMediaService {
	return &SocialMediaService{
		PostData: make(map[uuid.UUID]model.PostData),
		Repo: Repo,
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
	if DbFlag == true {
		if err := sms.Repo.CreatePost(postData); err != nil{
			return uuid.Nil, err
		}
		return id, nil
	}
	sms.PostData[postData.PostId] = postData
	return id, nil
}

func (sms *SocialMediaService) AddCommentService(postId uuid.UUID, comment string) error {
	
	commentData := model.CommentsData{
		CommentId:   uuid.New(),
		Comment:     comment,
		CommentedAt: time.Now(),
	}
	
	if DbFlag == true {
		if err := sms.Repo.AddComment(commentData, postId); err != nil{
			return err
		}
		return nil
	}
	post, exist := sms.PostData[postId]
	if !exist {
		fmt.Println("No posts found")
		return fmt.Errorf("No posts Found")
	}
	post.Comments = append(post.Comments, commentData)
	sms.PostData[postId] = post
	return nil
}

func (sms *SocialMediaService) LikePostService(postId uuid.UUID) error {
	if DbFlag == true {
		if err := sms.Repo.LikePost(postId); err != nil{
			return err
		}
		return nil
	}
	post, exist := sms.PostData[postId]
	if !exist {
		return fmt.Errorf("No posts Found")
	}
	post.Likes++
	sms.PostData[postId] = post
	return nil
}

func (sms *SocialMediaService) DislikePostService(postId uuid.UUID) error {
	if DbFlag == true {
		if err := sms.Repo.DislikePost(postId); err != nil{
			return err
		}
		return nil
	}
	post, exist := sms.PostData[postId]
	if !exist {
		return fmt.Errorf("No posts Found")
	}
	post.Dislikes++
	sms.PostData[postId] = post
	return nil
}

func (sms *SocialMediaService) GetPostService(postId uuid.UUID) (model.PostData, error) {
	if DbFlag == true{
		post, err := sms.Repo.GetPost(postId); 
		if err!=nil{
			return model.PostData{}, err
		}
		return post, nil
	}
	post, exist := sms.PostData[postId]
	if !exist {
		return post, fmt.Errorf("No posts Found")
	}
	
	return post, nil
}
