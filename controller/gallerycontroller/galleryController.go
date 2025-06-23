package gallerycontroller

import (
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

// func (controller *galleryController) UpdateCborDir(c gin.Context)        {}
// func (controller *galleryController) LoadImages(c gin.Context)           {}
// func (controller *galleryController) LoadCborFiles(c gin.Context)        {}
// func (controller *galleryController) SaveGalleryToCborDir(c gin.Context) {}
