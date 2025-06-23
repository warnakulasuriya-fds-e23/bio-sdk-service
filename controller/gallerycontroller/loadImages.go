package gallerycontroller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/warnakulasuriya-fds-e23/biometric-orchestration-go-server/responseobjects"
)

func (controller *galleryController) loadImages(c *gin.Context) {
	err := controller.sdk.LoadImages()
	if err != nil {
		err = fmt.Errorf("error occured when running sdk load images method : %w", err)
		resObj := responseobjects.ErrorResObj{Message: err.Error()}
		c.IndentedJSON(http.StatusInternalServerError, resObj)
	}

}
