package models

import (
	"coltstail-server/entity/res"
	"coltstail-server/utils"
	"errors"
	"fmt"
)

// 市场信息
type MarketInfo struct {
	//HasNext bool `json:"has_next"`
	Series     []Serie     `json:"series"`
	TotalCount int         `json:"total_count"`
	TotalPage  int         `json:"total_page"`
	Result     []NftDetail `json:"result"`
}

// 市场藏品信息
type MarketNftInfo struct {
	NftDetail
	Logs []MarketLog `json:"logs"`
}

// 流通记录
type MarketLog struct {
	Nickname   string  `json:"nickname"`
	TypText    string  `json:"typ_text"`
	Price      float64 `json:"price"`
	CreateTime int64   `json:"create_time"`
}

type Serie struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func GetList(page, typ, wid uint, sort string, keyword string) *MarketInfo {
	//查询品牌
	var marketInfo MarketInfo
	var series []Serie
	err := Eloquent.Raw("select id,name from tbl_works where status<2").Scan(&series).Error
	if err != nil {
		marketInfo.Series = make([]Serie, 0, 0)
	} else {
		marketInfo.Series = series
	}

	if typ == 2 {
		marketInfo.TotalPage = 0
		marketInfo.Result = make([]NftDetail, 0, 0)
		return &marketInfo
	}

	var nfts []NftDetail
	sql := "select tn.id,tn.work_no,tn.author_id,tn.price,tn.status,tn.chain_status,tw.name,tw.image from tbl_nfts tn,tbl_works tw,tbl_account ta where tn.wid = tw.id and tn.owner_id = ta.id and tn.status in(1,2)"

	// 参数绑定，防止 SQL 注入
	args := make([]interface{}, 0, 2)
	if keyword != "" {
		sql += " and tw.name like ?"
		args = append(args, "%"+keyword+"%")
	}
	if wid > 0 {
		sql += " and tn.wid = ?"
		args = append(args, wid)
	}

	if sort == "" || sort == "time_desc" {
		sql += " order by tn.update_time desc"
	} else if sort == "price_desc" {
		sql += " order by tn.price desc"
	} else if sort == "price_asc" {
		sql += " order by tn.price asc"
	}

	//查询总页数
	var totalPage int
	Eloquent.Raw(sql, args...).Scan(&nfts)

	//fmt.Println("sss=",len(nfts))
	if len(nfts)%8 == 0 {
		totalPage = len(nfts) / 8
	} else {
		totalPage = len(nfts)/8 + 1
	}
	//fmt.Println("totalPage=",totalPage)
	marketInfo.TotalCount = len(nfts)
	marketInfo.TotalPage = totalPage

	//每页数量
	n := 8
	//页码
	if page == 0 {
		page = 1
	}
	sql += fmt.Sprintf(" limit %d,%d", (int(page)-1)*n, n)
	err = Eloquent.Raw(sql, args...).Scan(&nfts).Error
	if err != nil || len(nfts) == 0 {
		marketInfo.Result = make([]NftDetail, 0, 0)
	} else {
		for k, v := range nfts {
			//获取IP方信息
			Ip, _ := GetIpInfo(v.AuthorId)
			nfts[k].AuthorName = Ip.Nickname
			nfts[k].AuthorAvatar = "https://bh-ipfs.oss-cn-hangzhou.aliyuncs.com" + Ip.Avatar

			nfts[k].Image = "https://bh-ipfs.oss-cn-hangzhou.aliyuncs.com" + v.Image
			//获取nft的hash和链上标识
		}
		marketInfo.Result = nfts
	}

	return &marketInfo

}

