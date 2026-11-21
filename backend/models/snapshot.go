package models

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type TblSnapshot struct {
	ID          int64  `json:"id" gorm:"primary_key"`
	AId         int64  `json:"a_id" comment:"快照作品"`
	BId         int64  `json:"b_id" comment:"关联作品"`
	StartTime   int64  `json:"start_time" comment:"快照开始时间"`
	Status      int    `json:"status" comment:"0未开始1进行中2已结束"`
	SnapshotMod int    `json:"snapshot_mod" comment:"0:手动1:自动"`
	Whitelist   string `json:"whitelist" comment:"手动名单"`
	UpdateTime  int64  `json:"update_time" comment:"更新时间"`
}

type TblSnapshotLog struct {
	ID          int64 `json:"id" gorm:"primary_key"`
	AccountId   int64 `json:"account_id" comment:"用户"`
	WorkId      int64 `json:"work_id" comment:"可提前购作品"`
	SnapshotId  int64 `json:"snapshot_id" comment:"快照id"`
	BuyTotal    int   `json:"buy_total" comment:"最大可购买份数"`
	UseTotal    int   `json:"use_total" comment:"已购买份数"`
	SnapshotMod int   `json:"snapshot_mod" comment:"快照模式:0手动,1:自动"`
	CreateTime  int64 `json:"create_time"`
	UpdateTime  int64 `json:"update_time" comment:"更新时间"`
}

type SnapshotList struct {
	Id  int64
	Num int
}

func (snapShotLog *TblSnapshotLog) Insert() (int64, error) {
	err := Eloquent.Create(snapShotLog).Error
	if err != nil {
		return 0, err
	}
	return snapShotLog.ID, nil
}

// 获取一条快照记录
func GetSanpShotLogByOne(account_id, work_id int64) (TblSnapshotLog, error) {
	var sanpShotLog TblSnapshotLog
	err := Eloquent.Model(&TblSnapshotLog{}).Select("id,account_id,work_id,buy_total").Where("account_id=? AND work_id=?", account_id, work_id).First(&sanpShotLog).Error
	return sanpShotLog, err
}

func StartSnapshot() {
	var list []TblSnapshot
	//查询所有未开始的项目
	err := Eloquent.Model(&TblSnapshot{}).Where("status = 0").Find(&list).Error
	if err == nil && len(list) > 0 {
		for _, v := range list {
			//查询开始时间
			fmt.Println("当前时间：", time.Now())
			if v.StartTime > time.Now().Unix() {
				continue
			}
			//先将未开始的改为进行中
			if txErr := Eloquent.Model(TblSnapshot{}).Where("id=?", v.ID).
				Updates(map[string]interface{}{
					"status":      1,
					"update_time": time.Now().Unix(),
				}).Error; txErr != nil {
				continue
			}

			//查询快照模式
			if v.SnapshotMod == 0 { //手动模式
				//解析名单
				if v.Whitelist == "" {
					continue
				}

				whitelist := strings.Split(v.Whitelist, ",")
				for _, wl := range whitelist {
					info := strings.Split(wl, "-")
					//查询用户是否存在
					account, errAcc := FindAccountByPhone(info[0])
					if errAcc != nil {
						fmt.Println("手动快照用户:", info[0], "不存在")
						continue
					}

					//获取数量
					num, err3 := strconv.Atoi(info[1])
					if err3 != nil {
						fmt.Println("手动快照用户:", info[0], "提前购数量不正确")
						continue
					}

					//存在则插入
					var tblSnapshotlog TblSnapshotLog
					err = Eloquent.Model(&TblSnapshotLog{}).Where("work_id=? and account_id=?", v.BId, account.ID).First(&tblSnapshotlog).Error
					if err != nil {
						tblSnapshotlog.AccountId = account.ID
						tblSnapshotlog.WorkId = v.BId
						tblSnapshotlog.SnapshotId = v.AId
						tblSnapshotlog.BuyTotal = num
						tblSnapshotlog.UseTotal = 0
						tblSnapshotlog.SnapshotMod = v.SnapshotMod
						tblSnapshotlog.CreateTime = time.Now().Unix()
						tblSnapshotlog.UpdateTime = time.Now().Unix()
						_, err = tblSnapshotlog.Insert()
						if err != nil {
							fmt.Println("用户", v.AId, "增加快照失败：", err.Error())
						}
					} else { //update
						if txErr := Eloquent.Model(TblSnapshotLog{}).Where("id=?", tblSnapshotlog.ID).
							Updates(map[string]interface{}{
								"buy_total":   tblSnapshotlog.BuyTotal + num,
								"update_time": time.Now().Unix(),
							}).Error; txErr != nil {
							continue
						}
					}

				}
			} else {
				var snapshotList []SnapshotList
				err = Eloquent.Raw("select ta.id,count(*) as num from tbl_account ta inner join tbl_nfts tn on ta.id = tn.owner_id and tn.wid = ? and tn.update_time < ? group by ta.id", v.AId, v.StartTime).Scan(&snapshotList).Error
				if err == nil && len(snapshotList) > 0 {
					//保存快照
					for _, item := range snapshotList {
						var tblSnapshotlog TblSnapshotLog
						err = Eloquent.Model(&TblSnapshotLog{}).Where("work_id=? and account_id=?", v.BId, item.Id).First(&tblSnapshotlog).Error
						if err != nil {
							tblSnapshotlog.AccountId = item.Id
							tblSnapshotlog.WorkId = v.BId
							tblSnapshotlog.SnapshotId = v.AId
							tblSnapshotlog.BuyTotal = item.Num
							tblSnapshotlog.UseTotal = 0
							tblSnapshotlog.SnapshotMod = v.SnapshotMod
							tblSnapshotlog.CreateTime = time.Now().Unix()
							tblSnapshotlog.UpdateTime = time.Now().Unix()
							_, err = tblSnapshotlog.Insert()
							if err != nil {
								fmt.Println("用户", v.AId, "增加快照失败：", err.Error())
							}
						} else { //update
							if txErr := Eloquent.Model(TblSnapshotLog{}).Where("id=?", tblSnapshotlog.ID).
								Updates(map[string]interface{}{
									"buy_total":   tblSnapshotlog.BuyTotal + item.Num,
									"update_time": time.Now().Unix(),
								}).Error; txErr != nil {
								continue
							}
						}
					}

				}
			}

			//改为已结束
			if txErr := Eloquent.Model(TblSnapshot{}).Where("id=?", v.ID).
				Updates(map[string]interface{}{
					"status":      2,
					"update_time": time.Now().Unix(),
				}).Error; txErr != nil {
				continue
			}

		}

	}

}

func TestSql() {
	var list []SnapshotList
	Eloquent.Raw("select ta.id,count(*) as num from tbl_account ta inner join tbl_nfts tn on ta.id = tn.owner_id and tn.wid = 1 group by ta.id").Scan(&list)
	fmt.Println("num=", list)
}
