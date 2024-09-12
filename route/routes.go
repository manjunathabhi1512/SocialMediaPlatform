package route

import (
	"socialmediaplatform/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter(controller *controller.SocialMediaController) *gin.Engine {
	r := gin.Default()
	r.POST("/create", controller.CreatePostController)
	r.POST("/:id/addComment", controller.AddCommentController)
	r.POST("/:id/likePost", controller.LikePostController)
	r.POST("/:id/dislikePost", controller.DislikePostController)
	r.GET("/:id/getPost", controller.GetPostController)
	r.GET("/:id/share", controller.SharePostController)

	return r
}
