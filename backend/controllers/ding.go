package controllers

import (
	"coltstail-server/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ding(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data":    models.Ding(),
		"success": true,
	})
}
