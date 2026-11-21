package controllers

import (
	"coltstail-server/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type Line struct {
	Id   int    `json:"id"`
	Name string `json:"name" comment:"作品名称"`
	List []Info `json:"list"`
}

type Info struct {
	No    int     `json:"no" comment:"作品编号"`
	Time  int64   `json:"time" comment:"支付时间"`
	User  string  `json:"user" comment:"支付用户"`
	Money float64 `json:"money" comment:"金额"`
}

type LineInfo struct {
	Id   int64  `form:"id"`
	Date string `form:"date"`
}

func GetLine(ctx *gin.Context) {
	var r LineInfo
	if err := ctx.ShouldBindQuery(&r); err != nil {
		EchoError(ctx, http.StatusBadRequest, 4000, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data":    models.GetLine2(r.Id),
		"success": true,
	})
}

func GetDone(ctx *gin.Context) {
	var r LineInfo
	if err := ctx.ShouldBindQuery(&r); err != nil {
		EchoError(ctx, http.StatusBadRequest, 4000, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data":    models.GetdownList(r.Id, r.Date),
		"success": true,
	})
}

func GetUsers(ctx *gin.Context) {
	nt := time.Now().Unix()
	nowTime := Unwrap(nt, 10)
	nowTimeStr := strconv.FormatFloat(nowTime, 'g', 11, 64)

	oldTime := Unwrap(nt-600, 10)
	oldTimeStr := strconv.FormatFloat(oldTime, 'g', 11, 64)

	//OldTime := nowTime - 600

	size, _ := models.Redis.ZCount("users", oldTimeStr, nowTimeStr).Result()

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data":    size,
		"success": true,
	})
}
