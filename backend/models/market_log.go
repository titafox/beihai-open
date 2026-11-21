package models

type TblMarketLog struct {
	ID         int64   `json:"id" gorm:"primary_key"`
	NftId      int64   `json:"nft_id"`
	AccountId  int64   `json:"account_id"`
	Typ        uint    `json:"typ"`
	Price      float64 `json:"price"`
	Status     uint    `json:"status"`
	CreateTime int64   `json:"create_time"`
}
