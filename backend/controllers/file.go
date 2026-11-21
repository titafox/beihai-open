package controllers

import (
	"coltstail-server/ecodes"
	"coltstail-server/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Credential(ctx *gin.Context) {
	data, err := utils.Token()
	if err != nil {
		ServerResponse(ctx, http.StatusBadRequest, ecodes.PARAM_ERROR, err.Error(), nil)
		return
	}

	ServerResponse(ctx, http.StatusOK, ecodes.SUCCESS, "", data)
}
