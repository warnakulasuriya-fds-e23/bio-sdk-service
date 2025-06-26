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

	router.GET("/api/gallery/get-images-dir", galController.GetImagesDir)
	router.GET("/api/gallery/get-cbor-dir", galController.GetCborDir)
	router.POST("/api/gallery/update-images-dir", galController.UpdateImagesDir)
	router.POST("/api/gallery/update-cbor-dir", galController.UpdateCborDir)
	router.GET("/api/gallery/load-images", galController.LoadImages)
	router.GET("/api/gallery/load-cbor-files", galController.LoadCborFiles)
	router.GET("/api/gallery/save-gallery", galController.SaveGalleryToCborDir)
	router.POST("/api/gallery/upload-cbor-zip", galController.UploadCborDirZip)

	router.GET("/api/fingerprint", finController.GetStatus)
	router.POST("/api/fingerprint/match", finController.MatchTemplates)
	router.POST("/api/fingerprint/identify", finController.IdentifyTemplate)
	router.POST("/api/fingerprint/enroll", finController.EnrollTemplate)

	// apparently in following way gin listens for 4000 on all network interfaces
	router.Run(":4000")
}
