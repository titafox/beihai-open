package controllers

import (
	"coltstail-server/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetNoticeType
// @Summary 获取公告类型
// @Description 获取公告类型
// @Tags 公告
// @Accept application/json
// @Produce application/json
// @Success 200
// @Router /notice/type [get]
func GetNoticeType(ctx *gin.Context) {
	list := models.GetNoticeType()
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data":    list,
		"success": true,
	})
}

// GetNoticeAll
// @Summary 综合公告
// @Description 综合公告
// @Tags 公告
// @Accept application/json
// @Produce application/json
// @Success 200
// @Router /notice/type/all [get]
func GetNoticeAll(ctx *gin.Context) {
	list := models.GetNoticeTypeAll()
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data":    list,
		"success": true,
	})
}

// GetTypeNoticeAll
// @Summary 某个分类下所有公告
// @Description 某个分类下所有公告
// @Tags 公告
// @Accept application/json
// @Produce application/json
// @Param   nid     path    int     true        "分类id"
// @Success 200
// @Router /notice/type/list/{nid} [get]
func GetTypeNoticeAll(ctx *gin.Context) {
	nid := ctx.Param("nid")
	i, _ := strconv.ParseInt(nid, 10, 64)
	//fmt.Println(nid)
	list := models.GetNoticeList(i)
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data":    list,
		"success": true,
	})
}

// GetNoticeLaba
// @Summary 公告小喇叭
// @Description 公告小喇叭
// @Tags 公告
// @Accept application/json
// @Produce application/json
// @Success 200
// @Router /notice [get]
func GetNoticeLaba(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data":    models.GetNoticeLaba(),
		"success": true,
	})
}

// GetNoticeDetail
// @Summary 公告详情
// @Description 公告详情
// @Tags 公告
// @Accept application/json
// @Produce application/json
// @Param   id     path    int     true        "公告id"
// @Success 200
// @Router /notice/detail/{id} [get]
func GetNoticeDetail(ctx *gin.Context) {
	nid := ctx.Param("id")
	i, _ := strconv.ParseInt(nid, 10, 64)
	//fmt.Println(nid)
	list := models.GetNoticeDetail(i)
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data":    list,
		"success": true,
	})
}
