package fingerprintcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/warnakulasuriya-fds-e23/biometric-orchestration-go-server/requestobjects"
	"github.com/warnakulasuriya-fds-e23/biometric-orchestration-go-server/responseobjects"
)

func (controller *fingerprintController) identifyTemplate(c *gin.Context) {
	var reqObj requestobjects.IdentifyTemplateReqObj
	err := c.BindJSON(&reqObj)
	if err != nil {
		res := responseobjects.ErrorResObj{Message: err.Error()}
		c.IndentedJSON(http.StatusInternalServerError, res)
	}

	// TODO: Add more error handling to sdk methods
	probeTemplate := controller.sdk.ParseByteArrayToTemplate(&reqObj.ProbeCbor)
	isMatched, discoveredId := controller.sdk.Identify(probeTemplate)
	resObj := responseobjects.IdentifyTemplateResObje{
		IsMatched:    isMatched,
		DiscoveredId: discoveredId,
	}
	c.IndentedJSON(http.StatusOK, resObj)
}
