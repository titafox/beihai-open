package models

type TblVersion struct {
	ID          int64  `json:"id" gorm:"primary_key"`
	DownloadUrl string `json:"download_url"`
	Oldversion  int    `json:"oldversion"`
	Newversion  int    `json:"newversion"`
	Content     string `json:"content"`
	Enforce     int    `json:"enforce"`
}

func GetVersion() (TblVersion, error) {
	var version TblVersion
	err := Eloquent.Model(&TblVersion{}).First(&version).Error
	return version, err

}
