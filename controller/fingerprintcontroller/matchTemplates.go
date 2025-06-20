package fingerprintcontroller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/warnakulasuriya-fds-e23/biometric-orchestration-go-server/requestobjects"
	"github.com/warnakulasuriya-fds-e23/biometric-orchestration-go-server/responseobjects"
)

func (controller *fingerprintController) matchTemplates(c *gin.Context) {
	var reqObj requestobjects.MatchTemplatesReqObj
	err := c.BindJSON(&reqObj)
	if err != nil {
		log.Fatal(err.Error())
	}
	probeTemplate := controller.sdk.ParseByteArrayToTemplate(&reqObj.ProbeCbor)
	candidateTemplate := controller.sdk.ParseByteArrayToTemplate(&reqObj.CandidateCbor)
	isMatch := controller.sdk.Match(probeTemplate, candidateTemplate)
	if isMatch {
		res := responseobjects.MatchTemplatesResObj{Status: "match"}
		c.IndentedJSON(http.StatusOK, res)
	} else {
		res := responseobjects.MatchTemplatesResObj{Status: "no match"}
		c.IndentedJSON(http.StatusOK, res)
	}
}
