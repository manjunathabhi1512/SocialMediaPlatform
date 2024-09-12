package repository

import (
	"fmt"
	"log"
	"socialmediaplatform/model"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type SocialMediaRepository interface {
	CreatePost(postData model.PostData) error
	AddComment(commentData model.CommentsData, postId uuid.UUID) error
	LikePost(postId uuid.UUID) error
	DislikePost(postId uuid.UUID) error
	GetPost(postId uuid.UUID) (model.PostData, error)
}

type SocialMediaRepo struct {
	Db *gorm.DB
}

func NewSocialMediaRepository() (*SocialMediaRepo, error) {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return &SocialMediaRepo{}, err
	}

	fmt.Println("Database connection established")
	return &SocialMediaRepo{
		Db: db,
	}, nil
}

func (smr *SocialMediaRepo) CreatePost(postData model.PostData) error {
	postDb := model.PostDataToDB{
		PostID:   postData.PostId,
		Context:  postData.Context,
		Likes:    postData.Likes,
		Dislikes: postData.Dislikes,
		PostedAt: postData.PostedAt,
	}
	if err := smr.Db.Debug().Table("socialMediaPlatform.postsdata").Create(&postDb).Error; err != nil {
		fmt.Println("Unable to create a post")
		return err
	}
	return nil
}

func (smr *SocialMediaRepo) AddComment(commentData model.CommentsData, postId uuid.UUID) error {
	commentsInDb := model.CommentDataToDB{
		CommentID:   commentData.CommentId,
		Comment:     commentData.Comment,
		CommentedAt: commentData.CommentedAt,
		PostID:      postId,
	}
	if err := smr.Db.Debug().Table("socialMediaPlatform.commentsdata").Create(&commentsInDb).Error; err != nil {
		fmt.Println("Unable to comment a post")
		return err
	}
	return nil
}

func (smr *SocialMediaRepo) LikePost(postId uuid.UUID) error {
	var post model.PostDataToDB
	if err := smr.Db.Debug().Table("socialMediaPlatform.postsdata").First(&post, "postid = ?", postId).Error; err != nil {
		fmt.Println("Post not found")
		return err
	}
	post.Likes++
	if err := smr.Db.Debug().Table("socialMediaPlatform.postsdata").Save(&post).Error; err != nil {
		return err
	}
	return nil
}

func (smr *SocialMediaRepo) DislikePost(postId uuid.UUID) error {
	var post model.PostDataToDB
	if err := smr.Db.Debug().Table("socialMediaPlatform.postsdata").First(&post, "postid = ?", postId).Error; err != nil {
		fmt.Println("Post not found")
		return err
	}
	post.Dislikes++
	if err := smr.Db.Debug().Table("socialMediaPlatform.postsdata").Save(&post).Error; err != nil {
		return err
	}
	return nil
}

func (smr *SocialMediaRepo) GetPost(postId uuid.UUID) (model.PostData, error) {
	var post model.PostDataToDB
	if err := smr.Db.Debug().Table("socialMediaPlatform.postsdata").First(&post, "postid = ?", postId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return model.PostData{}, fmt.Errorf("post with id %v not found", postId)
		}
		return model.PostData{}, err
	}
	dbData := model.PostData{
		PostId: post.PostID,
		Context: post.Context,
		Likes: post.Likes,
		Dislikes: post.Dislikes,
		PostedAt: post.PostedAt,
	}
	return dbData, nil
}
