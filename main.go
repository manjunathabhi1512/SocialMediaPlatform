package main

import (
	"log"
	"socialmediaplatform/controller"
	"socialmediaplatform/repository"
	"socialmediaplatform/route"
	"socialmediaplatform/service"
)

func main() {

	socialMediaRepo, err := repository.NewSocialMediaRepository()
	if err != nil {
		log.Fatalf("Failed to initialize repository: %v", err)
	}
	socialMediaService := service.NewSocialMediaService(socialMediaRepo)
	socialMediaController := controller.NewSocialMediaController(socialMediaService)
	socialMediaRoutes := route.InitRouter(socialMediaController)
	socialMediaRoutes.Run(":4500")
}
