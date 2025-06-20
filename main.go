package main

import (
	"github.com/gin-gonic/gin"
	"github.com/warnakulasuriya-fds-e23/biometric-orchestration-go-server/controller"
	"github.com/warnakulasuriya-fds-e23/biometric-orchestration-go-server/controller/fingerprintcontroller"
)

func main() {
	router := gin.Default()

	finController := fingerprintcontroller.NewFingerprintController()

	router.GET("/api/test", controller.GiveTestResponse)
	router.GET("/api/fingerprint", finController.GetStatus)
	router.Run("localhost:4000")
}
