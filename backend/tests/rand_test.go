package tests

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRand(t *testing.T) {
	//array := [5]int{11,12,13,14,15}
	//s := rand.NewSource(time.Now().Unix())
	//r := rand.New(s)
	//i := r.Intn(len(array))
	//t.Log(i)

	s := generateRandomNumber(1, 10001, 10000)
	for k, v := range s {
		fmt.Println("k=", k, ",v=", v)
	}

}

// 生成count个[start,end)结束的不重复的随机数
func generateRandomNumber(start int, end int, count int) []int {
	//范围检查
	if end < start || (end-start) < count {
		return nil
	}

	//存放结果的slice
	nums := make([]int, 0)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		//生成随机数
		num := r.Intn((end - start)) + start

		//查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}

		if !exist {
			nums = append(nums, num)
		}
	}

	return nums
}
