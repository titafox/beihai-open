package controllers

import (
	"coltstail-server/middlewares"
	"coltstail-server/models"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"math"
	"time"
)

func GetUserInfo(ctx *gin.Context) (*middlewares.CustomClaims, error) {
	user, bol := ctx.Get("jwt")
	if !bol {
		return nil, errors.New("token not found")
	}
	account := user.(*middlewares.CustomClaims)
	//访问人数
	f := Unwrap(time.Now().Unix(), 10)
	models.Redis.ZAdd("users", redis.Z{f, account.Username})

	return account, nil
}

// 将int64恢复成正常的float64
func Unwrap(num int64, retain int) float64 {
	return float64(num) / math.Pow10(retain)
}
