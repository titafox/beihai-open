package models

import (
	"fmt"
	"strconv"
)

type DingList struct {
	Id    string  `json:"id"`
	Name  string  `json:"name"`
	Image string  `json:"image"`
	Price float64 `json:"price"`
}

func Ding() []DingList {

	var works []TblWorks
	err := Eloquent.Model(&TblWorks{}).Select("id,name,image").Where("status=1").Find(&works).Error
	if err != nil || len(works) == 0 {
		return make([]DingList, 0, 0)
	}

	list := make([]DingList, len(works), len(works))
	for k, v := range works {
		list[k].Id = fmt.Sprintf("%d", v.ID)
		list[k].Name = v.Name
		list[k].Image = "https://bh-ipfs.oss-cn-hangzhou.aliyuncs.com" + v.Image

		//查询价格
		var nfts TblNfts
		Eloquent.Model(&TblNfts{}).Select("price").Where("wid=? and status >0", v.ID).Order("price").First(&nfts)
		list[k].Price = nfts.Price
	}

	return list

}

func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}
