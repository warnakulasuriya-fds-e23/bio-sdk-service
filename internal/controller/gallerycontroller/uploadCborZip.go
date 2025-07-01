package gallerycontroller

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/warnakulasuriya-fds-e23/bio-sdk-service/internal/configtomlreader"
	"github.com/warnakulasuriya-fds-e23/bio-sdk-service/internal/responseobjects"
	"github.com/warnakulasuriya-fds-e23/bio-sdk-service/internal/serverutilis"
)

func (controller *galleryController) uploadCborZip(c *gin.Context) {
	// single file
	file, _ := c.FormFile("file")
	if filepath.Ext(file.Filename) != ".zip" {
		resObj := responseobjects.ErrorResObj{Message: "error only zip files can be uploaded to this endpoint "}
		c.IndentedJSON(http.StatusInternalServerError, resObj)
		return
	}

	config := configtomlreader.ConfigTomlReader()
	zipfilePath := filepath.Join(config.StorageVolume, "/uploadedcbors", file.Filename)
	// Upload the file to specific dst.
	c.SaveUploadedFile(file, zipfilePath)

	err := serverutilis.UnzipCborDirZip(zipfilePath)
	if err != nil {
		resObj := responseobjects.ErrorResObj{Message: "error while trying to unzip, however file was uploaded : " + err.Error()}
		c.IndentedJSON(http.StatusInternalServerError, resObj)
		return
	}
	err = os.Remove(zipfilePath)
	if err != nil {
		resObj := responseobjects.ErrorResObj{Message: "error while trying to remove zip file after it was unzipped"}
		c.IndentedJSON(http.StatusInternalServerError, resObj)
		return
	}
	base := filepath.Base(zipfilePath)
	parent := filepath.Dir(zipfilePath)
	folderName := strings.TrimSuffix(base, filepath.Ext(base))
	folderPath := filepath.Join(parent, folderName)
	c.String(http.StatusOK, fmt.Sprintf("%s uploaded, was extracted to %s", file.Filename, folderPath))
}
