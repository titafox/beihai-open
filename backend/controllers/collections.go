package controllers

import (
	"coltstail-server/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetWorkList
// @Summary 首发作品列表
// @Description 首发作品列表
// @Tags 首发作品
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200
// @Router /product/list [get]
func GetWorkList(ctx *gin.Context) {
	user, err := GetUserInfo(ctx)
	if err != nil {
		EchoError(ctx, http.StatusBadRequest, 4000, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data":    models.GetWorkList(user.UserID),
		"success": true,
	})

}

func GetMyWorkList(ctx *gin.Context) {
	//nid := ctx.Param("id")
	//i, _ := strconv.ParseInt(nid, 10, 64)
	user, err := GetUserInfo(ctx)
	if err != nil {
		EchoError(ctx, http.StatusBadRequest, 4000, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data":    models.GetMyNftList(user.UserID),
		"success": true,
	})

}

// GetMyNftDetail
// @Summary 已购藏品合集内页
// @Description
// @Tags 二期接口列表
// @Accept application/json
// @Produce application/json
// @Param   id     query    int     true        "id"
// @Param   page     query    int     false        "分页"
// @Security ApiKeyAuth
// @Success 200
// @Router /account/nfts/series [get]
func GetMyNftDetail(ctx *gin.Context) {
	user, err := GetUserInfo(ctx)
	if err != nil {
		EchoError(ctx, http.StatusBadRequest, 4000, err.Error())
		return
	}

	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	info, err := models.GetMyNftDetail(id, user.UserID)
	if err != nil {
		fmt.Println("err=", err)
		EchoError(ctx, http.StatusBadRequest, 4000, "获取藏品信息失败")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data":    info,
		"success": true,
	})
}

// GetWorkInfo
// @Summary 首发作品详情
// @Description 首发作品详情
// @Tags 首发作品
// @Accept application/json
// @Produce application/json
// @Param   id     path    int     true        "首发作品id"
// @Security ApiKeyAuth
// @Success 200
// @Router /product/detail/{id} [get]
func GetWorkInfo(ctx *gin.Context) {
	user, err := GetUserInfo(ctx)
	if err != nil {
		EchoError(ctx, http.StatusBadRequest, 4000, err.Error())
		return
	}

	nid := ctx.Param("id")
	i, _ := strconv.ParseInt(nid, 10, 64)

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data":    models.GetWorkDetail(i, user.UserID),
		"success": true,
	})

}
