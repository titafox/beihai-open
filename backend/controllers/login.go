package controllers

import (
	"os"
	"coltstail-server/entity/req"
	"coltstail-server/middlewares"
	"coltstail-server/models"
	"coltstail-server/utils"
	"fmt"
	"github.com/gin-gonic/gin"

	"net/http"
	"time"
)

// Login
// @Summary 短信登录
// @Description 短信登录
// @Tags 用户
// @Accept application/json
// @Produce application/json
// @Param object body req.Login true "输入参数"
// @Success 200
// @Router /account/smsLogin [post]
func Login(ctx *gin.Context) {
	var r req.Login
	if err := ctx.ShouldBindJSON(&r); err != nil {
		EchoError(ctx, http.StatusBadRequest, 4000, err.Error())
		return
	}

	//验证手机号+验证码是否匹配
	var sms models.TblSms
	err := sms.CheckSmsCode(r.Phone, r.Code, 1)
	if err != nil {
		EchoError(ctx, http.StatusBadRequest, 4001, err.Error())
		return
	}

	//如果填写了邀请码，验证邀请码是否正确
	var pid, id int64
	var account *models.TblAccount
	if r.Invitation != "" {
		account, err = account.CheckInvitation(r.Invitation)
		if err != nil {
			EchoError(ctx, http.StatusBadRequest, 4002, "邀请码不正确")
			return
		}
		pid = account.ID
	}

	//验证用户是否存在
	account, err = account.FindAccountByPhone(r.Phone)
	fmt.Println("account=", account)
	if err != nil { //不存在，
		var avata utils.Avata
		bsnAddress, err3 := avata.CreateAccount(utils.GenerateUri("n"), utils.GenerateUri("op"))
		if err3 != nil {
			EchoError(ctx, http.StatusBadRequest, 4001, "获取bsn地址失败")
			return
		}

		account.Uid = utils.GenerateRandomString(8)
		account.Pid = pid
		account.Role = 0
		account.Phone = r.Phone
		account.Status = 1
		account.Authentication = 0
		account.Nickname = utils.HidePhone(r.Phone)
		account.Avatar = os.Getenv("DEFAULT_AVATAR_URL")
		account.BsnAddress = bsnAddress
		account.IsHandsel = 1
		account.Nums = 0
		account.CreateTime = time.Now().Unix()
		account.UpdateTime = time.Now().Unix()
		id, err = account.Insert()
		if err != nil {
			fmt.Println("创建用户失败：", err)
			EchoError(ctx, http.StatusBadRequest, 4002, "登录失败")
			return
		}
		account.ID = id
	}
	//写入jwt
	token, err := middlewares.CreateJWT(account)
	if err != nil {
		EchoError(ctx, http.StatusInternalServerError, 4003, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data":    "Bearer " + token,
		"success": false,
	})
}
