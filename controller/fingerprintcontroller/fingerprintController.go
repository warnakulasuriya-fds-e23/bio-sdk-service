package fingerprintcontroller

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"github.com/warnakulasuriya-fds-e23/biometric-orchestration-go-server/configuration"
	"github.com/warnakulasuriya-fds-e23/fingerprint-go-sdk/core"
)

type fingerprintController struct {
	sdk *core.SDKCore
}

func NewFingerprintController() *fingerprintController {
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err.Error())
	}

	tomlpath := filepath.Join(workingDir, "config.toml")
	var config configuration.Configuration
	toml.DecodeFile(tomlpath, &config)
	s, err := core.NewSDKCore(config.ImagesDir, config.CborDir)
	if err != nil {
		log.Fatal(err.Error())
	}

	controller := &fingerprintController{sdk: s}

	return controller
}

func (controller *fingerprintController) GetStatus(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "runnig")
}
