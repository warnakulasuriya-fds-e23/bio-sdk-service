package fingerprintcontroller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/warnakulasuriya-fds-e23/bio-sdk-service/internal/requestobjects"
	"github.com/warnakulasuriya-fds-e23/bio-sdk-service/internal/responseobjects"
)

func (controller *fingerprintController) enrollTemplate(c *gin.Context) {
	var reqObj requestobjects.EnrollTemplateReqObj
	err := c.BindJSON(&reqObj)
	if err != nil {
		resObj := responseobjects.ErrorResObj{Message: "Error when running BindJSON check response body contents, " + err.Error()}
		c.IndentedJSON(http.StatusInternalServerError, resObj)
		return
	}
	template, err := controller.sdk.ParseByteArrayToTemplate(&reqObj.Data)
	if err != nil {
		err = fmt.Errorf("error occured when parsing new entry byte data: %w", err)
		resObj := responseobjects.ErrorResObj{Message: err.Error()}
		c.IndentedJSON(http.StatusInternalServerError, resObj)
		return
	}
	err = controller.sdk.Enroll(template, reqObj.Id)
	if err != nil {
		err = fmt.Errorf("error occured run Enroll of SDKCore: %w", err)
		resObj := responseobjects.ErrorResObj{Message: err.Error()}
		c.IndentedJSON(http.StatusInternalServerError, resObj)
		return
	}
	resObj := responseobjects.EnrollTemplateResObj{Message: "Enrolled " + reqObj.Id + " successfully"}
	c.IndentedJSON(http.StatusOK, resObj)

}
