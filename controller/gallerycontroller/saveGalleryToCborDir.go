package gallerycontroller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (controller *galleryController) saveGalleryToCborDir(c *gin.Context) {
	go func() {
		err := controller.sdk.SaveGallery()
		if err != nil {
			err = fmt.Errorf("error occured when running sdk save gallery to cbor dir method : %w", err)
			log.Printf("Stopped saving gallery to cbor directory becasue: %s", err.Error())
			return
		}
	}()
	c.IndentedJSON(http.StatusOK, "Successfully initialized saving of the gallery to the cbor directory")
}
