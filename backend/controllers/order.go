package controllers

import (
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

type OrderForm struct {
	Id  int64 `json:"id" binding:"required" comment:"作品id"`
	Num int   `json:"num" binding:"required,min=1,max=100" comment:"购买数量"`
}

// Create 购买首发订单创建
// @Summary 购买首发订单创建
// @Description 购买首发订单创建
// @Tags 二期接口列表
// @Accept application/json
// @Produce application/json
// @Param object body OrderForm true "查询参数"
// @Security ApiKeyAuth
// @Success 200
// @Router /order/commit [post]
func Create(ctx *gin.Context) {
	var form OrderForm
	if err := ctx.ShouldBindJSON(&form); err != nil {
		EchoError(ctx, http.StatusBadRequest, 4000, err.Error())
		return
	}

	//获取用户信息
	user, err := GetUserInfo(ctx)
	if err != nil {
		EchoError(ctx, http.StatusBadRequest, 4001, "获取用户信息失败")
		return
	}

	//防止重复提交
	key := fmt.Sprintf("create%d", user.UserID)

	//lua脚本，防止并发
	sha, err := models.Redis.ScriptLoad(`
if redis.call('get',KEYS[1]) == false then
redis.call('set',KEYS[1],1)
redis.call("expire", KEYS[1], 1)
return 0
end
return 1
`).Result()
	if err != nil {
		fmt.Println("lua:执行失败：", err.Error())
		EchoError(ctx, http.StatusBadRequest, 4008, "脚本执行失败")
		return
	}

	sta, err := models.Redis.EvalSha(sha, []string{key}).Int64()
	if err != nil {
		fmt.Println("lua:执行失败2：", err.Error())
		EchoError(ctx, http.StatusBadRequest, 4008, "脚本执行失败")
		return
	}

	if sta == 1 {
		EchoError(ctx, http.StatusForbidden, 4000, "重复请求，请稍后再试")
		return
	}

	//查询用户是否通过实名认证
	account, err := models.FindAccountByID(user.UserID)
	if err != nil {
		EchoError(ctx, http.StatusBadRequest, 4000, "获取用户信息失败")
		return
	}

	if account.Authentication != 1 {
		EchoError(ctx, http.StatusBadRequest, 4008, "请先进行实名认证")
		return
	}

	//查询钱包开户信息（放在扣库存之前，未开钱包的请求不会占用库存）
	var p utils.SelectWalletParam
	p.UserId = fmt.Sprintf("%d", user.UserID)
	user_cust_id, _, err := p.SelectWalletStatus()
	if err != nil {
		EchoError(ctx, http.StatusBadRequest, 4009, "请先开通钱包")
		return
	}

	//bol := XianLiu("order",200,3)
	//if bol == false {
	//	EchoError(ctx, http.StatusBadRequest, 4000, "公售过于火爆,请稍等几秒再试")
	//	return
	//}

	//获取藏品信息
	work, err := models.GetWoksById(form.Id)
	if err != nil {
		EchoError(ctx, http.StatusBadRequest, 4000, "获取藏品信息失败")
		return
	}

	//  总份数       预留份数
	//if work.StockSell+form.Num > work.Stock-work.Reserved {
	//	EchoError(ctx, http.StatusBadRequest, 400, "已卖完,感谢支持")
	//	return
	//}

	//判断库存是否已售罄 总量-预留-购买的份数
	if work.SoldOut {
		EchoError(ctx, http.StatusBadRequest, 400, "已售罄")
		return
	}

	//查询可购买的最大数量
	maxNum := models.GetMaxNum(user.UserID, work.ID)
	if maxNum == 0 {
		EchoError(ctx, http.StatusBadRequest, 400, "暂无购买名额")
		return
	}

	if maxNum < uint(form.Num) {
		EchoError(ctx, http.StatusBadRequest, 400, "超出购买份额，请减少数量")
		return
	}

	////   发行总量       预留份数         已售出份数       购买份数
	//if work.Stock-work.Reserved < work.StockSell+form.Num {
	//	EchoError(ctx, http.StatusBadRequest, 400, "超出剩余份数，请减少购买数量")
	//	return
	//}

	//redis版本
	//var s int
	//stock, err := models.Redis.Get(fmt.Sprintf("work_%d", work.ID)).Result()
	//if err != nil {
	//	//添加一下库存到redis
	//	//key := fmt.Sprintf("work_%d", work.ID)
	//	//models.Redis.Set(key, work.Stock-work.Reserved, 0).Result()
	//	s = work.Stock - work.Reserved
	//} else {
	//	s, _ = strconv.Atoi(stock)
	//}
	//
	//if s < form.Num {
	//	EchoError(ctx, http.StatusBadRequest, 4004, "超出剩余份数，请减少购买数量")
	//	return
	//}

	bol,err := models.ReduceStock(fmt.Sprintf("work_%d", work.ID),form.Num)
	if err != nil {
		EchoError(ctx, http.StatusBadRequest, 4004, err.Error())
		return
	}

	if !bol {
		EchoError(ctx, http.StatusBadRequest, 4004,"库存不足，请稍后")
		return
	}
	//已扣库存：后续任何一步失败都要把库存加回去
	stockKey := fmt.Sprintf("work_%d", work.ID)
	orderCreated := false
	defer func() {
		if !orderCreated {
			if _, e := models.Redis.IncrBy(stockKey, int64(form.Num)).Result(); e != nil {
				fmt.Println("回补库存失败：", stockKey, e.Error())
			}
		}
	}()

	//挑起支付参数
	transAmt := work.Price * float64(form.Num)
	orderId := utils.GenerateUri("BH")
	divDetails := []req.DivDetails{
		{RiskDivType: "1", DivAmt: strconv.FormatFloat(transAmt, 'f', 2, 64), DivAcctId: conf.Conf().DivAcctId, DivCustId: conf.Conf().MER_CUST_ID},
	}

	divDetailsJson, err := json.Marshal(divDetails)
	if err != nil {
		EchoError(ctx, http.StatusBadRequest, 4011, "分账参数设置失败")
		return
	}

	payParam := utils.BalancePayParam{
		OrderId:    orderId,
		TransAmt:   strconv.FormatFloat(transAmt, 'f', 2, 64),
		GoodsDesc:  work.Name,
		UserCustId: user_cust_id,
		DivDetails: string(divDetailsJson),
		NotifyUrl:  conf.Conf().BG_RET_URL,
	}

	fmt.Println("支付参数：", payParam)
	payUrl, err := payParam.BalancePay()
	if err != nil {
		EchoError(ctx, http.StatusBadRequest, 4012, "拉起支付失败："+err.Error())
		return
	}

	//开始下单
	err = models.CreateOrder(orderId, user.UserID, work.ID, work.Price, form.Num, transAmt, payUrl)
	if err != nil {
		EchoError(ctx, http.StatusBadRequest, 400, "订单创建失败")
		return
	}
	orderCreated = true
	//redis减库存
	//_, err = models.Redis.DecrBy(fmt.Sprintf("work_%d", work.ID), int64(form.Num)).Result()
	//if err != nil {
	//	fmt.Println("下单减库存失败，订单：", orderId)
	//}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data":    payUrl,
		"success": true,
	})
}

