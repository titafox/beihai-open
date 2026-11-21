package controllers

import (
	"coltstail-server/entity/req"
	"coltstail-server/models"
	"coltstail-server/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"crypto/rand"
	"math/big"
	"net/http"
	"time"
)

// Send
// @Summary 获取登录验证码
// @Description 获取登录验证码
// @Tags 用户
// @Accept application/json
// @Produce application/json
// @Param object body req.Send true "输入参数"
// @Success 200
// @Router /account/smsLogin/code [post]
func Send(ctx *gin.Context) {
	var r req.Send
	if err := ctx.ShouldBindJSON(&r); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    "4000",
			"message": "参数错误",
			"data":    nil,
			"success": false,
		})
		return
	}

	//按 IP 限频：每小时最多 10 次，防止轮换手机号刷短信
	if models.Redis != nil {
		ipKey := "sms:ip:" + ctx.ClientIP()
		cnt, _ := models.Redis.Incr(ipKey).Result()
		if cnt == 1 {
			models.Redis.Expire(ipKey, time.Hour)
		}
		if cnt > 10 {
			EchoError(ctx, http.StatusTooManyRequests, 4000, "操作频繁，请稍后再试。")
			return
		}
	}

	var sms models.TblSms
	//查询间隔是否大于2分钟
	err := sms.FindSendCount(r.Phone, 1)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    "4000",
			"message": err.Error(),
			"data":    nil,
			"success": false,
		})
		return
	}

	//生成六位数验证码（crypto/rand）
	n, err := rand.Int(rand.Reader, big.NewInt(1000000))
	if err != nil {
		EchoError(ctx, http.StatusInternalServerError, 4000, "验证码生成失败")
		return
	}
	code := fmt.Sprintf("%06d", n.Int64())
	var createSms models.TblSms
	createSms.Mobile = r.Phone
	createSms.Code = code
	createSms.Event = 1
	createSms.SendTime = time.Now().Unix()
	err = createSms.Insert()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    "4000",
			"message": "验证码发送失败",
			"data":    nil,
			"success": false,
		})
		return
	}

	//开始发送
	bol, err := utils.SendMsCode(code, r.Phone)
	if !bol || err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    "4001",
			"message": "验证码发送失败",
			"data":    nil,
			"success": false,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data":    nil,
		"success": true,
	})
}
