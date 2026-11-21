package models

import (
	"encoding/json"
	"fmt"
	"time"
)


type TblSnapshot2 struct {
	Id int64 `json:"id"`
	AId int64 `json:"a_id"`
	RoleData string `json:"role_data"`
	RoleData2 string `json:"role_data2"`
	StartTime int64 `json:"start_time"`
	Status uint `json:"status"`
	UpdateTime int64 `json:"update_time"`
}

type SnapInfo struct {
 Wid int64 `json:"wid"`
 Num uint `json:"num"`
}

func GetAllSnapshot2()  {
	var list []TblSnapshot2
	err := Eloquent.Model(TblSnapshot2{}).Where("status=0").Find(&list).Error
	if err == nil && len(list) >0 {
		for _, v := range list {
			if txErr := Eloquent.Model(TblSnapshot2{}).Where("id=?", v.Id).
				Updates(map[string]interface{}{
					"status":      1,
					"update_time": time.Now().Unix(),
				}).Error; txErr != nil {
				continue
			}

			var endpoint  []SnapInfo
			err := json.Unmarshal([]byte(v.RoleData2),&endpoint)
			if err != nil || len(endpoint) == 0 {
				fmt.Println("关联作品解析失败:",err)
				break
			}

			//获取关联的作品和数量,变为map
			var snapInfo = make(map[int64]uint,len(endpoint))
			var ids []int64
			for _,v1 := range endpoint {
				//查询用户藏品数量
				ids = append(ids, v1.Wid)
				snapInfo[v1.Wid] = v1.Num
			}
			//查询所有用户
			var users []TblAccount
			err = Eloquent.Model(&TblAccount{}).Select("id,phone").Find(&users).Error
			if err == nil && len(users) >0 {
				for _, user := range users {
					var s1 []SnapInfo
					var nums []uint
					err = Eloquent.Raw("select wid,count(*) as num from tbl_nfts where wid in (?) and owner_id=? group by wid",ids,user.ID).Scan(&s1).Error
					if err == nil && len(s1) >0 {
						if len(s1) != len(endpoint) {
							//gg := fmt.Sprintf("A用户%d最多可合成数量%d",user.ID,GetMinNum(nums))
							//fmt.Println(gg)
							continue
						}
						for _, n := range s1 {
							//可快照的数量
							snapshotNum := n.Num / snapInfo[n.Wid]
							nums = append(nums,snapshotNum)
						}

						//gg := fmt.Sprintf("B用户%d最多可合成数量%d",user.ID,GetMinNum(nums))

						//fmt.Println(gg)

						//可合成数量
						snum := GetMinNum(nums)
						if snum == 0 {
							continue
						}

						//存在则插入
						var tblSnapshotlog TblSnapshotLog
						err = Eloquent.Model(&TblSnapshotLog{}).Where("work_id=? and account_id=?", v.AId, user.ID).First(&tblSnapshotlog).Error
						if err != nil {
							tblSnapshotlog.AccountId = user.ID
							tblSnapshotlog.WorkId = v.AId
							tblSnapshotlog.SnapshotId = 0
							tblSnapshotlog.BuyTotal = int(snum)
							tblSnapshotlog.UseTotal = 0
							tblSnapshotlog.SnapshotMod = 1
							tblSnapshotlog.CreateTime = time.Now().Unix()
							tblSnapshotlog.UpdateTime = time.Now().Unix()
							_, err = tblSnapshotlog.Insert()
							if err != nil {
								fmt.Println("用户", v.AId, "增加快照失败：", err.Error())
							}
						} else { //update
							if txErr := Eloquent.Model(TblSnapshotLog{}).Where("id=?", tblSnapshotlog.ID).
								Updates(map[string]interface{}{
									"buy_total":   tblSnapshotlog.BuyTotal + int(snum),
									"update_time": time.Now().Unix(),
								}).Error; txErr != nil {
								continue
							}
						}
					}
				}
			}


			//fmt.Println("s1=",s1)
		}
	}

	//ent := time.Now().Unix() -st
	//fmt.Println("共计用时：",ent)
}


func GetMinNum(ary []uint) uint {
	if len(ary) == 0 {
		return 0
	}

	maxVal := ary[0]
	for i := 1; i < len(ary); i++ {
		if maxVal > ary[i] {
			maxVal = ary[i]
		}
	}

	return maxVal

}
