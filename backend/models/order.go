package models

import (
	"coltstail-server/utils"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type TblOrders struct {
	ID         int64   `json:"id" gorm:"primary_key"`
	No         string  `json:"no"`
	AccountId  int64   `json:"account_id"`
	Wid        int64   `json:"wid" comment:"藏品id"`
	NftId      int64   `json:"nft_id" comment:"藏品id"`
	Price      float64 `json:"price"`
	Num        int     `json:"num"`
	OrderMoney float64 `json:"order_money"`
	PayUrl     string  `json:"pay_url"`
	Typ        uint    `json:"typ" comment:"1:一级市场，2：二级市场"`
	Status     uint    `json:"status"`
	CreateTime int64   `json:"create_time"`
	PayTime    int64   `json:"pay_time"`
	UpdateTime int64   `json:"update_time"`
}

type OrderListRes struct {
	Id         int64   `json:"id"`
	Name       string  `json:"name"`
	Image      string  `json:"image"`
	No         string  `json:"no"`
	Money      float64 `json:"money"`
	PayUrl     string  `json:"pay_url"`
	Countdown  int64   `json:"countdown"`
	Status     uint    `json:"status"`
	StatusText string  `json:"status_text"`
	CreateTime int64   `json:"create_time"`
}

func (order *TblOrders) Insert() error {
	var orderDb = Eloquent
	return orderDb.Create(order).Error
}

func SelectOrder(account_id int64) error {
	//Eloquent.Exec("update tbl_orders set status=3 where unix_timestamp(now()) - create_time > 6000")
	var info TblOrders
	err := Eloquent.Model(&TblOrders{}).Where("account_id = ? and status =1", account_id).Find(&info).Error
	if err == nil {
		return errors.New("你有未完成的支付订单,请先完成支付")
	}

	return nil

}

//huoqu

func GetOrders(uid int64, typ uint) []OrderListRes {
	statusMap := map[uint]string{1: "进行中", 2: "已完成", 3: "已取消"}

	var list []TblOrders
	var err error
	if typ == 0 {
		err = Eloquent.Model(&TblOrders{}).Where("account_id = ?", uid).Order("id desc").Find(&list).Error
	} else {
		err = Eloquent.Model(&TblOrders{}).Where("account_id = ? AND status = ?", uid, typ).Order("id desc").Find(&list).Error
	}

	if err != nil {
		return make([]OrderListRes, 0)
	}
	orderList := make([]OrderListRes, len(list), len(list))
	for k, v := range list {
		orderList[k].Id = v.ID
		//获取藏品信息
		w, _ := GetWoksById(v.Wid)
		var n TblNfts
		err = Eloquent.Model(&TblNfts{}).Where("id=?", v.NftId).First(&n).Error
		if err != nil {
			orderList[k].Name = w.Name
		} else {
			orderList[k].Name = w.Name + fmt.Sprintf(" #%d", n.WorkNo)
		}

		orderList[k].Image = "https://bh-ipfs.oss-cn-hangzhou.aliyuncs.com" + w.Image
		orderList[k].No = fmt.Sprintf("数量x%d", v.Num)
		orderList[k].Money = v.OrderMoney
		if v.Status == 1 {
			orderList[k].Countdown = v.CreateTime + int64(180) - time.Now().Unix()
		}

		orderList[k].PayUrl = v.PayUrl
		orderList[k].Status = v.Status
		orderList[k].StatusText = statusMap[v.Status]
		orderList[k].CreateTime = v.CreateTime
	}
	return orderList
}

func CreateOrder(order_no string, account_id, work_id int64, price float64, num int, order_money float64, pay_url string) error {
	//启用事务处理,
	var transDb = Eloquent
	//增加交易记录
	txErr := transDb.Create(&TblOrders{
		No:         order_no,
		AccountId:  account_id,
		Wid:        work_id,
		Price:      price,
		Num:        num,
		OrderMoney: order_money,
		PayUrl:     pay_url,
		Typ:        1,
		Status:     1,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}).Error

	return txErr
}

// 二级市场下单
func CreateMarketOrder(order_no string, account_id, work_id int64, nft_id int64, price float64, num int, order_money float64, pay_url string) error {
	//启用事务处理,
	var transDb = Eloquent
	err := transDb.Transaction(func(tx *gorm.DB) error {
		//增加交易记录
		if txErr := tx.Create(&TblOrders{
			No:         order_no,
			AccountId:  account_id,
			Wid:        work_id,
			NftId:      nft_id,
			Price:      price,
			Num:        num,
			OrderMoney: order_money,
			PayUrl:     pay_url,
			Typ:        2,
			Status:     1,
			CreateTime: time.Now().Unix(),
			UpdateTime: time.Now().Unix(),
		}).Error; txErr != nil {
			return txErr
		}

		//挂单改为锁单中
		txErr := tx.Model(&TblNfts{}).Where("id=?", nft_id).UpdateColumns(TblNfts{Status: 2, UpdateTime: time.Now().Unix()}).Error
		if txErr != nil {
			return txErr
		}
		return nil
	})
	return err
}

// LockNFTForSale 原子地把挂售中的藏品改为锁单中；不能购买自己的藏品。
// 返回 false 表示已被其他买家锁定、已下架或已售出。
func LockNFTForSale(nftId, buyerId int64) bool {
	res := Eloquent.Model(&TblNfts{}).
		Where("id = ? AND status = 1 AND owner_id <> ?", nftId, buyerId).
		UpdateColumns(map[string]interface{}{"status": 2, "update_time": time.Now().Unix()})
	return res.Error == nil && res.RowsAffected == 1
}

// UnlockNFTForSale 下单失败时释放锁单
func UnlockNFTForSale(nftId int64) {
	Eloquent.Model(&TblNfts{}).Where("id = ? AND status = 2", nftId).
		UpdateColumns(map[string]interface{}{"status": 1, "update_time": time.Now().Unix()})
}

// todo 定时查询快过期的订单 需要更新减少的库存
func GetExpiredOrder() {
	var list []TblOrders
	err := Eloquent.Model(&TblOrders{}).Where("status = 1 and unix_timestamp(now()) - create_time > 180").Find(&list)
	if err != nil && len(list) > 0 {
		for _, v := range list {
			// 只取消仍未支付的订单；若此刻支付回调已把它改成已支付，则跳过，避免释放已售出的藏品/库存
			res := Eloquent.Model(&TblOrders{}).Where("id = ? AND status = 1", v.ID).UpdateColumns(TblOrders{Status: 3, UpdateTime: time.Now().Unix()})
			if res.Error != nil || res.RowsAffected != 1 {
				continue
			}
			//将订单设置为已取消
			if v.NftId > 0 {
				Eloquent.Model(TblNfts{}).Where("id = ? AND status = 2", v.NftId).Updates(map[string]interface{}{
					"status":      1,
					"update_time": time.Now().Unix(),
				})
			} else {
				//redis加库存（原先误用 DecrBy，会再扣一次库存）
				_, e := Redis.IncrBy(fmt.Sprintf("work_%d", v.Wid), int64(v.Num)).Result()
				if e != nil {
					fmt.Println("redis加库存失败,订单：", v.No, "err:", e.Error())
				}
			}

			//Eloquent.Model(&TblWorks{}).Where("id=?", v.Wid).UpdateColumn("stock_sell", gorm.Expr("stock_sell-?", v.Num))
		}
	}
}

// 更新回调订单
func UpdateOrderPayNotify(orderId string) {
	defer func() {
		r := recover()
		fmt.Println("sssssss:", r)
	}()

	//查询订单状态
	var order TblOrders
	err := Eloquent.Model(&TblOrders{}).Where("no = ?", orderId).First(&order).Error
	if err != nil {
		fmt.Println("更新回调订单失败：", err.Error())
		return
	}

	if order.Status != 1 {
		fmt.Println("订单已经处理完成或已取消")
		return
	}

	//启用事务处理,
	var transDb = Eloquent
	res := transDb.Model(TblOrders{}).Where("no=? and status=1", orderId).
		Updates(map[string]interface{}{
			"status":      2,
			"pay_time":    time.Now().Unix(),
			"update_time": time.Now().Unix(),
		})
	if res.Error != nil {
		fmt.Println(fmt.Sprintf("更新订单%s失败：%s", orderId, res.Error.Error()))
		return
	}
	// 幂等：并发/重复回调时只有一个能把 status 从 1 改成 2，其余直接返回，避免重复铸造
	if res.RowsAffected != 1 {
		fmt.Println("订单已被其他回调处理：", orderId)
		return
	}

	//支付成功订单推送到mq
	fmt.Println("orderid=", orderId)
	r := RabbitMQ{Conn: OrderConn}
	if err = r.Publish("mint", "DD:"+orderId); err != nil {
		fmt.Println("rabbitMQ 错误：" + err.Error())
	}
}

// 更新回调订单二级市场
func UpdateOrderBalancePayNotify(orderId string) {
	//查询订单状态
	var order TblOrders
	err := Eloquent.Model(&TblOrders{}).Where("no = ?", orderId).First(&order).Error
	if err != nil {
		fmt.Println("更新回调订单失败：", err.Error())
		return
	}

	if order.Status != 1 {
		fmt.Println("订单已经处理完成或已取消")
		return
	}

	//查询form_address
	var chain struct {
		NftId      string `json:"nft_id"`
		BsnAddress string `json:"bsn_address"`
	}
	Eloquent.Raw("select tn.nft_id,ta.bsn_address from tbl_nfts tn,tbl_account ta where tn.id = ? and tn.owner_id = ta.id", order.NftId).Scan(&chain)
	formAddress := chain.BsnAddress
	nft_id := chain.NftId
	//查询to_address
	Eloquent.Raw("select bsn_address from tbl_account where id = ?", order.AccountId).Scan(&chain)
	toAddress := chain.BsnAddress

	op := utils.GenerateUri("trans")

	//启用事务处理,
	var transDb = Eloquent
	err = transDb.Transaction(func(tx *gorm.DB) error {
		res := tx.Model(TblOrders{}).Where("no=? and status=1", orderId).
			Updates(map[string]interface{}{
				"status":      2,
				"pay_time":    time.Now().Unix(),
				"update_time": time.Now().Unix(),
			})
		if res.Error != nil {
			return res.Error
		}
		// 幂等：重复回调直接回滚
		if res.RowsAffected != 1 {
			return errors.New("订单已处理")
		}

		//增加买入记录
		if txErr := tx.Create(&TblMarketLog{
			NftId:      order.NftId,
			AccountId:  order.AccountId,
			Typ:        2,
			Price:      order.OrderMoney,
			Status:     1,
			CreateTime: time.Now().Unix(),
		}).Error; txErr != nil {
			return txErr
		}

		//txErr := tx.Model(&TblNfts{}).Where("id=?", order.NftId).UpdateColumns(TblNfts{OperationId: op, OwnerId: order.AccountId, Status: 0, UpdateTime: time.Now().Unix()}).Error
		//if txErr != nil {
		//	return txErr
		//}

		// 只有仍处于锁单状态（status=2）的藏品才能过户
		nres := tx.Model(TblNfts{}).Where("id=? AND status = 2", order.NftId).
			Updates(map[string]interface{}{
				"owner_id":     order.AccountId,
				"operation_id": op,
				//"transfer_num": 2,
				"chain_status": 1,
				"status":       0,
				"update_time":  time.Now().Unix(),
			})
		if nres.Error != nil {
			return nres.Error
		}
		if nres.RowsAffected != 1 {
			return errors.New("藏品状态异常，无法过户")
		}
		return nil
	})

	if err != nil {
		fmt.Println(fmt.Sprintf("更新订单%s失败：%s", orderId, err.Error()))
		return
	}

	//开始转移NFT
	var avata utils.Avata
	a, b := avata.NftTransfer(formAddress, nft_id, toAddress, op)
	fmt.Println("操作id:", op, "转移状态：", a, b)
}

// 查询是否有未支付的订单
func GetUnpaidOrder(account_id, id int64) (TblOrders, error) {
	var order TblOrders
	err := Eloquent.Model(&TblOrders{}).Where("status = 1 and nft_id != ? and account_id = ?", id, account_id).First(&order).Error
	return order, err
}
