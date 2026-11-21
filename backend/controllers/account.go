package controllers

import (
	"strings"
	"coltstail-server/conf"
	"coltstail-server/entity/req"
	"coltstail-server/models"
	"coltstail-server/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// GetAccountInfo
// @Summary 获取用户信息
// @Description 获取用户信息
// @Tags 用户
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200
// @Router /account/info [get]
func GetAccountInfo(ctx *gin.Context) {
	user, err := GetUserInfo(ctx)
	if err != nil {
		EchoError(ctx, http.StatusBadRequest, 4000, err.Error())
		return
	}

	account, err := models.FindAccountByID(user.UserID)
	if err != nil {
		EchoError(ctx, http.StatusBadRequest, 403, err.Error())
		return
	}
	//查询钱包开户信息
	var p utils.SelectWalletParam
	p.UserId = fmt.Sprintf("%d", account.ID)
	_, _, err = p.SelectWalletStatus()
	if err == nil {
		account.IsWallet = 1
	}
	// 身份证号打码后再返回（前端只用它判断是否已实名）
	if n := len(account.IdCardNum); n > 8 {
		account.IdCardNum = account.IdCardNum[:4] + strings.Repeat("*", n-8) + account.IdCardNum[n-4:]
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data":    account,
		"success": true,
	})
}

// AccountAuthentication
// @Summary 实名认证
// @Description 实名认证
// @Tags 用户
// @Accept application/json
// @Produce application/json
// @Param object body req.Authentication true "输入参数"
// @Security ApiKeyAuth
// @Success 200
// @Router /account/authentication [post]
func AccountAuthentication(ctx *gin.Context) {
	var r req.Authentication
	if err := ctx.ShouldBindJSON(&r); err != nil {
		EchoError(ctx, http.StatusBadRequest, 4000, err.Error())
		return
	}

	//获取用户信息
	user, err := GetUserInfo(ctx)
	if err != nil {
		EchoError(ctx, http.StatusBadRequest, 4000, "获取用户信息失败")
		return
	}

	//查询身份证信息是否已经实名过
	a, err := models.FindAccountByCard(r.IdCard)
	if err == nil {
		EchoError(ctx, http.StatusBadRequest, 4000, "当前身份信息已被手机号"+utils.HidePhone(a.Phone)+"绑定")
		return
	}

	//调用实名第三方接口
	bol, err := utils.CheckIdCard(r.IdCard, r.Name, user.Username)
	fmt.Println(bol, err)
	if err != nil {
		EchoError(ctx, http.StatusBadRequest, 4002, "实名认证失败")
		return
	}

	if !bol {
		EchoError(ctx, http.StatusBadRequest, 4000, "实名信息不一致")
		return
	}

	//保存认证信息
	var account models.TblAccount
	account.Authentication = 1
	account.Name = r.Name
	account.IdCardNum = r.IdCard
	account.UpdateTime = time.Now().Unix()
	err = account.UpdateAccountAuthentication(user.UserID)
	if err != nil {
		EchoError(ctx, http.StatusBadRequest, 4000, "保存认证信息失败")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "认证成功",
		"data":    nil,
		"success": true,
	})

}

// UpdateAccountInfo
// @Summary 更新个人信息
// @Description 更新个人信息
// @Tags 用户
// @Accept application/json
// @Produce application/json
// @Param object body req.UpdateAccountInfo true "查询参数"
// @Security ApiKeyAuth
// @Success 200
// @Router /account/update [post]
func UpdateAccountInfo(ctx *gin.Context) {
	var r req.UpdateAccountInfo
	if err := ctx.ShouldBindJSON(&r); err != nil {
		EchoError(ctx, http.StatusBadRequest, 4000, err.Error())
		return
	}

	//获取用户信息
	user, err := GetUserInfo(ctx)
	if err != nil {
		EchoError(ctx, http.StatusBadRequest, 4000, err.Error())
		return
	}

	err = models.UpdateAccountInfo(r.Nickname, r.Avatar, r.Desc, user.UserID)
	if err != nil {
		EchoError(ctx, http.StatusBadRequest, 4000, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data":    nil,
		"success": true,
	})

}

