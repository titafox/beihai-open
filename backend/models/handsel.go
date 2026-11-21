package models

import (
	"coltstail-server/utils"
	"errors"
	"fmt"
	"time"
)

type TblHandsel struct {
	ID          int64  `json:"id" gorm:"primary_key"`
	AccountId   int64  `json:"account_id"`
	WId         int64  `json:"w_id"`
	WorkNo      int64  `json:"work_no"`
	Typ         uint   `json:"typ"`
	Status      int    `json:"status"`
	OperationId string `json:"operation_id"`
	CreteTime   int64  `json:"crete_time"`
}

// 转增
func SetHandsel(account_id, id int64, phone string) error {
	//查询发送方
	faccount, err := FindAccountByID(account_id)
	if err != nil {
		return errors.New("发送方不存在")
	}
	if faccount.Phone == phone {
		return errors.New("不能自己转给自己")
	}

	//查询是否有转增权限
	if faccount.IsHandsel != 1 {
		return errors.New("暂无转赠权限，请联系客服处理")
	}

	//查询用户是否存在此作品
	var nft TblNfts
	err = Eloquent.Model(&TblNfts{}).Where("owner_id = ? AND id = ?", account_id, id).First(&nft).Error
	if err != nil {
		return errors.New("您不存在此NFT")
	}

	if nft.Wid > 2 && time.Now().Unix()-nft.UpdateTime < 7200 {
		return errors.New("转赠冷却中")

	}

	if nft.ChainStatus != 1 {
		return errors.New("NFT正在打包上链中,请先等待处理完成")
	}
	//查询对方是否存在
	taccount, err := FindAccountByPhone(phone)
	if err != nil {
		return errors.New("接收方不存在")
	}

	op := utils.GenerateUri("handsel")

	var han TblHandsel
	han.AccountId = account_id
	han.WId = nft.Wid
	han.WorkNo = nft.WorkNo
	han.Typ = 1
	han.Status = 0
	han.OperationId = op
	han.CreteTime = time.Now().Unix()
	err = Eloquent.Create(&han).Error
	if err != nil {
		return errors.New("转增失败")
	}

	var han2 TblHandsel
	han2.AccountId = taccount.ID
	han2.WId = nft.Wid
	han2.WorkNo = nft.WorkNo
	han2.Typ = 0
	han2.Status = 0
	han2.OperationId = op
	han2.CreteTime = time.Now().Unix()
	err = Eloquent.Create(&han2).Error
	if err != nil {
		return errors.New("转增失败")
	}

	var workDB = Eloquent
	if txErr := workDB.Model(TblNfts{}).Where("id = ?", nft.ID).
		Updates(map[string]interface{}{
			"owner_id":     taccount.ID,
			"operation_id": op,
			"chain_status": 0,
		}).Error; txErr != nil {
		return errors.New("更新NFT信息失败")
	}

	var avata utils.Avata
	bol, _ := avata.NftTransfer(faccount.BsnAddress, nft.NftId, taccount.BsnAddress, op)
	if !bol {
		return errors.New("转增：链上操作失败")
	}

	return err

}

type HandselListRes struct {
	Name       string `json:"name"`
	Image      string `json:"image"`
	No         string `json:"no"`
	Mobile     string `json:"mobile"`
	Des        string `json:"des"`
	CreateTime int64  `json:"create_time"`
}

func GetHandselList(account_id int64) []HandselListRes {
	// todo 需要加上编号
	var s []TblHandsel
	err := Eloquent.Model(&TblHandsel{}).Where("account_id = ?", account_id).Find(&s).Error
	if err != nil {
		return make([]HandselListRes, 0, 0)
	}

	list := make([]HandselListRes, len(s), len(s))
	for k, v := range s {
		//获取作品信息
		var work TblWorks
		Eloquent.Model(&TblWorks{}).Where("id = ?", v.WId).First(&work)
		list[k].Name = work.Name
		list[k].Image = "https://bh-ipfs.oss-cn-hangzhou.aliyuncs.com" + work.Image
		list[k].No = fmt.Sprintf("#%d", v.WorkNo)

		//获取对方信息
		msg := ""
		var h TblHandsel
		err = Eloquent.Model(&TblHandsel{}).Where("operation_id = ? and typ != ?", v.OperationId, v.Typ).First(&h).Error

		if err == nil {
			//获取用户信息
			account, _ := FindAccountByID(h.AccountId)
			list[k].Mobile = account.Phone
			if v.Typ == 0 {
				msg = fmt.Sprintf("已成功收到%s的转增", account.Phone)
			} else {
				msg = fmt.Sprintf("已成功转让给%s", account.Phone)
			}
		}
		list[k].Des = msg
		list[k].CreateTime = v.CreteTime
	}

	return list
	//list[0].Name = "极光"
	//list[0].Image = "https://lihood-1306623008.cos.ap-nanjing.myqcloud.com/products/20221006/d7cd0637-4542-11ed-a0a6-525400ec5ffb.jpeg?imageMogr2/thumbnail/!50p"
	//list[0].No = "#123/10000"
	//list[0].Mobile = "18888888"
	//list[0].CreateTime = time.Now().Unix()
}
