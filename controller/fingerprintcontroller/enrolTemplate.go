package fingerprintcontroller

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/warnakulasuriya-fds-e23/biometric-orchestration-go-server/requestobjects"
	"github.com/warnakulasuriya-fds-e23/biometric-orchestration-go-server/responseobjects"
)

func (controller *fingerprintController) enrollTemplate(c *gin.Context) {
	bodyBytes, err1 := io.ReadAll(c.Request.Body)
	if err1 != nil {
		log.Printf("Error reading request body: %v", err1)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
		return
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	requestBodyString := string(bodyBytes)
	log.Printf("Incoming JSON Request Body: %s\n", requestBodyString)

	var reqObj requestobjects.EnrollTemplateReqObj
	err := c.BindJSON(&reqObj)
	if err != nil {
		resObj := responseobjects.ErrorResObj{Message: "Error when running BindJSON check response body contents, " + err.Error()}
		c.IndentedJSON(http.StatusInternalServerError, resObj)
	}
	template, err := controller.sdk.ParseByteArrayToTemplate(&reqObj.Data)
	if err != nil {
		err = fmt.Errorf("error occured when parsing new entry byte data: %w", err)
		resObj := responseobjects.ErrorResObj{Message: err.Error()}
		c.IndentedJSON(http.StatusInternalServerError, resObj)
	}
	err = controller.sdk.Enroll(template, reqObj.Id)
	if err != nil {
		err = fmt.Errorf("error occured run Enroll of SDKCore: %w", err)
		resObj := responseobjects.ErrorResObj{Message: err.Error()}
		c.IndentedJSON(http.StatusInternalServerError, resObj)
	}
	resObj := responseobjects.EnrollTemplateResObj{Message: "Enrolled " + reqObj.Id + " successfully"}
	c.IndentedJSON(http.StatusOK, resObj)

}