func GetMarketNftInfo(id, account_id int64) (*MarketNftInfo, error) {
	//查询品牌
	var marketNftInfo MarketNftInfo

	var nft NftDetail
	sql := "select tn.id,tn.wid,tn.work_no,tn.price,tn.nft_id,tn.chain_status,tn.status,tw.stock,tw.name,tw.image,tw.is_synthetic_material,ti.nickname as author_name,ti.avatar as author_avatar,ti.introduction as author_desc, tn.owner_id,ta.nickname as owner_name,ta.avatar as owner_avatar from tbl_nfts tn,tbl_works tw,tbl_account ta,tbl_ips ti where tn.wid = tw.id and tn.owner_id = ta.id and tn.author_id = ti.id and tn.id = ? and tn.status < 3"
	err := Eloquent.Raw(sql, id).Scan(&nft).Error
	if err != nil {
		return nil, errors.New("获取藏品信息失败")
	}

	nft.Image = "https://bh-ipfs.oss-cn-hangzhou.aliyuncs.com" + nft.Image
	nft.AuthorAvatar = "https://bh-ipfs.oss-cn-hangzhou.aliyuncs.com" + nft.AuthorAvatar
	nft.OwnerName = "*****"
	//3762D09AF25E5F84D3DF787EF6F1647F1BC66743A3DF74B7BF94A02CAA791221
	var tx TblTx
	Eloquent.Model(&TblTx{}).Where("work_id=? and work_no=?", nft.Wid, nft.WorkNo).Last(&tx)
	nft.Hash = utils.HideHash(tx.TxHash)

	//当前藏品是否有未支付的
	var order TblOrders
	err = Eloquent.Model(&TblOrders{}).Where("status = 1 and nft_id = ? and account_id = ?", id, account_id).First(&order).Error
	if err == nil {
		nft.PayUrl = order.PayUrl
	}
	//nft.NftId = nft.

	//获取铸件信息
	work, _ := GetWoksById(nft.Wid)

	//获取流通量
	var circulation int
	Eloquent.Model(&TblNfts{}).Where("wid=?  and status < 3 and nft_id !=''", nft.Wid).Count(&circulation)
	nft.Circulation = circulation

	marketNftInfo.NftDetail = nft

	//流转信息
	var logs []MarketLog
	err = Eloquent.Raw(`
select ta.nickname,tml.price,
       case tml.typ
           when 1 then '寄售'
           when 2 then '买入'
           end as 'typ_text',ta.create_time
from tbl_account ta,tbl_market_log tml where ta.id=tml.account_id and tml.nft_id=?
order by tml.id desc
`, nft.Id).Scan(&logs).Error
	if err != nil || len(logs) == 0 {
		marketNftInfo.Logs = []MarketLog{{Nickname: nft.AuthorName, TypText: "铸件", Price: 0.00, CreateTime: work.CreateTime}}
	} else {
		for k, _ := range logs {
			logs[k].Nickname = "*****"
		}
		logs = append(logs, MarketLog{Nickname: nft.AuthorName, TypText: "铸件", Price: 0.00, CreateTime: work.CreateTime})
		marketNftInfo.Logs = logs
	}

	return &marketNftInfo, nil
}

// 已购藏品
func GetMyNFTs(account_id int64, page uint) []res.Nfts {
	fmt.Println("account_id", account_id)
	var nfts []res.Nfts
	sql := "select count(tn.work_no) as nums,tn.wid as id, tw.name,tw.image,tw.is_synthetic_material,ti.nickname as author_name,ti.avatar as author_avatar from tbl_nfts tn,tbl_works tw,tbl_ips ti where tn.owner_id=? and tn.status < 3 and tn.wid= tw.id and tn.author_id=ti.id group by tn.wid,author_name,author_avatar"
	err := Eloquent.Raw(sql, account_id).Scan(&nfts).Error
	if err != nil || len(nfts) == 0 {
		return make([]res.Nfts, 0, 0)
	}

	for k, v := range nfts {
		nfts[k].Image = "https://bh-ipfs.oss-cn-hangzhou.aliyuncs.com" + v.Image
		nfts[k].AuthorAvatar = "https://bh-ipfs.oss-cn-hangzhou.aliyuncs.com" + v.AuthorAvatar
	}

	return nfts
}

