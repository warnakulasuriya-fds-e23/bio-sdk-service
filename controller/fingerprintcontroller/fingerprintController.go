package fingerprintcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/warnakulasuriya-fds-e23/fingerprint-go-sdk/core"
)

type fingerprintController struct {
	sdk *core.SDKCore
}

func NewFingerprintController(sdkptr *core.SDKCore) (controller *fingerprintController) {
	controller = &fingerprintController{sdk: sdkptr}
	return
}

func (controller *fingerprintController) GetStatus(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "runnig")
}
func (controller *fingerprintController) MatchTemplates(c *gin.Context) {
	controller.matchTemplates(c)
}
func (controller *fingerprintController) IdentifyTemplate(c *gin.Context) {
	controller.identifyTemplate(c)
}
