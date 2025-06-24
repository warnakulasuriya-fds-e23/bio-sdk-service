package gallerycontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/warnakulasuriya-fds-e23/fingerprint-go-sdk/core"
)

type galleryController struct {
	sdk *core.SDKCore
}

func NewGalleryController(sdkptr *core.SDKCore) *galleryController {
	return &galleryController{sdk: sdkptr}
}
func (controller *galleryController) UpdateImagesDir(c *gin.Context) {
	controller.updateImagesDir(c)
}
func (controller *galleryController) GetImagesDir(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, controller.sdk.GetImagesDir())
}
func (controller *galleryController) UpdateCborDir(c *gin.Context) {
	controller.updateCborDir(c)
}
func (controller *galleryController) GetCborDir(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, controller.sdk.GetCborDir())
}
func (controller *galleryController) LoadImages(c *gin.Context) {
	controller.loadImages(c)
}
func (controller *galleryController) LoadCborFiles(c *gin.Context) {
	controller.loadCborFiles(c)
}
func (controller *galleryController) SaveGalleryToCborDir(c *gin.Context) {
	controller.saveGalleryToCborDir(c)
}
