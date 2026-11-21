package models

import (
	"coltstail-server/utils"
	"fmt"
	"time"
)

func BSNCrete() {

	var nfts []TblNfts
	err := Eloquent.Model(&nfts).Where("wid = 11 and chain_status=2").Find(&nfts).Error
	if err == nil && len(nfts) > 0 {
		for _, v := range nfts {
			w, _ := GetWoksById(v.Wid)
			u, _ := FindAccountByID(v.OwnerId)
			var avata utils.Avata
			a, b := avata.CreateNFT(v.OperationId, u.BsnAddress, w.Name+fmt.Sprintf("#%d", v.WorkNo))
			fmt.Println("上链信息:", a, b)

			time.Sleep(time.Second * 2)
		}
	}

}
