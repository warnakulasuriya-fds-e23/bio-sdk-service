package fingerprintcontroller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/warnakulasuriya-fds-e23/biometric-orchestration-go-server/requestobjects"
	"github.com/warnakulasuriya-fds-e23/biometric-orchestration-go-server/responseobjects"
)

func (controller *fingerprintController) identifyTemplate(c *gin.Context) {
	var reqObj requestobjects.IdentifyTemplateReqObj
	err := c.BindJSON(&reqObj)
	if err != nil {
		resObj := responseobjects.ErrorResObj{Message: "Error when running BindJSON check response body contents, " + err.Error()}
		c.IndentedJSON(http.StatusInternalServerError, resObj)
		return
	}
	// TODO: Add more error handling to sdk methods
	probeTemplate, err := controller.sdk.ParseByteArrayToTemplate(&reqObj.ProbeCbor)
	if err != nil {
		err = fmt.Errorf("error occured when parsing probe byte data: %w", err)
		resObj := responseobjects.ErrorResObj{Message: err.Error()}
		c.IndentedJSON(http.StatusInternalServerError, resObj)
		return
	}
	isMatched, discoveredId, err := controller.sdk.Identify(probeTemplate)
	if err != nil {
		err = fmt.Errorf("error occured when running sdk identify method for probe : %w", err)
		resObj := responseobjects.ErrorResObj{Message: err.Error()}
		c.IndentedJSON(http.StatusInternalServerError, resObj)
		return
	}

	resObj := responseobjects.IdentifyTemplateResObje{
		IsMatched:    isMatched,
		DiscoveredId: discoveredId,
	}
	c.IndentedJSON(http.StatusOK, resObj)
}
