package controllers

import (
	"coltstail-server/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ArtList(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data":    models.GetIpsAll(),
		"success": true,
	})
}
