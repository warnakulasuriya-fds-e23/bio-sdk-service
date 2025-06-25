package gallerycontroller

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/warnakulasuriya-fds-e23/biometric-orchestration-go-server/requestobjects"
	"github.com/warnakulasuriya-fds-e23/biometric-orchestration-go-server/responseobjects"
)

func (controller *galleryController) updateImagesDir(c *gin.Context) {
	var reqObj requestobjects.UpdateImagesDirReqObj
	err := c.BindJSON(&reqObj)
	if err != nil {
		resObj := responseobjects.ErrorResObj{Message: "Error when running BindJSON check response body contents, " + err.Error()}
		c.IndentedJSON(http.StatusInternalServerError, resObj)
		return
	}
	message, err := controller.sdk.UpdateImageDir(reqObj.PathString)
	if err != nil {
		workingDir, errGetwd := os.Getwd()
		if errGetwd != nil {
			resObj := responseobjects.ErrorResObj{Message: errGetwd.Error() + " + " + err.Error()}
			c.IndentedJSON(http.StatusInternalServerError, resObj)
			return
		}
		resObj := responseobjects.ErrorResObj{Message: "Current working directory: " + workingDir + " ," + err.Error()}
		c.IndentedJSON(http.StatusInternalServerError, resObj)
		return
	}
	resObj := responseobjects.UpdateImagesDirResObj{Message: message}
	c.IndentedJSON(http.StatusInternalServerError, resObj)
}
