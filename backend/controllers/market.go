package controllers

import (
	"coltstail-server/entity/req"
	"coltstail-server/models"
	"coltstail-server/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Market
// @Summary market
// @Description typ:目前默认1,wid:接口返回series对象里的id,sort:默认最新上架，sort=time_desc 最新上架；price_desc 价格高到低；price_asc 底到高
// @Tags 二期接口列表
// @Accept application/json
// @Produce application/json
// @Param object query req.Market true "查询参数"
// @Success 200
// @Router /market [get]
func Market(ctx *gin.Context) {
	var r req.Market
	if err := ctx.ShouldBindQuery(&r); err != nil {
		EchoError(ctx, http.StatusBadRequest, 4000, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data":    models.GetList(r.Page, r.Typ, r.Wid, r.Sort,r.Keyword),
		"success": true,
	})

}

// MarketInfo
// @Summary 藏品详情
// @Description
// @Tags 二期接口列表
// @Accept application/json
// @Produce application/json
// @Param   id     query    int     true        "藏品id"
// @Security ApiKeyAuth
// @Success 200
// @Router /market/detail [get]
func MarketInfo(ctx *gin.Context) {
	var r req.Page
	if err := ctx.ShouldBindQuery(&r); err != nil {
		EchoError(ctx, http.StatusBadRequest, 4000, err.Error())
		return
	}

	//获取用户信息
	user, err := GetUserInfo(ctx)
	if err != nil {
		EchoError(ctx, http.StatusBadRequest, 4001, "获取用户信息失败")
		return
	}

	info, err := models.GetMarketNftInfo(r.Id, user.UserID)
	if err != nil {
		EchoError(ctx, http.StatusBadRequest, 4001, "获取藏品信息失败")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data":    info,
		"success": true,
	})

}

// MarketOnTheNFT
// @Summary 寄售
// @Description
// @Tags 二期接口列表
// @Accept application/json
// @Produce application/json
// @Param object body req.MarketOnTheNFT true "输入参数"
// @Security ApiKeyAuth
// @Success 200
// @Router /market/nft/on [post]
func MarketOnTheNFT(ctx *gin.Context) {
	var r req.MarketOnTheNFT
	if err := ctx.ShouldBindJSON(&r); err != nil {
		EchoError(ctx, http.StatusBadRequest, 4000, err.Error())
		return
	}
	//获取用户信息
	user, err := GetUserInfo(ctx)
	if err != nil {
		EchoError(ctx, http.StatusBadRequest, 4001, "获取用户信息失败")
		return
	}

	//查询卖家钱包开户信息
	var p utils.SelectWalletParam
	p.UserId = fmt.Sprintf("%d", user.UserID)
	_, _, err = p.SelectWalletStatus()
	if err != nil {
		EchoError(ctx, http.StatusBadRequest, 4010, "请先开通钱包")
		return
	}
	err = models.OnTheNFT(r.Id, user.UserID, r.Price)
	if err != nil {
		EchoError(ctx, http.StatusBadRequest, 4002, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data":    nil,
		"success": true,
	})

}

// MarketOffTheNFT
// @Summary 下架
// @Description
// @Tags 二期接口列表
// @Accept application/json
// @Produce application/json
// @Param object body req.MarketOffTheNFT true "输入参数"
// @Security ApiKeyAuth
// @Success 200
// @Router /market/nft/off [post]
func MarketOffTheNFT(ctx *gin.Context) {
	var r req.MarketOffTheNFT
	if err := ctx.ShouldBindJSON(&r); err != nil {
		EchoError(ctx, http.StatusBadRequest, 4000, err.Error())
		return
	}
	//获取用户信息
	user, err := GetUserInfo(ctx)
	if err != nil {
		EchoError(ctx, http.StatusBadRequest, 4001, "获取用户信息失败")
		return
	}
	err = models.OffTheNFT(r.Id, user.UserID)
	if err != nil {
		EchoError(ctx, http.StatusBadRequest, 4002, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data":    nil,
		"success": true,
	})

}
