package models

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type TblNfts struct {
	ID          int64   `json:"id" gorm:"primary_key"`
	Wid         int64   `json:"wid" comment:"作品id"`
	WorkNo      int64   `json:"work_no" comment:"作品编号"`
	AuthorId    int64   `json:"author_id"`
	OwnerId     int64   `json:"owner_id"`
	TransferNum int     `json:"transfer_num"`
	Price       float64 `json:"price"`
	NftId       string  `json:"nft_id"`
	OperationId string  `json:"operation_id"`
	ChainStatus uint    `json:"chain_status"`
	Status      uint    `json:"status"`
	UpdateTime  int64   `json:"update_time"`
}

func (nft *TblNfts) Insert() error {
	var nftDb = Eloquent
	return nftDb.Create(nft).Error
}

func GetNFTDetail(id int64) (TblNfts, error) {
	var nft TblNfts
	err := Eloquent.Model(&TblNfts{}).Where("id= ?", id).First(&nft).Error
	return nft, err
}

// 上架藏品
func OnTheNFT(id, account_id int64, price float64) error {
	var nft TblNfts
	err := Eloquent.Model(&TblNfts{}).Select("id,wid,status").Where("id= ? and owner_id=?", id, account_id).First(&nft).Error
	if err != nil {
		return errors.New("你不存在此藏品")
	}

	//查询可售卖时间
	work, _ := GetWoksById(nft.Wid)
	if work.SellTime > time.Now().Unix() {
		s := time.Unix(work.SellTime, 0).Format("2006-01-02 15:04:05")
		return errors.New("藏品可出售时间为:" + s)
	}

	if nft.Status == 1 {
		return errors.New("藏品已在上架状态")
	}

	if nft.Status == 2 {
		return errors.New("藏品锁单中，无法上架")
	}

	if price <= 1 {
		return errors.New("价格非法，请调整价格")
	}

	//开始更新
	err = Eloquent.Transaction(func(tx *gorm.DB) error {
		//增加挂售记录
		if txErr := tx.Create(&TblMarketLog{
			NftId:      nft.ID,
			AccountId:  account_id,
			Typ:        1,
			Price:      price,
			Status:     1,
			CreateTime: time.Now().Unix(),
		}).Error; txErr != nil {
			return txErr
		}

		txErr := tx.Model(&TblNfts{}).Where("id=?", nft.ID).UpdateColumns(TblNfts{Price: price, Status: 1, UpdateTime: time.Now().Unix()}).Error
		if txErr != nil {
			return txErr
		}
		return nil
	})

	if err != nil {
		fmt.Println("上架失败：", err)
		return errors.New("内部错误，请稍后重试")
	}

	return nil
}

// 下架藏品
func OffTheNFT(id, account_id int64) error {
	var nft TblNfts
	err := Eloquent.Model(&TblNfts{}).Select("id,status").Where("id= ? and owner_id=?", id, account_id).First(&nft).Error
	if err != nil {
		return errors.New("你不存在此藏品")
	}
	if nft.Status == 0 {
		return errors.New("藏品已在下架状态")
	}
	if nft.Status == 2 {
		return errors.New("藏品锁单中，无法下架")
	}

	//开始更新
	err = Eloquent.Transaction(func(tx *gorm.DB) error {
		txErr := tx.Model(TblNfts{}).Where("id = ?", id).Updates(map[string]interface{}{
			"status":      0,
			"update_time": time.Now().Unix(),
		}).Error
		if txErr != nil {
			return txErr
		}

		//txErr = tx.Model(&TblMarketLog{}).Where("nft_id=? and account_id =? and typ=1 and status = 1", nft.ID, account_id).UpdateColumns(TblMarketLog{Status: 0, CreateTime: time.Now().Unix()}).Error
		txErr = Eloquent.Model(TblMarketLog{}).Where("nft_id=? and account_id =? and typ=1 and status = 1", nft.ID, account_id).Updates(map[string]interface{}{
			"status":      0,
			"create_time": time.Now().Unix(),
		}).Error

		if txErr != nil {
			return txErr
		}
		return nil
	})

	if err != nil {
		fmt.Println("下失败：", err)
		return errors.New("内部错误，请稍后重试")
	}
	return nil
}

//获取某个时刻nft 数量
func GetNftCount(wid,account_id,sometime int64) uint  {
	var count uint
	Eloquent.Model(TblNfts{}).Where("wid=? and owner_id=？and update_time<=?",wid,account_id,sometime).Count(&count)
	return count
}