// 已购藏品系列列表
func GetMyNFTsSeries(id, account_id int64, page uint) res.SeriesNftsList {
	var sn res.SeriesNftsList
	var nfts []res.SeriesNfts
	sql := "select tn.id, tn.work_no,tn.price,tn.chain_status, tw.name,tw.image,tw.is_synthetic_material,ti.nickname as author_name,ti.avatar as author_avatar from tbl_nfts tn,tbl_works tw,tbl_ips ti where tn.wid=? and tn.owner_id=? and tn.status < 3  and tn.wid= tw.id and tn.author_id=ti.id order by tn.update_time desc"

	var total struct {
		Nums int
	}
	sqlTotal := "select count(tn.id) as nums from tbl_nfts tn,tbl_works tw,tbl_ips ti where tn.wid=? and tn.owner_id=? and tn.status < 3 and tn.wid= tw.id and tn.author_id=ti.id order by tn.update_time desc"
	Eloquent.Raw(sqlTotal, id, account_id).Scan(&total)

	if total.Nums%8 == 0 {
		sn.TotalPage = total.Nums / 8
	} else {
		sn.TotalPage = total.Nums/8 + 1
	}
	//sn.TotalPage = total.Nums

	//每页数量
	n := 8
	//页码
	if page == 0 {
		page = 1
	}
	sql += fmt.Sprintf(" limit %d,%d", (int(page)-1)*n, n)

	err := Eloquent.Raw(sql, id, account_id).Scan(&nfts).Error

	if err != nil || len(nfts) == 0 {
		sn.List = make([]res.SeriesNfts, 0, 0)
		return sn
	}

	for k, v := range nfts {
		nfts[k].Image = "https://bh-ipfs.oss-cn-hangzhou.aliyuncs.com" + v.Image
		nfts[k].AuthorAvatar = "https://bh-ipfs.oss-cn-hangzhou.aliyuncs.com" + v.AuthorAvatar
	}
	sn.List = nfts
	return sn
}

// 流通列表
func GetWorksNFT(wid int64, sort string, page uint) res.SeriesNftsList {
	var sn res.SeriesNftsList
	var nfts []res.SeriesNfts
	var work TblWorks
	Eloquent.Model(&TblWorks{}).Select("name").Where("id=?", wid).First(&work)
	sn.Name = work.Name
	sql := "select tn.id, tn.work_no,tn.price,tn.status,tn.chain_status, tw.name,tw.image,ti.nickname as author_name,ti.avatar as author_avatar from tbl_nfts tn,tbl_works tw,tbl_ips ti where tn.wid=? and tn.wid= tw.id and tn.author_id=ti.id and tn.status>0"
	var total struct {
		Nums int
	}
	sqlTotal := "select count(tn.id) as nums from tbl_nfts tn,tbl_works tw,tbl_ips ti where tn.wid=? and tn.wid= tw.id and tn.author_id=ti.id and tn.status>0"
	Eloquent.Raw(sqlTotal, wid).Scan(&total)

	if total.Nums%8 == 0 {
		sn.TotalPage = total.Nums / 8
	} else {
		sn.TotalPage = total.Nums/8 + 1
	}
	//sn.TotalPage = total.Nums

	if sort == "" || sort == "price" {
		sql += " order by tn.price"
	} else {
		sql += " order by tn.work_no"
	}

	//每页数量
	n := 8
	//页码
	if page == 0 {
		page = 1
	}
	sql += fmt.Sprintf(" limit %d,%d", (int(page)-1)*n, n)

	err := Eloquent.Raw(sql, wid).Scan(&nfts).Error

	if err != nil || len(nfts) == 0 {
		sn.List = make([]res.SeriesNfts, 0, 0)
		return sn
	}

	for k, v := range nfts {
		nfts[k].Image = "https://bh-ipfs.oss-cn-hangzhou.aliyuncs.com" + v.Image
		nfts[k].AuthorAvatar = "https://bh-ipfs.oss-cn-hangzhou.aliyuncs.com" + v.AuthorAvatar
	}
	sn.List = nfts
	return sn
}

// 我的寄售
func GetMySellsNFTs(account_id int64, page uint) []res.SeriesNfts {
	var nfts []res.SeriesNfts
	sql := "select tn.id, tn.work_no,tn.price,tn.chain_status,tn.status, tw.name,tw.image,ti.nickname as author_name,ti.avatar as author_avatar from tbl_nfts tn,tbl_works tw,tbl_ips ti where tn.owner_id=? and tn.status > 0 and tn.wid= tw.id and tn.author_id=ti.id order by tn.update_time desc"

	//每页数量
	n := 8
	//页码
	if page == 0 {
		page = 1
	}
	sql += fmt.Sprintf(" limit %d,%d", (int(page)-1)*n, n)

	err := Eloquent.Raw(sql, account_id).Scan(&nfts).Error

	if err != nil || len(nfts) == 0 {
		return make([]res.SeriesNfts, 0, 0)
	}
	for k, v := range nfts {
		nfts[k].Image = "https://bh-ipfs.oss-cn-hangzhou.aliyuncs.com" + v.Image
		nfts[k].AuthorAvatar = "https://bh-ipfs.oss-cn-hangzhou.aliyuncs.com" + v.AuthorAvatar
	}
	return nfts
}
