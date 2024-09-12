package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"socialmediaplatform/model"
	"socialmediaplatform/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SocialMediaController struct {
	SocialMediaService *service.SocialMediaService
}

func NewSocialMediaController(service *service.SocialMediaService) *SocialMediaController {
	return &SocialMediaController{
		SocialMediaService: service,
	}
}

func (smc *SocialMediaController) CreatePostController(r *gin.Context) {
	var payloadData model.PostPayload
	body, err := io.ReadAll(r.Request.Body)
	if err != nil {
		fmt.Println("Unable to Read the data")
		return
	}
	if err := json.Unmarshal(body, &payloadData); err != nil {
		fmt.Println("Unable to Unmarshal the data")
		return
	}
	postId, err := smc.SocialMediaService.CreatePostService(payloadData.ContextData)
	if err != nil {
		fmt.Println("Failed to create the post")
		return
	}
	r.JSON(http.StatusOK, gin.H{
		"message": "Post created successfully",
		"Post_Id": postId,
	})
}

func (smc *SocialMediaController) AddCommentController(r *gin.Context) {
	postId, err := uuid.Parse(r.Param("id"))
	if err != nil {
		return
	}
	var payloadData model.PostPayload

	body, err := io.ReadAll(r.Request.Body)
	if err != nil {
		fmt.Println("Unable to Read the data")
		return
	}
	if err := json.Unmarshal(body, &payloadData); err != nil {
		fmt.Println("Unable to Unmarshal the data")
		return
	}

	if err := smc.SocialMediaService.AddCommentService(postId, payloadData.ContextData); err != nil {
		fmt.Println("Failed to comment the post")
		return
	}
	r.JSON(http.StatusOK, gin.H{
		"message": "Commented successfully",
		"data":    payloadData,
	})
}

func (smc *SocialMediaController) LikePostController(r *gin.Context) {
	postId, err := uuid.Parse(r.Param("id"))
	if err != nil {
		return
	}
	if err := smc.SocialMediaService.LikePostService(postId); err != nil {
		fmt.Println("Failed to Like the post")
		return
	}
	r.JSON(http.StatusOK, gin.H{
		"message": "Liked successfully",
	})
}

func (smc *SocialMediaController) DislikePostController(r *gin.Context) {
	postId, err := uuid.Parse(r.Param("id"))
	if err != nil {
		return
	}
	if err := smc.SocialMediaService.DislikePostService(postId); err != nil {
		fmt.Println("Failed to Dislike the post")
		return
	}
	r.JSON(http.StatusOK, gin.H{
		"message": "Disliked successfully",
	})
}

func (smc *SocialMediaController) GetPostController(r *gin.Context) {
	postId, err := uuid.Parse(r.Param("id"))
	if err != nil {
		return
	}
	postData, err := smc.SocialMediaService.GetPostService(postId)
	if err != nil {
		fmt.Println("Failed to fetch the post")
		return
	}
	r.JSON(http.StatusOK, gin.H{
		"message": "fetched post successfully",
		"data":    postData,
	})
}

func (smc *SocialMediaController) SharePostController(r *gin.Context) {
	shareableLink := fmt.Sprintf("https://socialmediaplatform.com/post/%s", r.Param("id"))
	r.JSON(http.StatusOK, gin.H{
		"message": "Unique Link Generated successfully",
		"data":    shareableLink,
	})
}
