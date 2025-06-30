package gallerycontroller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (controller *galleryController) loadCborFiles(c *gin.Context) {
	go func() {
		err := controller.sdk.LoadCborfiles()
		if err != nil {
			err = fmt.Errorf("error occured when running sdk load cbor files method : %w", err)
			log.Printf("Stopped loading to gallery from cbor files becasue: %s", err.Error())
			return
		}
	}()
	c.IndentedJSON(http.StatusOK, "Successfully initialized loading of cbor files from the specified cbor directory into the in memory gallery")

}
