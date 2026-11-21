package models

import (
	"fmt"
	"time"
)

type TblWorks struct {
	ID          int64   `json:"id" gorm:"primary_key"`
	Collection  string  `json:"collection" comment:"合集标识"`
	No          int     `json:"no" comment:"编号"`
	AuthorId    int64   `json:"author_id"`
	OwnerId     int64   `json:"owner_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	TokenId     string  `json:"token_id"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Reserved    int     `json:"reserved"`
	StockSell   int     `json:"stock_sell"`
	PubTime     int64   `json:"pub_time"`
	AdvanceHour int64   `json:"advance_hour"`
	Hash        string  `json:"hash"`
	IsAirdrop   int     `json:"is_airdrop"`
	SoldOut     bool    `json:"sold_out"`
	Status      int     `json:"status"`
	SellTime    int64   `json:"sell_time"`
	CreateTime  int64   `json:"create_time"`
	UpdateTime  int64   `json:"update_time"`
}

type WorkList struct {
	Count  int    `json:"count"`
	Result []Work `json:"result"`
}

type Work struct {
	Id           int64   `json:"id"`
	Name         string  `json:"name"`
	Image        string  `json:"image"`
	Stock        int     `json:"stock"`
	AuthorName   string  `json:"author_name"`
	AuthorAvatar string  `json:"author_avatar"`
	Price        float64 `json:"price"`
	PubTime      int64   `json:"pub_time"`
	AdvanceHour  int64   `json:"advance_hour"`
	Countdown    int64   `json:"countdown"`
	SoldOut      bool    `json:"sold_out"`
	ChainStatus  uint    `json:"chain_status"`
}

type WorkDesc struct {
	Work
	AuthorDesc  string `json:"author_desc"`
	Description string `json:"description"`
	Hash        string `json:"hash"`
	TokenId     string `json:"token_id"`
	MaxSupply   uint   `json:"max_supply" comment:"最大可购买的数据"`
	CanBuy      bool   `json:"can_buy" comment:"false:不可购买,true:可购买"`
	SoldOut     bool   `json:"sold_out" comment:"false:未售罄,true:已售罄"`
}

type NftDetail struct {
	Id           int64   `json:"id" comment:"藏品id"`
	Wid          int64   `json:"wid"`
	Name         string  `json:"name" comment:"藏品名称"`
	Image        string  `json:"image" comment:"藏品图片"`
	WorkNo       int64   `json:"work_no" comment:"作品编号"`
	Stock        int     `json:"stock" comment:"发行总量"`
	Circulation  int     `json:"circulation" comment:"流通量"`
	AuthorId     int64   `json:"author_id"`
	AuthorName   string  `json:"author_name" commnet:"创造者名称"`
	AuthorAvatar string  `json:"author_avatar" comment:"创作者图像"`
	AuthorDesc   string  `json:"author_desc" comment:"创作者描述"`
	OwnerId      int64   `json:"owner_id" commnet:"拥有者id"`
	OwnerName    string  `json:"owner_name" commnet:"拥有者名称"`
	OwnerAvatar  string  `json:"owner_avatar" comment:"拥有者图像"`
	Price        float64 `json:"price" comment:"藏品价格"`
	Status       uint    `json:"status" comment:"状态"`
	ChainStatus  uint    `json:"chain_status" comment:"上链状态"`
	Hash         string  `json:"hash" comment:"交易hash"`
	NftId        string  `json:"nft_id" comment:"链上状态"`
	PayUrl       string  `json:"pay_url"`
	IsSyntheticMaterial uint `json:"is_synthetic_material"`
}

func GetWorkList(account_id int64) *WorkList {
	var works []Work
	err := Eloquent.Raw("select tw.id,tw.name,tw.image,ti.nickname as author_name,ti.avatar as author_avatar,tw.price,tw.pub_time,tw.advance_hour,tw.sold_out from tbl_works tw join tbl_ips ti on tw.author_id = ti.id and tw.status=1 order by tw.id desc").Scan(&works).Error
	if err != nil {
		return nil
	}

	var workList WorkList
	for k, v := range works {
		works[k].Image = "https://bh-ipfs.oss-cn-hangzhou.aliyuncs.com" + v.Image
		works[k].AuthorAvatar = "https://bh-ipfs.oss-cn-hangzhou.aliyuncs.com" + v.AuthorAvatar
		//查询是否有白名单
		_, err2 := GetSanpShotLogByOne(account_id, works[k].Id)
		fmt.Println("err2=", err2)
		if err2 != nil { //没有白名单
			if v.PubTime-time.Now().Unix() > 0 {
				works[k].Countdown = v.PubTime - time.Now().Unix()
			} else {
				works[k].Countdown = 0
			}
		} else { //有白名单
			if v.PubTime-v.AdvanceHour*3600-time.Now().Unix() > 0 {
				works[k].Countdown = v.PubTime - v.AdvanceHour*3600 - time.Now().Unix()
			} else {
				works[k].Countdown = 0
			}
		}
	}
	workList.Count = len(works)
	workList.Result = works
	return &workList
}

