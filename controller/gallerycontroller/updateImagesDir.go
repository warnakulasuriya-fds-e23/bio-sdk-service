package gallerycontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/warnakulasuriya-fds-e23/biometric-orchestration-go-server/requestobjects"
	"github.com/warnakulasuriya-fds-e23/biometric-orchestration-go-server/responseobjects"
)

func (controller *galleryController) updateImagesDir(c *gin.Context) {
	var reqObj requestobjects.UpdateImagesDirReqObj
	err := c.BindJSON(&reqObj)
	if err != nil {
		res := responseobjects.ErrorResObj{Message: err.Error()}
		c.IndentedJSON(http.StatusInternalServerError, res)
	}
	message, err := controller.sdk.UpdateImageDir(reqObj.PathString)
	if err != nil {
		res := responseobjects.ErrorResObj{Message: err.Error()}
		c.IndentedJSON(http.StatusInternalServerError, res)
	}
	res := responseobjects.UpdateImagesDirResObj{Message: message}
	c.IndentedJSON(http.StatusInternalServerError, res)
}
