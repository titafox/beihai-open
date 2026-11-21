package tests

import (
	"coltstail-server/models"
	"fmt"
	"github.com/go-redis/redis"
	"math"
	"strconv"
	"testing"
	"time"
)

func TestBol(t *testing.T) {
	test2(false)
}

func test2(bol bool)  {
	if !bol {
		fmt.Println("jia")
	}else {
		fmt.Println("zhen")
	}

}

func TestGet(t *testing.T) {
	models.InitRedis()
	a,b := models.ReApply("x2")
	t.Log(a,b)
	//a,b := models.ReduceStock("test2",500)
	//t.Log(a,b)
	////a,e :=models.Redis.Get("work_2").Result()
	////t.Log(a,e)
	////a,e :=models.Redis.Get("work_2111").Result()
	////fmt.Println(a,e)
	//bol := models.Lock("d")
	//fmt.Println(bol)
	//
	//a, b := models.Redis.Set("d", 4, 0).Result()
	//t.Log(a, b)
}

func TestIncr(t *testing.T) {
	models.InitRedis()
	//s,e :=services.Redis.IncrBy("work_2",10).Result()
	//t.Log(s,e)

	a, b := models.Redis.Get("c").Result()
	t.Log(a, b)

	a, b = models.Redis.Set("c", 4, 0).Result()
	t.Log(a, b)

	//a,b := models.Redis.Set("b",4,0).Result()
	//t.Log(a,b)
}

//func TestDecr(t *testing.T) {
//	services.InitRedis()
//	s,e :=services.Redis.DecrBy("KC123",2000).Result()
//	t.Log(s,e)
//}

func TestFloat4(t *testing.T) {
	ts := time.Now().Unix()
	f := Unwrap(ts, 10)
	t.Log(ts)
	t.Log(f)

}

func TestSet(t *testing.T) {
	models.InitRedis()
	ts := time.Now().Unix()
	//t.Log(ts)
	f := Unwrap(ts, 10)
	//t.Log(f)
	//t.Log(Wrap(f,10))
	a, b := models.Redis.ZAdd("users", redis.Z{f, "13800000001"}).Result()
	t.Log(a, b)
}

func TestZcount(t *testing.T) {
	//string:=strconv.FormatInt(int64,10)
	models.InitRedis()
	nt := time.Now().Unix()
	nowTime := Unwrap(nt, 10)
	nowTimeStr := strconv.FormatFloat(nowTime, 'g', 11, 64)

	oldTime := Unwrap(nt-600, 10)
	oldTimeStr := strconv.FormatFloat(oldTime, 'g', 11, 64)
	t.Log(nowTime)
	t.Log(nowTimeStr)

	t.Log(oldTime)
	t.Log(oldTimeStr)

	//OldTime := nowTime - 600

	size, err := models.Redis.ZCount("users", oldTimeStr, nowTimeStr).Result()
	t.Log(size, err)

}

func TestFF(t *testing.T) {
	models.InitRedis()

	for i := 0; i < 100; i++ {
		bol := XianLiu("test2", 5, 20)
		if bol {
			t.Log(bol)
		} else {
			t.Error("公售过于火爆")
			break
		}

	}

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

// 将float64转成精确的int64
func Wrap(num float64, retain int) int64 {
	return int64(num * math.Pow10(retain))
}

// 将int64恢复成正常的float64
func Unwrap(num int64, retain int) float64 {
	return float64(num) / math.Pow10(retain)
}

// 精准float64
func WrapToFloat64(num float64, retain int) float64 {
	return num * math.Pow10(retain)
}

// 精准int64
func UnwrapToInt64(num int64, retain int) int64 {
	return int64(Unwrap(num, retain))
}
