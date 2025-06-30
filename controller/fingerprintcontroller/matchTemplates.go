package fingerprintcontroller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/warnakulasuriya-fds-e23/biometric-orchestration-go-server/requestobjects"
	"github.com/warnakulasuriya-fds-e23/biometric-orchestration-go-server/responseobjects"
)

func (controller *fingerprintController) matchTemplates(c *gin.Context) {
	var reqObj requestobjects.MatchTemplatesReqObj
	err := c.BindJSON(&reqObj)
	if err != nil {
		resObj := responseobjects.ErrorResObj{Message: "Error when running BindJSON check response body contents, " + err.Error()}
		c.IndentedJSON(http.StatusInternalServerError, resObj)
		return
	}
	probeTemplate, err := controller.sdk.ParseByteArrayToTemplate(&reqObj.ProbeCbor)
	if err != nil {
		err = fmt.Errorf("error occured when parsing probe byte data: %w", err)
		resObj := responseobjects.ErrorResObj{Message: err.Error()}
		c.IndentedJSON(http.StatusInternalServerError, resObj)
		return
	}
	candidateTemplate, err := controller.sdk.ParseByteArrayToTemplate(&reqObj.CandidateCbor)
	if err != nil {
		err = fmt.Errorf("error occured when parsing candidate byte data: %w", err)
		resObj := responseobjects.ErrorResObj{Message: err.Error()}
		c.IndentedJSON(http.StatusInternalServerError, resObj)
		return
	}
	isMatch, err := controller.sdk.Match(probeTemplate, candidateTemplate)
	if err != nil {
		err = fmt.Errorf("error occured when running match method of sdk: %w", err)
		resObj := responseobjects.ErrorResObj{Message: err.Error()}
		c.IndentedJSON(http.StatusInternalServerError, resObj)
		return
	}
	resObj := responseobjects.MatchTemplatesResObj{IsMatch: isMatch}
	c.IndentedJSON(http.StatusOK, resObj)
}
