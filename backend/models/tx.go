package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type TblTx struct {
	ID          int64  `json:"id" gorm:"primary_key"`
	FromId      int64  `json:"from_id"`
	ToId        int64  `json:"to_id"`
	WorkId      int64  `json:"work_id"`
	WorkNo      int64  `json:"work_no"`
	BlockHeight int    `json:"block_height"`
	TxHash      string `json:"tx_hash"`
	Typ         string `json:"typ"`
	Timestamp   int64  `json:"timestamp"`
}

func TxHook(operation_id string, block_height int, hash string, typ string, nft_id string) {
	//transfer_nft

	// 幂等：同一笔链上交易只处理一次（防止回调重放）
	var seen int
	Eloquent.Model(&TblTx{}).Where("tx_hash = ?", hash).Count(&seen)
	if hash == "" || seen > 0 {
		fmt.Println("bsn: duplicate or empty tx_hash, skip", operation_id, hash)
		return
	}

	//查询作品
	var nft TblNfts
	err := Eloquent.Model(&TblNfts{}).Where("operation_id = ?", operation_id).First(&nft).Error
	if err != nil {
		fmt.Println("bsn:", err)
		return
	}

	//查询空投是否存在此NFT
	var from_id, to_id int64
	if typ == "transfer_nft" {

	}

	//更新上链数据
	var txDb = Eloquent
	err = txDb.Transaction(func(tx *gorm.DB) error {
		//增加链上数据
		if txErr := tx.Create(&TblTx{
			FromId:      from_id,
			ToId:        to_id,
			WorkId:      nft.Wid,
			WorkNo:      nft.WorkNo,
			BlockHeight: block_height,
			TxHash:      hash,
			Typ:         typ,
			Timestamp:   time.Now().Unix(),
		}).Error; txErr != nil {
			fmt.Println("bsn:", txErr)
			return txErr
		}

		//更新nfts的上链状态
		if txErr := tx.Model(TblNfts{}).Where("id=?", nft.ID).
			Updates(map[string]interface{}{
				"transfer_num": nft.TransferNum + 1,
				"nft_id":       nft_id,
				"chain_status": 1,
				"update_time":  time.Now().Unix(),
			}).Error; txErr != nil {
			fmt.Println("bsn:", txErr)
			return txErr
		}

		// 返回 nil 提交事务
		return nil
	})
}
