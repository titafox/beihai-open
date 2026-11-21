package models

import (
	"math"
	"strconv"
	"time"
)

type TblUsers struct {
	Id    int64 `json:"id"`
	Users int64 `json:"users"`
}

func GetOnlineUser() {
	nt := time.Now().Unix()
	nowTime := Unwrap(nt, 10)
	nowTimeStr := strconv.FormatFloat(nowTime, 'g', 11, 64)

	oldTime := Unwrap(nt-600, 10)
	oldTimeStr := strconv.FormatFloat(oldTime, 'g', 11, 64)

	size, _ := Redis.ZCount("users", oldTimeStr, nowTimeStr).Result()

	if size > 0 {
		Eloquent.Model(&TblUsers{}).Where("id=1").UpdateColumns(&TblUsers{
			Users: size,
		})
	}

}

// 将int64恢复成正常的float64
func Unwrap(num int64, retain int) float64 {
	return float64(num) / math.Pow10(retain)
}
