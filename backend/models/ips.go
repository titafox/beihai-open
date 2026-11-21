package models

import "fmt"

type TblIps struct {
	ID           int64  `json:"id" gorm:"primary_key"`
	Nickname     string `json:"nickname"`
	Avatar       string `json:"avatar"`
	Introduction string `json:"introduction"`
}

type ArtListRes struct {
	Count  int            `json:"count"`
	Result []AuthorEntity `json:"result"`
}

type AuthorEntity struct {
	ID         int64  `json:"id" column:"id"`
	Nickname   string `json:"nickname" column:"nickname"`
	Avatar     string `json:"avatar" column:"avatar"`
	WorksCount int64  `json:"works_count" column:"works_count"`
}

func GetIpsAll() *ArtListRes {
	var list ArtListRes
	var ips []TblIps
	err := Eloquent.Model(&TblIps{}).Find(&ips).Error
	fmt.Println("ips====", ips)
	if err != nil {
		list.Count = 0
		list.Result = make([]AuthorEntity, 0, 0)
	} else {
		list.Count = len(ips)
		r := make([]AuthorEntity, len(ips), len(ips))
		for k, v := range ips {
			var c int64
			Eloquent.Model(&TblWorks{}).Where("author_id=?", v.ID).Count(&c)
			r[k].ID = v.ID
			r[k].Nickname = v.Nickname
			r[k].Avatar = "https://bh-ipfs.oss-cn-hangzhou.aliyuncs.com" + v.Avatar
			r[k].WorksCount = c
		}
		list.Result = r
	}

	return &list
}

func GetIpInfo(id int64) (TblIps, error) {
	var ip TblIps
	err := Eloquent.Model(&TblIps{}).Where("id = ?", id).First(&ip).Error
	return ip, err
}
