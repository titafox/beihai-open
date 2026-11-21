package controllers

import (
	"coltstail-server/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAppVersion
// @Summary APP版本更新
// @Description APP版本更新
// @Tags 版本更新
// @Accept application/json
// @Produce application/json
// @Success 200
// @Router /app/version [get]
func GetAppVersion(ctx *gin.Context) {
	version, err := models.GetVersion()
	if err != nil {
		EchoError(ctx, http.StatusBadRequest, 4000, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data":    version,
		"success": true,
	})

}
