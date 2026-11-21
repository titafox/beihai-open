package models

import (
	"coltstail-server/entity/res"
	"fmt"
	"time"
)

type TblAccount struct {
	ID             int64  `json:"id" gorm:"primary_key"`
	Uid            string `json:"uid"`
	Pid            int64  `json:"pid"`
	Role           int    `json:"role"`
	Status         int    `json:"status"`
	Phone          string `json:"phone"`
	Nickname       string `json:"nickname"`
	Avatar         string `json:"avatar"`
	BsnAddress     string `json:"bsn_address"`
	BsnUsername    string `json:"bsn_username"`
	Name           string `json:"name"`
	Amount         int    `json:"amount"`
	IdCardNum      string `json:"id_card_num"`
	Authentication int    `json:"authentication"`
	Description    string `json:"description"`
	IsHandsel      uint   `json:"is_handsel"`
	IsWallet       uint   `json:"is_wallet"`
	Nums           int    `json:"nums"`
	CreateTime     int64  `json:"create_time"`
	UpdateTime     int64  `json:"update_time"`
}

// 创建用户
func (account *TblAccount) Insert() (int64, error) {
	err := Eloquent.Create(account).Error
	if err != nil {
		return 0, err
	}
	return account.ID, nil
}

// 更新实名信息
func (account *TblAccount) UpdateAccountAuthentication(id int64) error {
	var accDb = Eloquent
	err := accDb.Model(&account).Where("id = ?", id).Updates(account).Error
	if err != nil {
		return err
	}

	//更新邀请数量
	//var a TblAccount
	//err = Eloquent.Model(&TblAccount{}).Where("id=?", id).First(&a).Error
	//if err == nil {
	//	if a.Pid > 0 {
	//		var b TblAccount
	//		err = Eloquent.Model(&TblAccount{}).Where("id=?", a.Pid).First(&b).Error
	//		if err == nil {
	//			Eloquent.Model(&TblAccount{}).Where("id=?", b.ID).UpdateColumn("nums", gorm.Expr("nums+?", 1))
	//		}
	//	}
	//}

	return nil
}

// 查询邀请码
func (account *TblAccount) CheckInvitation(invitation string) (*TblAccount, error) {
	var accountDb = Eloquent
	var s TblAccount
	err := accountDb.Model(&TblAccount{}).Where("uid = ?", invitation).First(&s).Error
	return &s, err
}

// 通过手机号查询用户
func (account *TblAccount) FindAccountByPhone(phone string) (*TblAccount, error) {
	var accountDb = Eloquent
	var s TblAccount
	err := accountDb.Model(&TblAccount{}).Where("phone = ?", phone).First(&s).Error
	return &s, err
}

// 通过手机号查询用户
func FindAccountByPhone(phone string) (*TblAccount, error) {
	var accountDb = Eloquent
	var s TblAccount
	err := accountDb.Model(&TblAccount{}).Where("phone = ?", phone).First(&s).Error
	return &s, err
}

// 通过ID查询用户
func FindAccountByID(id int64) (*TblAccount, error) {
	var accountDb = Eloquent
	var s TblAccount
	err := accountDb.Model(&TblAccount{}).Where("id = ?", id).First(&s).Error
	return &s, err
}

// 通过身份证查找用户
func FindAccountByCard(card_no string) (*TblAccount, error) {
	var accountDb = Eloquent
	var s TblAccount
	err := accountDb.Model(&TblAccount{}).Select("phone").Where("id_card_num = ?", card_no).First(&s).Error
	return &s, err
}

// 更新用户信息
func UpdateAccountInfo(nickname, avatar, desc string, accountId int64) error {

	if txErr := Eloquent.Model(TblAccount{}).Where("id=?", accountId).
		Updates(map[string]interface{}{
			"nickname":    nickname,
			"avatar":      avatar,
			"description": desc,
		}).Error; txErr != nil {
		return txErr
	}

	return nil
}

