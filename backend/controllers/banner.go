package controllers

import (
	"coltstail-server/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetBannerList
// @Summary 首页Banner轮播图
// @Description 首页Banner轮播图
// @Tags Banner
// @Accept application/json
// @Produce application/json
// @Success 200
// @Router /banner/list [get]
func GetBannerList(ctx *gin.Context) {
	list, err := models.GetBannerList()
	if err != nil {
		EchoError(ctx, http.StatusBadRequest, 4000, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data":    list,
		"success": true,
	})

}