func GetQrCode(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data":    "iVBORw0KGgoAAAANSUhEUgAAAZAAAAGQAQMAAAC6caSPAAAABlBMVEX///8AAABVwtN+AAACJ0lEQVR42uybQY6lMAxEK8qCJUfgKP9ocDSOwhFYskB4ZDuBQDNM70afVEmtrwYeK8suOwYURVEURVFfq1aSNnxE0C1AtwKI+ZJrIULEkc1/mw2fBTI1UobTZ06/RN6K9CIyNhvQiCKmOYrMigT9jwiRG0TTyQTIpNllRa8ZhgiRpxizkgSPr/53YUmkOiRXMQBqVjTOVI48Fz4i9SGF6U22p1vRz1GG9hc+mchXI6XD1euKiAztesTSg4hUh1hJ8mAaNcOIx1gOu9ne9KNHJlIvogGljwcrR/a3Wjnq3b/om36UJCLVImpWxFtqVapi/iYNPR/OXUWkWkRLld5bgoyNOLoeb2p3L0zknYimE3s0ZxhPOgOCZRe9NYIIEVcws7LLTa+PboOPbq9DFSLVI66xOBRqM2K56tJWE6kZKdpqsfiyGJvjnpQ+y0MnTqRCJLdDGkidd0k+6w/WPI0gQqRMSkNOTj7jN9tTWGg0QuSliLVBIc/5LWCmmBDZhypEiJQZBtZW70eC0bsk8zHIE10iRHIyMmebxm15Cpdn/WcvTITI+SSxSwdDcxrc4R9dEpHakPOuglcx2K6CHLtwl1kfkdcgxa61yD62hW+q4K7lIVI5sm9C5hOgmDZV5NBChMjNrrV/yyN+mpj2EcaHrzmIVI0cy5MWZ+m2Pk6EyO2u9XGKGMu1pr/aHiJfjpy+FXWtyfQeK9hEiFAURVEURf1//QkAAP//ul/1E5qlzEgAAAAASUVORK5CYII=",
		"success": true,
	})
}

// MyNFTs
// @Summary 已购藏品
// @Description
// @Tags 二期接口列表
// @Accept application/json
// @Produce application/json
// @Param   page     query   int     false        "分页"
// @Security ApiKeyAuth
// @Success 200
// @Router /account/nfts [get]
func MyNFTs(ctx *gin.Context) {
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
		"data":    models.GetMyNFTs(user.UserID, r.Page),
		"success": true,
	})

}

// MyNFTsSeries
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
func MyNFTsSeries(ctx *gin.Context) {
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
		"data":    models.GetMyNFTsSeries(r.Id, user.UserID, r.Page),
		"success": true,
	})

}

// GetWorksNFT
// @Summary 单个作品流通列表
// @Description sort = price ,按价格排序，sort=no 按编号排序 默认价格排序
// @Tags 二期接口列表
// @Accept application/json
// @Produce application/json
// @Param   wid     query    int     true        "wid"
// @Param   sort     query    string     false        "排序"
// @Param   page     query    int     false        "分页"
// @Success 200
// @Router /market/nfts/circulating [get]
func GetWorksNFT(ctx *gin.Context) {
	var r req.GetWorksNft
	if err := ctx.ShouldBindQuery(&r); err != nil {
		EchoError(ctx, http.StatusBadRequest, 4000, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data":    models.GetWorksNFT(r.Wid, r.Sort, r.Page),
		"success": true,
	})

}

// MyNFTDetail
// @Summary 藏品详情
// @Description
// @Tags 二期接口列表
// @Accept application/json
// @Produce application/json
// @Param   id     path    int     true        "id"
// @Security ApiKeyAuth
// @Success 200
// @Router /account/nfts/{id} [get]
//func MyNFTDetail(ctx *gin.Context) {
//	var r req.MarketInfo
//	if err := ctx.ShouldBindUri(&r); err != nil {
//		EchoError(ctx, http.StatusBadRequest, 4000, err.Error())
//		return
//	}
//
//	info,err := models.GetMarketNftInfo(r.Id)
//	if err != nil {
//		EchoError(ctx, http.StatusBadRequest, 4001, "获取藏品信息失败")
//		return
//	}
//
//	ctx.JSON(http.StatusOK, gin.H{
//		"code":    0,
//		"message": "ok",
//		"data":  info  ,
//		"success": true,
//	})
//
//}

// MyNFTSells
// @Summary 我的寄售
// @Description
// @Tags 二期接口列表
// @Accept application/json
// @Produce application/json
// @Param   page     query    int     false        "page"
// @Security ApiKeyAuth
// @Success 200
// @Router /account/nfts/sells [get]
func MyNFTSells(ctx *gin.Context) {
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
		"data":    models.GetMySellsNFTs(user.UserID, r.Page),
		"success": true,
	})
}

