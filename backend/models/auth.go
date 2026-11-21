package models

type Auth struct {
	ID       int64  `json:"id" gorm:"primary_key"`
	AuthType uint   `json:"auth_type"`
	Address  string `json:"address"`
	Status   uint   `json:"status"`
}

func GetIsUploadByAddress(address string, authType uint) (*Auth, error) {
	var auth Auth
	err := Eloquent.Model(&Auth{}).Where("address = ? AND auth_type = ? AND status = 1", address, authType).First(&auth).Error
	return &auth, err
}

func SetAuth(address string) error {
	var auth Auth
	err := Eloquent.Model(&Auth{}).Where("address = ? AND status = 1", address).First(&auth).Error
	if err == nil {
		return nil
	}

	//开始授权
	if txErr := Eloquent.Create(&Auth{
		AuthType: 0,
		Address:  address,
		Status:   1,
	}).Error; txErr != nil {
		return txErr
	}

	if txErr := Eloquent.Create(&Auth{
		AuthType: 1,
		Address:  address,
		Status:   1,
	}).Error; txErr != nil {
		return txErr
	}

	return nil

}