// 获取用户可购买的数量
func GetMaxNum(account_id, work_id int64) uint {
	var num uint

	//获取藏品信息
	work, err := GetWoksById(work_id)
	if err != nil {
		return 0
	}

	//是否有白名单
	whiteInfo, err := GetSanpShotLogByOne(account_id, work_id)
	if err != nil { //没有白名单
		//是否已到公售时间
		if work.PubTime > time.Now().Unix() { //未到
			num = 0
		} else { //已到
			//查询是否有正在进行中/已完成的订单
			var orderCount int64
			Eloquent.Model(&TblOrders{}).Where("account_id=? AND wid=? AND status<3", account_id, work_id).Count(&orderCount)
			if orderCount == 0 {
				num = 1
			}
		}
	} else { //todo 有白名单 有点问题没考虑预售期没有买到最大量的问题
		//是否已到预售时间
		if work.PubTime-work.AdvanceHour*3600 > time.Now().Unix() { //未到
			num = 0
		} else { //已到
			//查询是否有正在进行中/已完成的订单
			var order TblOrders
			Eloquent.Model(&TblOrders{}).Select("sum(num) as num").Where("account_id=? AND wid=? AND status<3", account_id, work_id).First(&order)
			//如果白名单最大份数等于正在进行中/已完成的份数，那么他不能再买了
			if work.PubTime > time.Now().Unix() { //还未到公售时间
				num = uint(whiteInfo.BuyTotal) - uint(order.Num)
			} else { //已经到了公售时间
				//查询在公售期间是否已经买过1件
				var o TblOrders
				err = Eloquent.Model(&TblOrders{}).Where("account_id=? AND wid=? AND status<3 and create_time > ?", account_id, work_id, work.PubTime).First(&o).Error
				if err != nil { //公售期没有买过
					if order.Num == whiteInfo.BuyTotal {
						num = 0
					} else {
						num = 1
					}
				} else { //公售期买过
					num = 0
				}
			}
		}
	}
	return num
}

func GetMyInviteList(account_id int64, page uint) res.InviteList {
	var inviteList res.InviteList
	//更新我的邀请
	var list []res.Invite
	//1668754800 1668765600
	sql := "select ta.nickname,ta.create_time from tbl_account ta,tbl_nfts tn  where ta.pid = ? and ta.id= tn.owner_id and tn.update_time > 1668765600 group by tn.owner_id"
	var counts int
	err := Eloquent.Raw(sql, account_id).Scan(&list).Error
	if err != nil && len(list) == 0 {
		inviteList.List = make([]res.Invite, 0, 0)
		return inviteList
	}

	counts = len(list)
	if counts > 0 { //更新下人数
		Eloquent.Model(&TblAccount{}).Where("id=?", account_id).UpdateColumn(TblAccount{Nums: counts})
	}
	//Eloquent.Model(&TblAccount{}).Where("pid=?",account_id).Count(&counts)
	inviteList.Nums = uint(counts)
	if counts%8 == 0 {
		inviteList.TotalPage = int(counts) / 8
	} else {
		inviteList.TotalPage = int(counts)/8 + 1
	}

	//var list []res.Invite
	//sql := "select nickname,create_time from tbl_account where pid = ? order by id desc"

	//每页数量
	n := 15
	//页码
	if page == 0 {
		page = 1
	}
	sql += fmt.Sprintf(" limit %d,%d", (int(page)-1)*n, n)

	err = Eloquent.Raw(sql, account_id).Scan(&list).Error
	if err != nil && len(list) == 0 {
		inviteList.List = make([]res.Invite, 0, 0)
		return inviteList
	}

	inviteList.List = list
	return inviteList

}

func InviteTop() []res.InviteTop {
	var accounts []TblAccount
	err := Eloquent.Model(&TblAccount{}).Select("nickname,nums").Where("nums >0").Order("nums desc").Limit(10).Find(&accounts).Error

	if err != nil && len(accounts) == 0 {
		return make([]res.InviteTop, 0, 0)
	}

	list := make([]res.InviteTop, len(accounts), len(accounts))
	for k, v := range accounts {
		list[k].Nickname = v.Nickname
		list[k].Nums = v.Nums
	}

	return list

}
