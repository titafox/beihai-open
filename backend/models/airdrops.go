package models

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type TblAirdrops struct {
	ID         int64  `json:"id" gorm:"primary_key"`
	WorkId     int64  `json:"work_id"`
	Whitelist  string `json:"whitelist"`
	StartTime  int64  `json:"start_time"`
	Status     int    `json:"status"`
	UpdateTime int64  `json:"update_time"`
}

type TblAirdropsLog struct {
	ID         int64  `json:"id" gorm:"primary_key"`
	AccountId  int64  `json:"account_id"`
	AirdropsId int64  `json:"airdrops_id"`
	PId        int64  `json:"p_id"`
	NId        int    `json:"n_id"`
	Message    string `json:"whitelist"`
	Status     int    `json:"status"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
}

func (log *TblAirdropsLog) Insert() error {
	var logDb = Eloquent
	return logDb.Create(log).Error
}

func StartAirdrops() {
	//查询未开始的空投
	var list []TblAirdrops
	err := Eloquent.Model(&TblAirdrops{}).Where("status = 0").Find(&list).Error
	if err == nil && len(list) > 0 {
		for _, v := range list {
			//时间未到
			if time.Now().Unix() < v.StartTime {
				continue
			}

			//先将未开始的改为进行中
			if txErr := Eloquent.Model(TblAirdrops{}).Where("id=?", v.ID).
				Updates(map[string]interface{}{
					"status": 1,
				}).Error; txErr != nil {
				fmt.Println("空投订单更新失败：" + txErr.Error())
				continue
			}

			//开始解析白名单
			wList := strings.Split(v.Whitelist, ",") //13800000001-5,
			for _, item := range wList {
				info := strings.Split(item, "-")
				//查询用户是否存在
				account, err := FindAccountByPhone(info[0])
				if err == nil {
					//获取空投份数
					num, _ := strconv.Atoi(info[1])
					var l1 []TblNfts
					err = Eloquent.Model(TblNfts{}).Where("wid= ? AND transfer_num=1 and chain_status=0", v.WorkId).Limit(num).Find(&l1).Error
					if err == nil && len(l1) > 0 {
						for _, j := range l1 {
							Eloquent.Model(TblNfts{}).Where("id=?", j.ID).
								Updates(map[string]interface{}{
									"chain_status": 4,
								})

							var log TblAirdropsLog
							log.AccountId = account.ID
							log.AirdropsId = v.ID
							log.PId = j.Wid
							log.NId = int(j.WorkNo)
							log.Status = 0
							log.CreateTime = time.Now().Unix()
							log.UpdateTime = time.Now().Unix()
							_ = log.Insert()

							r := RabbitMQ{Conn: OrderConn}
							err = r.Publish("mint", fmt.Sprintf("KT:%d", log.ID))
						}

					}
				}
			}

			//已结束
			if txErr := Eloquent.Model(TblAirdrops{}).Where("id=?", v.ID).
				Updates(map[string]interface{}{
					"status": 2,
				}).Error; txErr != nil {
				fmt.Println("空投订单更新失败：" + txErr.Error())
				continue
			}

		}
	}
}
