package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GiveTestResponse(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "this is a test response")
}
