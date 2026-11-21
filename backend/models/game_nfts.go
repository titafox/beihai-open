package models

import "time"

type GameNfts struct {
	ID              int64     `json:"id" gorm:"primary_key"`
	ContractAddress string    `json:"contract_address"`
	TokenId         string    `json:"token_id"`
	OwnerAddress    string    `json:"owner_address"`
	Status          uint      `json:"status"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func (gameNft *GameNfts) Insert() error {
	return Eloquent.Create(gameNft).Error
}

func UpdateGameNftById(id int64, data map[string]interface{}) error {
	//var userDb = Eloquent
	return Eloquent.Model(GameNfts{}).Where("id = ?", id).Updates(data).Error
}

func GetGameNft(contract_address, ownerAddress string, status uint) (*GameNfts, error) {
	var nft GameNfts
	err := Eloquent.Model(&nft).Select("id,token_id").Where(" token_id > 10 AND contract_address = ? AND owner_address = ? AND status = ? ", contract_address, ownerAddress, status).Find(&nft).Error
	return &nft, err
}

func GetGameNftByTokenId(contract_address, token_id string) (GameNfts, error) {
	var nft GameNfts
	err := Eloquent.Model(&nft).Select("id,token_id").Where("contract_address = ? AND token_id = ? AND status=1", contract_address, token_id).Find(&nft).Error
	return nft, err
}
