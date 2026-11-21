package models

import (
	"coltstail-server/conf"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"sync"
)

var Redis *redis.Client
var mutex sync.Mutex

func InitRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     conf.Conf().REDIS_ADDR, // redis地址
		Password: "",                     // redis密码，没有则留空
		DB:       0,                      // 默认数据库，默认是0
	})
}

//减库存
func ReduceStock(key string ,val int) (bool,error) {

	//lua脚本，防止并发
	sha, err :=Redis.ScriptLoad(`
local kc = tonumber(redis.call('get',KEYS[1]))
if kc - ARGV[1] > 0 then
redis.call('INCRBY',KEYS[1],-ARGV[1])
return 1
end
return 0
`).Result()
	if err != nil {
		return false, errors.New("脚本执行失败")
	}

	sta, err := Redis.EvalSha(sha, []string{key},val).Int64()
	if err != nil {
		return false, errors.New("脚本执行失败2")
	}


	if sta == 0 {
		return false,errors.New("库存不足")
	}

	return true,nil
}

//加库存
func AddStock(key string ,val int64) (bool,error) {

	//lua脚本，防止并发
	sha, err :=Redis.ScriptLoad(`
local kc = tonumber(redis.call('get',KEYS[1]))
if kc - ARGV[1] > 0 then
redis.call('INCRBY',KEYS[1],ARGV[1])
return 1
end
return 0
`).Result()
	if err != nil {
		return false, errors.New("脚本执行失败")
	}

	sta, err := Redis.EvalSha(sha, []string{key},val).Int64()
	if err != nil {
		return false, errors.New("脚本执行失败2")
	}


	if sta == 0 {
		return false,errors.New("库存不足")
	}

	return true,nil
}

//防止重复请求
func ReApply(key string) (bool,error) {

	//lua脚本，防止并发
	sha, err :=Redis.ScriptLoad(`
local is_exists = redis.call('EXISTS',KEYS[1])
if is_exists==1 then
return 1
end
redis.call('SET',KEYS[1],10)
return 0
`).Result()
	if err != nil {
		fmt.Println("err：",err)
		return true, errors.New("脚本执行失败")
	}
	fmt.Println("sha=",sha)

	sta, e := Redis.EvalSha(sha, []string{key}).Int64()
	fmt.Println("sta=",sta,"e=",e)


	if sta == 1 {
		return true,errors.New("重复请求")
	}

	return false,nil
}


func ReApply2(key string) (bool,error) {

	//lua脚本，防止并发
	sha, err :=Redis.ScriptLoad(`
return tonumber(redis.call('get',KEYS[1]))
`).Result()
	if err != nil {
		return false, errors.New("脚本执行失败")
	}
	fmt.Println("sha=",sha)

	sta, _ := Redis.EvalSha(sha, []string{key}).Int64()

	fmt.Println("sta=",sta)


	if sta == 1 {
		return false,errors.New("重复请求")
	}

	return true,nil
}
