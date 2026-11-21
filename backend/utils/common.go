package utils

import (
	"fmt"
	"time"
)

func GetUri(n int, pre string) string {
	now := time.Now()
	uri := pre + fmt.Sprintf("%d%d%d%d%d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute())
	uri = uri + RandSeq(8)
	return uri
}
