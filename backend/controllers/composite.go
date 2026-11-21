package controllers

import (
	"coltstail-server/entity/req"
	"coltstail-server/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CompositeNFT
// @Summary 合成活动列表
// @Description 合成活动列表
// @Tags 合成
// @Accept application/json
// @Produce application/json
// @Success 200
// @Router /nft/composite/list [get]
func CompositeNFT(ctx *gin.Context) {
	//str := `[{"wid":5,"nums":5},{"wid":5,"nums":5}]`
	//go models.BSNCrete()
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data":    models.GetCompositeActivityList(),
		"success": true,
	})
}

// GetCompositeNFTMaterial
// @Summary 获取合成材料
// @Description 获取合成材料
// @Tags 合成
// @Accept application/json
// @Produce application/json
// @Param object query req.Page true "输入参数"
// @Security ApiKeyAuth
// @Success 200
// @Router /nft/composite/material [get]
func GetCompositeNFTMaterial(ctx *gin.Context) {
	var r req.Page
	if err := ctx.ShouldBindQuery(&r); err != nil {
		EchoError(ctx, http.StatusBadRequest, 4000, err.Error())
		return
	}

	user, err := GetUserInfo(ctx)
	if err != nil {
		EchoError(ctx, http.StatusBadRequest, 4000, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data":    models.GetCompositeNFTMaterial(r.Id,user.UserID,r.Page),
		"success": true,
	})
}

// CompositeStart
// @Summary 合成
// @Description 获取合成材料
// @Tags 合成
// @Accept application/json
// @Produce application/json
// @Param object body req.StartCompositeData true "输入参数"
// @Security ApiKeyAuth
// @Success 200
// @Router /nft/composite/start [post]
func CompositeStart(ctx *gin.Context) {
	var r req.StartCompositeData
	if err := ctx.ShouldBindJSON(&r); err != nil {
		EchoError(ctx, http.StatusBadRequest, 4000, err.Error())
		return
	}

	user, err := GetUserInfo(ctx)
	if err != nil {
		EchoError(ctx, http.StatusBadRequest, 4007, err.Error())
		return
	}

	//if user.Username != "13800000001"{
	//	EchoError(ctx, http.StatusBadRequest, 4006, "您暂无合成权限")
	//	return
	//}
	key := fmt.Sprintf("co%d",user.UserID)
	bol ,_ := models.ReApply(key)
	if bol == true {
		EchoError(ctx, http.StatusBadRequest, 4006, "重复请求,请稍后再试")
		return
	}

	dt,err :=  models.CompositeStart(user.UserID,r)
	if err != nil {
		models.Redis.Del(key)
		EchoError(ctx, http.StatusBadRequest, 4004, err.Error())
		return
	}

	models.Redis.Del(key)

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data": dt  ,
		"success": true,
	})
}
