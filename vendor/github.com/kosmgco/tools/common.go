package tools

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckHealth(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

func RegisterCommonRoutesOn(router *gin.RouterGroup) {
	router.HEAD("", CheckHealth)
}
