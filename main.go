package main

import (
	"github.com/gin-gonic/gin"
	"github.com/warnakulasuriya-fds-e23/biometric-orchestration-go-server/controller"
	"github.com/warnakulasuriya-fds-e23/biometric-orchestration-go-server/controller/fingerprintcontroller"
	"github.com/warnakulasuriya-fds-e23/biometric-orchestration-go-server/controller/gallerycontroller"
	"github.com/warnakulasuriya-fds-e23/biometric-orchestration-go-server/initializer"
)

func main() {
	router := gin.Default()
	sdkptr := initializer.Initializer()

	finController := fingerprintcontroller.NewFingerprintController(sdkptr)
	galController := gallerycontroller.NewGalleryController(sdkptr)
	router.GET("/api/test", controller.GiveTestResponse)

	router.GET("/api/gallery/update-images-dir", galController.UpdateImagesDir)
	router.GET("/api/gallery/update-cbor-dir", galController.UpdateCborDir)

	router.GET("/api/fingerprint", finController.GetStatus)
	router.POST("/api/fingerprint/match", finController.MatchTemplates)
	router.POST("/api/fingerprint/identify", finController.IdentifyTemplate)

	router.Run("localhost:4000")
}
