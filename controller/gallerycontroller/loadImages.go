package gallerycontroller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (controller *galleryController) loadImages(c *gin.Context) {
	go func() {
		err := controller.sdk.LoadImages()
		if err != nil {
			err = fmt.Errorf("error occured when running sdk load images method : %w", err)
			log.Printf("Stopped loading to gallery from images becasue: %s", err.Error())
			return
		}
	}()
	c.IndentedJSON(http.StatusOK, "Successfully initialized loading of images from the specified image directory into the in memory gallery")
}
