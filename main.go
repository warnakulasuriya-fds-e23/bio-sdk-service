package main

import (
	"bytes"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/warnakulasuriya-fds-e23/bio-sdk-service/internal/controller"
	"github.com/warnakulasuriya-fds-e23/bio-sdk-service/internal/controller/fingerprintcontroller"
	"github.com/warnakulasuriya-fds-e23/bio-sdk-service/internal/controller/gallerycontroller"
	"github.com/warnakulasuriya-fds-e23/bio-sdk-service/internal/initializer"
)

func RequestLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var buf bytes.Buffer
		tee := io.TeeReader(c.Request.Body, &buf)
		body, _ := io.ReadAll(tee)
		c.Request.Body = io.NopCloser(&buf)
		log.Println(string(body))
		log.Println(c.Request.Header)
		c.Next()
	}
}

func main() {
	_, err := os.Stat(".env")
	if err == nil {
		log.Println("discovered .env file")
		err := godotenv.Load()
		if err != nil {
			log.Println("however failed to load .env file")
		} else {
			log.Println(".env successfully loaded")
		}
	}
	router := gin.Default()
	sdkptr := initializer.Initializer()

	finController := fingerprintcontroller.NewFingerprintController(sdkptr)
	galController := gallerycontroller.NewGalleryController(sdkptr)

	router.GET("/api/test", controller.GiveTestResponse)
	router.Use(RequestLoggerMiddleware())

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
	router.POST("/api/fingerping/authorize", finController.Authorize)
	// apparently in following way gin listens for 4000 on all network interfaces
	router.Run(":4000")
}
