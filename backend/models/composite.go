package models

import (
	"coltstail-server/entity/req"
	"coltstail-server/entity/res"
	"coltstail-server/utils"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type TblComposite struct {
	ID             int64  `json:"id" gorm:"primary_key"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	CompositeId    int64  `json:"composite_id"`
	Stock          int  `json:"stock"`
	RemainingStock int  `json:"remaining_stock"`
	Material       string `json:"material,omitempty"`
	MaterialInfo    string `json:"material_info"`
	MaterialData    []materialData `json:"material_data"`
	FailureRate    uint   `json:"failure_rate"`
	Status         uint   `json:"status"`
}

type TblCompositeLog struct {
	ID             int64  `json:"id" gorm:"primary_key"`
	Op string `json:"op"`
	CompositeId          int64 `json:"composite_id"`
	AccountId int64 `json:"account_id"`
	Status uint `json:"status"`
	FailText string `json:"fail_text"`
	CreateTime int64 `json:"create_time"`
	UpdateTime int64 `json:"update_time"`
}

type materialData struct {
	Wid int64 `json:"wid"`
	Name string `json:"name"`
	Nums uint `json:"nums"`
}

//合成活动列表
func GetCompositeActivityList() []TblComposite {
	var list []TblComposite
	err := Eloquent.Model(&TblComposite{}).Where("status=1").Find(&list).Error
	if err != nil || len(list) == 0 {
		return make([]TblComposite,0,0)
	}

	for k, v := range list {
		var data []materialData
		_ = json.Unmarshal([]byte(v.MaterialInfo),&data)
		list[k].MaterialData = data
		list[k].Material = ""
	}

	return list
}




func GetCompositeNFTMaterial(id, account_id int64, page uint) res.SeriesNftsList {
	var sn res.SeriesNftsList
	//获取一下合成失败的卡
	var failNfts []res.SeriesNfts
	fail_sql := "select tw.id as wid,tn.id, tn.work_no,tn.price,tn.chain_status, tw.name,tw.image,ti.nickname as author_name,ti.avatar as author_avatar from tbl_nfts tn,tbl_works tw,tbl_ips ti where tw.is_synthetic_material=1 and tw.typ=3 and tn.owner_id=? and tn.chain_status=1 and tn.status=0 and tn.wid=tw.id and tn.author_id=ti.id order by tn.update_time desc"
	err2 := Eloquent.Raw(fail_sql, account_id).Scan(&failNfts).Error
	if err2 != nil || len(failNfts) == 0 {
		sn.FailCard = make([]res.SeriesNfts, 0, 0)
	}else{
		for k, v := range failNfts {
			failNfts[k].Image = "https://bh-ipfs.oss-cn-hangzhou.aliyuncs.com" + v.Image
			failNfts[k].AuthorAvatar = "https://bh-ipfs.oss-cn-hangzhou.aliyuncs.com" + v.AuthorAvatar
			failNfts[k].IsFailCard = true
		}
	}


	sn.FailCard = failNfts

	var nfts []res.SeriesNfts
	sql := "select tw.id as wid,tn.id, tn.work_no,tn.price,tn.chain_status, tw.name,tw.image,ti.nickname as author_name,ti.avatar as author_avatar from tbl_nfts tn,tbl_works tw,tbl_ips ti where tn.wid=? and tn.owner_id=? and tn.chain_status=1 and tn.status=0 and tn.wid=tw.id and tn.author_id=ti.id order by tn.update_time desc"

	var total struct {
		Nums int
	}
	sqlTotal := "select count(tn.id) as nums from tbl_nfts tn,tbl_works tw,tbl_ips ti where tn.wid=? and tn.owner_id=? and tn.chain_status=1 and tn.status=0 and tn.wid= tw.id and tn.author_id=ti.id order by tn.update_time desc"
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
		nfts[k].IsFailCard = false
	}
	sn.List = nfts
	return sn
}

func CompositeStart(account_id int64, r req.StartCompositeData) (*req.CompositeRes,error){
	//fmt.Println("ss",r.CompositeId)
	//查询合成活动是否正常
	var data TblComposite
	err := Eloquent.Model(&TblComposite{}).Where("id=?",r.CompositeId).First(&data).Error
	if err != nil {
		return nil,errors.New("合成活动不存在")
	}
	if data.Status != 1 {
		return nil,errors.New("合成活动未开始或已结束")
	}

	//获取库存
	work,_ :=GetWoksById(data.CompositeId)
	fmt.Println("work=",work.Stock,"-",work.StockSell,"=",r.CompositeId)
	if work.Stock - work.Reserved == work.StockSell {
		return nil,errors.New("已合空，暂无名额")
	}

	//
	fmt.Println("提交的材料：",r.Data)
	var mateData []materialData
	_ = json.Unmarshal([]byte(data.MaterialInfo),&mateData)
	fmt.Println("需要的材料：",getNums(5,mateData))


	// 去重，防止同一个藏品被重复计数
	var ids []int64
	seen := make(map[int64]bool)
	for _, v := range r.Data {
		for _, id := range v.Ids {
			if !seen[id] {
				seen[id] = true
				ids = append(ids, id)
			}
		}
	}
	if len(ids) == 0 {
		return nil, errors.New("未获取到材料数据")
	}


	fmt.Println("ids=",ids)
	//查询数量是否正确
	//var num int64
	//Eloquent.Model(TblNfts{}).Where("owner_id = ? and chain_status=1 and id IN(?)",account_id,ids).Count(&num)
	//if len(ids) != int(num) {
	//	return errors.New("材料数量不正确")
	//}

	//查询材料数量是否正确
	var q []struct {
		Wid int64
		Num uint
	}
	err =Eloquent.Raw("SELECT wid,count(*) as num FROM `tbl_nfts` WHERE owner_id = ? and status=0 and chain_status=1 and id IN(?) group by wid",account_id,ids).Scan(&q).Error
	if err != nil || len(q) == 0 {
		return nil,errors.New("未获取到材料数据")
	}
	//fmt.Println("q=",q)

	// 所有提交的 id 都必须属于当前用户且可用
	var owned uint
	have := make(map[int64]uint)
	for _, v := range q {
		owned += v.Num
		have[v.Wid] = v.Num
	}
	flag := owned != uint(len(ids))
	// 配方要求的每种材料都必须足量提供
	for _, m := range mateData {
		if m.Nums > 0 && have[m.Wid] != m.Nums {
			flag = true
		}
	}
	for _, v := range q {
		num := getNums(v.Wid,mateData)
		fmt.Println("num--",v.Num,"---",num)
		if  v.Num != num {
			flag = true
			break
		}
	}

	if flag {
		fmt.Println("合成材料不正确")
		return nil,errors.New("材料数量不正确")
	}
	//
	//var mateData []materialData
	//_ = json.Unmarshal([]byte(data.Material),&mateData)
	//fmt.Println("需要的材料：",mateData)
	//
	//fmt.Println("提交的材料长度：",len(r.Data))
	//fmt.Println("需要的材料长度：",len(mateData))
	//if len(r.Data) == len(mateData) { //说明没有合成卡
	//
	//}

   op := utils.GenerateUri("hcid")
	err = Eloquent.Transaction(func(tx *gorm.DB) error {
		//将所有材料销毁
		// 只能销毁自己名下、未挂售的材料，且数量必须完全一致
		res := tx.Model(&TblNfts{}).Where("id IN(?) AND owner_id = ? AND status = 0 AND chain_status = 1", ids, account_id).UpdateColumn(TblNfts{Status: 3, UpdateTime: time.Now().Unix()})
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected != int64(len(ids)) {
			return errors.New("材料状态已变化")
		}
		//增加待处理合成
		if txErr := tx.Create(&TblCompositeLog{
			Op: op,
			CompositeId: data.ID,
			AccountId: account_id,
			Status: 0,
			CreateTime: time.Now().Unix(),
			UpdateTime: time.Now().Unix(),
		}).Error; txErr != nil {
			fmt.Println("bsn:", txErr)
			return txErr
		}
		// 返回 nil 提交事务
		return nil
	})
	if err != nil {
		return nil,errors.New("合成失败")
	}
	//加入mq
	mq := RabbitMQ{Conn: OrderConn}
	err3 := mq.Publish("mint", "HC:"+op)
	if err3 != nil {
		fmt.Println("rabbitMQ 错误：" + err.Error())
	}


	var resData req.CompositeRes
	//获取合成作品
	//work,_ =GetWoksById(data.CompositeId)
	resData.Name = work.Name
	resData.Image = "https://bh-ipfs.oss-cn-hangzhou.aliyuncs.com" + work.Image

	return &resData,nil
}

//获取需要的材料数量
func getNums(wid int64,list []materialData) uint {

	fmt.Println("tttttt:",list)
	var nums uint
	for _, v := range list {
		if v.Wid == wid {
			nums = v.Nums
			break
		}
	}

	return nums
}

func GetCompositeAll()  {
	var list []TblCompositeLog
	err := Eloquent.Model(&TblCompositeLog{}).Where("status = 0").Limit(10).Find(&list).Error
	if err == nil && len(list) > 0 {
		for _, v := range list {
			var dt TblComposite
			Eloquent.Model(&TblComposite{}).Where("id=?",v.ID).First(&dt)
			var work TblWorks
			Eloquent.Model(&TblWorks{}).Where("id=?",dt.CompositeId).First(&dt)
			var count int64
			Eloquent.Model(&TblNfts{}).Where("wid = ?", dt.CompositeId).Count(&count)
			if count < 1 { //说明没有记录，需要添加编号
				nos := utils.GenerateRandomNumber(1, dt.Stock+1, dt.Stock)
				//定义一个需要插入的TblNfts切片 长度为发行总量的长途
				//nfts := make([]TblNfts,works.Stock,works.Stock)
				for _, k := range nos {
					var nfts TblNfts
					nfts.Wid = dt.CompositeId
					nfts.WorkNo = int64(k)
					nfts.AuthorId = work.AuthorId
					nfts.OwnerId = work.AuthorId
					nfts.TransferNum = 1
					nfts.ChainStatus = 0
					nfts.UpdateTime = time.Now().Unix()
					Eloquent.Create(&nfts)
				}
			}


			//获取接收者bsn地址
			account, _ := FindAccountByID(v.AccountId)
			//更新拥有者
			var nft TblNfts
			err = Eloquent.Model(TblNfts{}).Where("wid= ? AND transfer_num=1", work.ID).Find(&nft).Error
			if err == nil {
				op := utils.GenerateUri("create")
				Eloquent.Model(TblNfts{}).Where("id=?", v.ID).
					Updates(map[string]interface{}{
						"owner_id":     v.AccountId,
						"operation_id": op,
						"transfer_num": 2,
						"chain_status": 2,
						"price":        0.0,
						"update_time":  time.Now().Unix(),
					})

				var avata utils.Avata
				a, b := avata.CreateNFT(op, account.BsnAddress, work.Name+fmt.Sprintf("#%d", nft.WorkNo))
				fmt.Println("aaaaaaaa:", a, b)
			}
		}

	}
}