func XianLiu(key string, count uint, ttl int64) bool {
	reqCounts, _ := models.Redis.Get(key).Int64()
	if uint(reqCounts) < count {
		reqCounts, _ = models.Redis.Incr(key).Result()
		if reqCounts == 1 {
			models.Redis.Expire(key, time.Duration(ttl)*time.Second)
		}
		return true
	}

	if models.Redis.TTL(key).Val() <= 0 {
		err := models.Redis.Set(key, 1, time.Duration(ttl)*time.Second).Err()
		if err != nil {
			return false
		}

		return true
	}
	return false
}

// OrderList
// @Summary 订单记录
// @Description typ 0: "全部"，,1: "进行中", 2: "已完成", 3: "已取消"
// @Tags 支付相关
// @Accept application/json
// @Produce application/json
// @Param   typ     path    int     true        "订单类型"
// @Security ApiKeyAuth
// @Success 200
// @Router /order/list/{typ} [get]
func OrderList(ctx *gin.Context) {
	var r req.SelectOrder
	if err := ctx.ShouldBindUri(&r); err != nil {
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
		"data":    models.GetOrders(user.UserID, r.Typ),
		"success": true,
	})
}

func HandselList(ctx *gin.Context) {
	user, err := GetUserInfo(ctx)
	if err != nil {
		EchoError(ctx, http.StatusBadRequest, 4000, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data":    models.GetHandselList(user.UserID),
		"success": true,
	})
}

type HandselForm struct {
	Id    int64  `json:"id"`
	Phone string `json:"phone"`
}

func Handsel(ctx *gin.Context) {
	EchoError(ctx, http.StatusBadRequest, 4000, "转增功能已关闭")
	return

	var form HandselForm
	if err := ctx.ShouldBindJSON(&form); err != nil {
		EchoError(ctx, http.StatusBadRequest, 4000, err.Error())
		return
	}
	//

	user, err := GetUserInfo(ctx)
	if err != nil {
		EchoError(ctx, http.StatusBadRequest, 4000, err.Error())
		return
	}

	err = models.SetHandsel(user.UserID, form.Id, form.Phone)
	if err != nil {
		EchoError(ctx, http.StatusBadRequest, 4000, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "转增成功",
		"data":    nil,
		"success": true,
	})
}
