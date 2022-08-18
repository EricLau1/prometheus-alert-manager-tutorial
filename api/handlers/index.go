package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterIndex(router *gin.Engine) {
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"hello": "world!"})
	})
}
