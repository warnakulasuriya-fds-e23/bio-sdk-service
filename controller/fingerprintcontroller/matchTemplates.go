package fingerprintcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/warnakulasuriya-fds-e23/biometric-orchestration-go-server/requestobjects"
	"github.com/warnakulasuriya-fds-e23/biometric-orchestration-go-server/responseobjects"
)

func (controller *fingerprintController) matchTemplates(c *gin.Context) {
	var reqObj requestobjects.MatchTemplatesReqObj
	err := c.BindJSON(&reqObj)
	if err != nil {
		resObj := responseobjects.ErrorResObj{Message: err.Error()}
		c.IndentedJSON(http.StatusInternalServerError, resObj)
	}
	// TODO: Add more error handling to sdk methods
	probeTemplate := controller.sdk.ParseByteArrayToTemplate(&reqObj.ProbeCbor)
	candidateTemplate := controller.sdk.ParseByteArrayToTemplate(&reqObj.CandidateCbor)
	isMatch := controller.sdk.Match(probeTemplate, candidateTemplate)
	resObj := responseobjects.MatchTemplatesResObj{IsMatch: isMatch}
	c.IndentedJSON(http.StatusOK, resObj)
}