func GetMyNftList(uid int64) *WorkList {
	var nfts []TblNfts
	err := Eloquent.Model(&TblNfts{}).Where("owner_id=?", uid).Find(&nfts).Error
	if err != nil {
		return nil
	}

	var workList WorkList
	work := make([]Work, len(nfts), len(nfts))
	for k, v := range nfts {
		//获取作品详情
		var workInfo TblWorks
		Eloquent.Model(&TblWorks{}).Where("id = ?", v.Wid).First(&workInfo)

		work[k].Id = v.ID
		work[k].Name = workInfo.Name + "#" + fmt.Sprintf("%d", v.WorkNo)
		work[k].Image = "https://bh-ipfs.oss-cn-hangzhou.aliyuncs.com" + workInfo.Image
		//work[k].AuthorId = v.OwnerId
		//获取用户信息
		var user TblAccount
		err = Eloquent.Model(&TblAccount{}).Where("id=?", v.OwnerId).First(&user).Error
		if err == nil {
			work[k].AuthorName = user.Nickname
			work[k].AuthorAvatar = user.Avatar
		}
		//work[k].ChainStatus = v.ChainStatus
		work[k].Price = workInfo.Price
		work[k].Countdown = 0
		work[k].ChainStatus = v.ChainStatus

		//work[k].CanBuy = false
		//work[k].SaleTime = 0
	}
	workList.Count = len(nfts)
	workList.Result = work
	return &workList
}

// 作品详情
func GetWorkDetail(id, account_id int64) *WorkDesc {
	//作品信息
	var work TblWorks
	err := Eloquent.Model(&TblWorks{}).Where("id=?", id).First(&work).Error
	if err != nil {
		return nil
	}

	//获取IP方信息
	var ips TblIps
	err = Eloquent.Model(&TblIps{}).Where("id=?", work.AuthorId).First(&ips).Error
	if err != nil {
		return nil
	}

	fmt.Println("ips=", ips)

	var des WorkDesc
	des.Id = work.ID
	des.Name = work.Name
	des.Stock = work.Stock
	des.Image = "https://bh-ipfs.oss-cn-hangzhou.aliyuncs.com" + work.Image
	des.Description = work.Description
	des.AuthorName = ips.Nickname
	des.AuthorAvatar = "https://bh-ipfs.oss-cn-hangzhou.aliyuncs.com" + ips.Avatar
	des.AuthorDesc = ips.Introduction
	des.Price = work.Price
	des.Description = work.Description
	des.MaxSupply = GetMaxNum(account_id, work.ID)
	des.SoldOut = work.SoldOut
	//查询是否有白名单
	_, err2 := GetSanpShotLogByOne(account_id, work.ID)
	if err2 != nil { //没有白名单
		if work.PubTime-time.Now().Unix() > 0 {
			des.Countdown = work.PubTime - time.Now().Unix()
		} else {
			des.Countdown = 0
		}
	} else { //有白名单
		if work.PubTime-work.AdvanceHour*3600-time.Now().Unix() > 0 {
			des.Countdown = work.PubTime - work.AdvanceHour*3600 - time.Now().Unix()
		} else {
			des.Countdown = 0
		}
	}
	return &des
}

// 藏品详情
func GetMyNftDetail(id, account_id int64) (*NftDetail, error) {
	//获取藏品详情
	var nft TblNfts
	err := Eloquent.Model(TblNfts{}).Where("id=? and owner_id=?", id, account_id).First(&nft).Error
	if err != nil {
		return nil, err
	}

	fmt.Println("nft=", nft)

	//获取IP方详情
	var ips TblIps
	err = Eloquent.Model(TblIps{}).Where("id=?", nft.AuthorId).First(&ips).Error
	if err != nil {
		return nil, err
	}

	//获取作品详情
	var work TblWorks
	err = Eloquent.Model(TblWorks{}).Where("id=?", nft.Wid).First(&work).Error
	if err != nil {
		return nil, err
	}
	var nftDes NftDetail
	nftDes.Id = nft.ID
	nftDes.Name = work.Name + fmt.Sprintf("#%d", nft.WorkNo)
	nftDes.Image = "https://bh-ipfs.oss-cn-hangzhou.aliyuncs.com" + work.Image
	nftDes.WorkNo = nft.WorkNo
	nftDes.Stock = work.Stock
	nftDes.AuthorName = ips.Nickname
	nftDes.AuthorAvatar = "https://bh-ipfs.oss-cn-hangzhou.aliyuncs.com" + ips.Avatar
	nftDes.AuthorDesc = ips.Introduction
	nftDes.ChainStatus = nft.ChainStatus
	nftDes.Hash = ""
	nftDes.NftId = ""
	return &nftDes, nil
}

// 通过id获取作品详情
func GetWoksById(id int64) (*TblWorks, error) {
	var work TblWorks
	err := Eloquent.Model(&TblWorks{}).Where("id = ?", id).First(&work).Error
	return &work, err
}

type Sa struct {
	Id           int64   `json:"id" comment:"藏品id"`
	Name         string  `json:"name" comment:"藏品名称"`
	Image        string  `json:"image" comment:"藏品图片"`
	AuthorName   string  `json:"author_name" comment:"创作者名称"`
	AuthorAvatar string  `json:"author_avatar" comment:"创作者图像"`
	Price        float64 `json:"price" comment:"价格"`
	Countdown    int64   `json:"countdown" comment:"倒计时秒数"`
}

func TestWork() []Sa {
	var works []Sa
	Eloquent.Raw("select tn.id,tn.wid,owner_id,ti.nickname,tw.name,tw.image from tbl_works tw, tbl_nfts tn,tbl_ips ti where tn.transfer_num=0 and tn.work_no=0 and tn.wid = tw.id and tn.owner_id = ti.id").Scan(&works)
	return works
}
