package main

import (
	"socialmediaplatform/controller"
	"socialmediaplatform/route"
	"socialmediaplatform/service"
)

func main() {

	socialMediaService := service.NewSocialMediaService()
	socialMediaController := controller.NewSocialMediaController(socialMediaService)
	socialMediaRoutes := route.InitRouter(socialMediaController)
	socialMediaRoutes.Run(":4500")
}