// MyNFTSells
// @Summary 下单
// @Description
// @Tags 二期接口列表
// @Accept application/json
// @Produce application/json
// @Param object body req.MarketBuyTheNFT true "输入参数"
// @Security ApiKeyAuth
// @Success 200
// @Router /market/nft/buy [post]
func BuyNFT(ctx *gin.Context) {
	var r req.MarketBuyTheNFT
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

	//查询用户是否通过实名认证
	account, err := models.FindAccountByID(user.UserID)
	if err != nil {
		EchoError(ctx, http.StatusBadRequest, 4002, "获取用户信息失败")
		return
	}

	if account.Authentication != 1 {
		EchoError(ctx, http.StatusBadRequest, 4003, "请先进行实名认证")
		return
	}

	nft, err := models.GetNFTDetail(r.Id)
	if err != nil {
		EchoError(ctx, http.StatusBadRequest, 4005, "不存在此藏品")
		return
	}
	if nft.Status == 0 {
		EchoError(ctx, http.StatusBadRequest, 4006, "藏品已下架")
		return
	}

	if nft.Status == 2 {
		EchoError(ctx, http.StatusBadRequest, 4007, "藏品锁单中，无法购买")
		return
	}

	_, err = models.GetUnpaidOrder(user.UserID, r.Id)
	if err == nil {
		EchoError(ctx, http.StatusBadRequest, 4008, "你还有未支付的订单，请先完上一单")
		return
	}

	//获取藏品信息
	work, err := models.GetWoksById(nft.Wid)
	if err != nil {
		EchoError(ctx, http.StatusBadRequest, 4008, "获取藏品信息失败")
		return
	}

	//防止重复提交（SETNX 原子加锁）
	key := fmt.Sprintf("buy%d", r.Id)
	if got, _ := models.Redis.SetNX(key, 1, time.Second*10).Result(); !got {
		EchoError(ctx, http.StatusBadRequest, 4010, "藏品锁定中") //重复请求已退回
		return
	}

	//在拉起支付前原子锁单：只有 status=1（挂售中）才能改成 2（锁单），保证一件藏品只会被一个买家锁定
	if !models.LockNFTForSale(nft.ID, user.UserID) {
		EchoError(ctx, http.StatusBadRequest, 4007, "藏品锁单中，无法购买")
		return
	}
	orderCreated := false
	defer func() {
		if !orderCreated {
			models.UnlockNFTForSale(nft.ID)
		}
	}()
	//锁单后重新读取，确保价格/归属是最新的
	nft, err = models.GetNFTDetail(r.Id)
	if err != nil || nft.Status != 2 {
		EchoError(ctx, http.StatusBadRequest, 4005, "不存在此藏品")
		return
	}

	//查询钱包开户信息
	var p utils.SelectWalletParam
	p.UserId = fmt.Sprintf("%d", account.ID)
	user_cust_id, _, err := p.SelectWalletStatus()
	if err != nil {
		EchoError(ctx, http.StatusBadRequest, 4009, "请先开通钱包")
		return
	}

	//查询卖家钱包开户信息
	p.UserId = fmt.Sprintf("%d", nft.OwnerId)
	seller_user_cust_id, seller_acc_id, err := p.SelectWalletStatus()
	if err != nil {
		EchoError(ctx, http.StatusBadRequest, 4010, "卖家暂未开通钱包,无法支付")
		return
	}

	//平台收取4%服务费
	servieFee := utils.Decimal(nft.Price * 4 / 100)
	//卖家收益
	sellerMoney := nft.Price - servieFee

	//挑起支付参数
	divDetails := []req.DivDetails{
		{RiskDivType: "1", DivAmt: strconv.FormatFloat(sellerMoney, 'f', 2, 64), DivAcctId: seller_acc_id, DivCustId: seller_user_cust_id},
		{RiskDivType: "1", DivAmt: strconv.FormatFloat(servieFee, 'f', 2, 64), DivAcctId: conf.Conf().DivAcctId, DivCustId: conf.Conf().MER_CUST_ID},
	}

	divDetailsJson, err := json.Marshal(divDetails)
	if err != nil {
		EchoError(ctx, http.StatusBadRequest, 4011, "分账参数设置失败")
		return
	}

	orderId := utils.GenerateUri("BY")
	payParam := utils.BalancePayParam{
		OrderId:    orderId,
		TransAmt:   strconv.FormatFloat(nft.Price, 'f', 2, 64),
		GoodsDesc:  work.Name + fmt.Sprintf("#%d", nft.WorkNo),
		UserCustId: user_cust_id,
		DivDetails: string(divDetailsJson),
		NotifyUrl:  conf.Conf().BALANCE_BG_RET_URL,
	}
	payUrl, err := payParam.BalancePay()
	if err != nil {
		EchoError(ctx, http.StatusBadRequest, 4012, "拉起支付失败："+err.Error())
		return
	}

	//开始下单
	err = models.CreateMarketOrder(orderId, user.UserID, work.ID, nft.ID, nft.Price, 1, nft.Price, payUrl)
	if err != nil {
		EchoError(ctx, http.StatusBadRequest, 4013, "订单创建失败")
		return
	}
	orderCreated = true

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data":    payUrl,
		"success": true,
	})
}

// InviteTop
// @Summary 排行邀请
// @Description
// @Tags 二期接口列表
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200
// @Router /account/invite/top [get]
func InviteTop(ctx *gin.Context) {
	_, err := GetUserInfo(ctx)
	if err != nil {
		EchoError(ctx, http.StatusBadRequest, 4000, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data":    models.InviteTop(),
		"success": true,
	})

}

// MyInviteList
// @Summary 我的邀请
// @Description
// @Tags 二期接口列表
// @Accept application/json
// @Produce application/json
// @Param   page     query   int     false        "分页"
// @Security ApiKeyAuth
// @Success 200
// @Router /account/invite/list [get]
func MyInviteList(ctx *gin.Context) {
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
		"data":    models.GetMyInviteList(user.UserID, r.Page),
		"success": true,
	})

}
