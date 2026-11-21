package models

type TblBanner struct {
	ID         int64  `json:"id" gorm:"primary_key"`
	ProducId   int    `json:"produc_id"`
	Name       string `json:"name"`
	Link       string `json:"link"`
	Image      string `json:"image"`
	BannerType int    `json:"banner_type"`
	Val        string `json:"val"`
}

func GetBannerList() ([]TblBanner, error) {
	var list []TblBanner
	err := Eloquent.Model(&TblBanner{}).Where("status = 1").Find(&list).Error
	return list, err
}
