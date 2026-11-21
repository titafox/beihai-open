package controllers

import (
	"coltstail-server/models"
	"coltstail-server/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func Pay(ctx *gin.Context) {
	data, err := io.ReadAll(ctx.Request.Body)
	//fmt.Println("收到支付回调：",err,string(data))
	if err != nil {
		EchoError(ctx, http.StatusBadRequest, 4000, "回调错误:"+err.Error())
		return
	}
	//解签
	orderId, err := utils.VerifySign(string(data)[12:])
	if err != nil {
		EchoError(ctx, http.StatusBadRequest, 4000, "验签失败："+err.Error())
		return
	}
	//更新订单
	models.UpdateOrderPayNotify(orderId)
}
func BalancePay(ctx *gin.Context) {
	data, err := io.ReadAll(ctx.Request.Body)
	//fmt.Println("收到支付回调：",err,string(data))
	if err != nil {
		EchoError(ctx, http.StatusBadRequest, 4000, "回调错误:"+err.Error())
		return
	}
	//解签
	orderId, err := utils.VerifySign(string(data)[12:])
	fmt.Println("orderId=", orderId)
	if err != nil {
		EchoError(ctx, http.StatusBadRequest, 4000, "验签失败："+err.Error())
		return
	}
	//更新订单
	models.UpdateOrderBalancePayNotify(orderId)
}
