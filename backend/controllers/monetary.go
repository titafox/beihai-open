package controllers

import (
	"coltstail-server/models"
	"coltstail-server/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

// WalletCreate
// @Summary 汇付钱包开户
// @Description 先查询用户是否已开通钱包,如果未开通调用此接口，否则调用钱包查询接口获取钱包链接。
// @Tags 二期接口列表
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200
// @Router /wallet/create [post]
func WalletCreate(ctx *gin.Context) {
	user, err := GetUserInfo(ctx)
	if err != nil {
		EchoError(ctx, http.StatusBadRequest, 4000, err.Error())
		return
	}

	account, err := models.FindAccountByID(user.UserID)
	if err != nil {
		EchoError(ctx, http.StatusBadRequest, 4001, err.Error())
		return
	}

	if account.IsWallet == 1 {
		EchoError(ctx, http.StatusBadRequest, 4002, "您已完成钱包开户,无需重复开户")
		return
	}

	if account.Authentication == 0 {
		EchoError(ctx, http.StatusBadRequest, 4003, "请先完成实名认证")
		return
	}

	var p utils.CreateWalletParam
	p.UserId = fmt.Sprintf("%d", account.ID)
	p.OrderId = utils.GenerateUri("cw")
	p.UserName = account.Name
	p.IdCard = account.IdCardNum
	walletUrl, err := p.CreateWallet()
	if err != nil {
		EchoError(ctx, http.StatusBadRequest, 4004, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data":    walletUrl,
		"success": true,
	})

}

// WalletSelect
// @Summary 汇付钱包查询
// @Description 汇付钱包查询
// @Tags 二期接口列表
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200
// @Router /wallet/info [get]
func WalletSelect(ctx *gin.Context) {
	user, err := GetUserInfo(ctx)
	if err != nil {
		EchoError(ctx, http.StatusBadRequest, 4000, err.Error())
		return
	}

	account, err := models.FindAccountByID(user.UserID)
	if err != nil {
		EchoError(ctx, http.StatusBadRequest, 4001, err.Error())
		return
	}

	var p utils.SelectWalletParam
	p.UserId = fmt.Sprintf("%d", account.ID)
	p.OrderId = utils.GenerateUri("sw")

	walletUrl, err := p.SelectWallet()
	if err != nil {
		EchoError(ctx, http.StatusBadRequest, 4001, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data":    walletUrl,
		"success": true,
	})
}

// 钱包开户回调
func WalletNotify(ctx *gin.Context) {
	data, err := io.ReadAll(ctx.Request.Body)
	fmt.Println("收到开户回调：", err, string(data))
	if err != nil {
		EchoError(ctx, http.StatusBadRequest, 4000, "回调错误:"+err.Error())
		return
	}
	//解签
	_, err = utils.CreateWalletVerifySign(string(data)[12:])
	if err != nil {
		EchoError(ctx, http.StatusBadRequest, 4000, "验签失败："+err.Error())
		return
	}
	//更新订单
	//models.UpdateOrderPayNotify(orderId)

}
