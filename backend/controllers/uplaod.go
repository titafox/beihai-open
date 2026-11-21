package controllers

import (
	"coltstail-server/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// AvatarUpload
// @Summary 上传个人图像
// @Description 上传个人图像
// @Tags 用户
// @Accept application/json
// @Produce application/json
// @Param file formData file true "文件"
// @Security ApiKeyAuth
// @Success 200
// @Router /storage/avatarUpload [post]
func AvatarUpload(ctx *gin.Context) {
	file, _, err := ctx.Request.FormFile("file")
	if err != nil {
		EchoError(ctx, http.StatusBadRequest, 4000, "请上传图像")
		return
	}

	defer file.Close()
	var data = make([]byte, 512)
	if _, err = file.Read(data); err != nil {
		EchoError(ctx, http.StatusBadRequest, 4000, "上传出错了")
		return
	}

	contentType := http.DetectContentType(data)
	if _, err = file.Seek(0, 0); err != nil {
		EchoError(ctx, http.StatusBadRequest, 4000, "上传出错了")
		return
	}

	if contentType != "image/jpeg" && contentType != "image/png" {
		EchoError(ctx, http.StatusBadRequest, 4000, "上传出错了")
		return
	}

	filename := fmt.Sprintf("idcard/%s/%s.%s", time.Now().Format("20060102"), utils.GenerateUri("i"),
		contentType[6:])
	sto := utils.NewOssFileStorage(filename, contentType)
	path, err := sto.Save(file)
	if err != nil {
		EchoError(ctx, http.StatusBadRequest, 4000, "上传出错")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data":    path,
		"success": true,
	})
}
