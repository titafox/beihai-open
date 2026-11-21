package models

import (
	"errors"
	"coltstail-server/conf"
	"coltstail-server/utils"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/streadway/amqp"
	"log"
	"strings"
	"time"
)

var (
	OrderConn *amqp.Connection
)

type RabbitMQ struct {
	Conn *amqp.Connection
}

func QueueInit() {
	var sErr error
	OrderConn, sErr = amqp.Dial(conf.Conf().RABBMIT_MQ_MINT_URL)
	if sErr != nil {
		var err interface{} = "Failed to connect to OrderConn RabbitMQ"
		panic(err)
	}

}

// 异步列队
func (r *RabbitMQ) MintNft() {
	ch, err := r.Conn.Channel()
	if err != nil {
		fmt.Println(err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		"mint", // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			//time.Sleep(10 * time.Second)
			fmt.Println("收到MQ消息：", string(d.Body))

			//获取消息
			msg := strings.Split(string(d.Body), ":")
			if msg[0] == "DD" {
				orderNo := msg[1]
				fmt.Println("收到订单数据：" + orderNo)
				//获取订单信息
				var order TblOrders
				err := Eloquent.Model(&TblOrders{}).Where("no = ?", orderNo).First(&order).Error
				// 幂等：只处理已支付订单，且每个订单只铸造一次（防止消息重投 / 伪造消息重复铸造）
				if err == nil && (order.Status != 2 || !claimOnce("minted:DD:"+orderNo)) {
					fmt.Println("跳过订单（未支付或已处理）：", orderNo)
					err = errors.New("skip")
				}
				if err == nil {
					//获取作品信息
					var works TblWorks
					err = Eloquent.Model(&TblWorks{}).Where("id = ?", order.Wid).First(&works).Error
					if err == nil {
						var count int64
						Eloquent.Model(&TblNfts{}).Where("wid = ?", order.Wid).Count(&count)
						if count < 1 { //说明没有记录，需要添加编号
							nos := utils.GenerateRandomNumber(1, works.Stock+1, works.Stock)
							//定义一个需要插入的TblNfts切片 长度为发行总量的长途
							//nfts := make([]TblNfts,works.Stock,works.Stock)
							for _, v := range nos {
								var nfts TblNfts
								nfts.Wid = works.ID
								nfts.WorkNo = int64(v)
								nfts.AuthorId = works.AuthorId
								nfts.OwnerId = works.AuthorId
								nfts.TransferNum = 1
								nfts.ChainStatus = 0
								nfts.UpdateTime = time.Now().Unix()
								Eloquent.Create(&nfts)
							}

						}

						//获取接收者bsn地址
						account, _ := FindAccountByID(order.AccountId)
						//更新拥有者
						var list []TblNfts
						err = Eloquent.Model(TblNfts{}).Where("wid= ? AND transfer_num=1", works.ID).Limit(order.Num).Find(&list).Error
						if err == nil && len(list) > 0 {
							//更新库存
							//Eloquent.Model(TblWorks{}).Where("id=?", works.ID).
							//	Updates(map[string]interface{}{
							//		"stock_sell":  works.StockSell + order.Num,
							//		"update_time": time.Now().Unix(),
							//	})

							//更新最大购买量
							//var tbls TblSnapshotLog
							//err = Eloquent.Model(&TblSnapshotLog{}).Where("account_id=? and work_id=?", account.ID, works.ID).First(&tbls).Error
							//if err == nil {
							//	Eloquent.Model(TblSnapshotLog{}).Where("id=?", tbls.ID).
							//		Updates(map[string]interface{}{
							//			"use_total":   tbls.UseTotal + order.Num,
							//			"update_time": time.Now().Unix(),
							//		})
							//}

							for _, v := range list {
								op := utils.GenerateUri("create")
								res := Eloquent.Model(TblNfts{}).Where("id=? AND transfer_num=1", v.ID).
									Updates(map[string]interface{}{
										"owner_id":     order.AccountId,
										"operation_id": op,
										"transfer_num": 2,
										"chain_status": 2,
										"price":        order.Price,
										"update_time":  time.Now().Unix(),
									})
								if res.Error != nil || res.RowsAffected != 1 {
									continue // 已被分配，跳过，避免一个编号发给两个人
								}

								//获取作品名称
								w, _ := GetWoksById(v.Wid)
								var avata utils.Avata
								a, b := avata.CreateNFT(op, account.BsnAddress, w.Name+fmt.Sprintf("#%d", v.WorkNo))
								fmt.Println("aaaaaaaa:", a, b)
							}
						}

					}

					//查找当前作品的随机数是否存在

				}

			} else if msg[0] == "KT" { //空投消息
				//获取空投记录
				var airLog TblAirdropsLog
				err = Eloquent.Model(&TblAirdropsLog{}).Where("id = ?", msg[1]).First(&airLog).Error
				if err == nil {
					//查询接收空投者
					var account TblAccount
					err = Eloquent.Model(&TblAccount{}).Where("id = ?", airLog.AccountId).First(&account).Error
					if err == nil {
						//获取空投作品
						var nft TblNfts
						err = Eloquent.Model(TblNfts{}).Where("wid =? and work_no= ? AND transfer_num=1", airLog.PId, airLog.NId).First(&nft).Error
						if err == nil {
							op := utils.GenerateUri("create")
							Eloquent.Model(TblNfts{}).Where("id=?", nft.ID).
								Updates(map[string]interface{}{
									"owner_id":     account.ID,
									"operation_id": op,
									"transfer_num": 2,
									"chain_status": 2,
								})

							var works TblWorks
							Eloquent.Model(&TblWorks{}).Where("id = ?", nft.Wid).First(&works)

							//更新库存
							Eloquent.Model(TblWorks{}).Where("id=?", works.ID).
								Updates(map[string]interface{}{
									"stock_sell":  works.StockSell + 1,
									"update_time": time.Now().Unix(),
								})

							//获取作品名称 todo 需要优化
							w, _ := GetWoksById(airLog.PId)
							var avata utils.Avata
							a, b := avata.CreateNFT(op, account.BsnAddress, w.Name+fmt.Sprintf("#%d", nft.WorkNo))
							fmt.Println("bbbbb:", a, b)

						}

					}

				}
			} else if msg[0] == "HC" { //合成消息
				//获取合成记录
				var tblCompositeLog TblCompositeLog
				err = Eloquent.Model(&TblCompositeLog{}).Where("op = ? and status = 0", msg[1]).Find(&tblCompositeLog).Error
				if err == nil {
					//获取合成活动
					var tblComposite TblComposite
					err = Eloquent.Model(&TblComposite{}).Where("id = ?", tblCompositeLog.CompositeId).First(&tblComposite).Error
					if err == nil {


						//获取作品信息
						var works TblWorks
						Eloquent.Model(&TblWorks{}).Where("id = ?", tblComposite.CompositeId).First(&works)

						//获取当前合成作品是否已生成
						var count int64
						Eloquent.Model(&TblNfts{}).Where("wid = ?", tblComposite.CompositeId).Count(&count)
						if count == 0 { //没有开始生成
							for i := 0; i < works.Stock ; i++ {
								var n TblNfts
								n.Wid = works.ID
								n.WorkNo = int64(i+1)
								n.AuthorId = works.AuthorId
								n.OwnerId = works.AuthorId
								n.TransferNum = 1
								n.Price = 0.0
								n.ChainStatus = 0
								n.Status = 0
								n.UpdateTime = time.Now().Unix()
								Eloquent.Create(&n)
							}
						}

						//获取一条
						var nft TblNfts
						Eloquent.Model(&TblNfts{}).Where("wid = ? and status = 0  and owner_id < 10000", works.ID).First(&nft)

						//开始合成
						op := utils.GenerateUri("hc")
						var txDb = Eloquent
						err = txDb.Transaction(func(tx *gorm.DB) error {
							if txErr := tx.Model(TblNfts{}).Where("id=?", nft.ID).
								Updates(map[string]interface{}{
									"operation_id":op,
									"status":      0,
									"owner_id": tblCompositeLog.AccountId,
									"update_time": time.Now().Unix(),
								}).Error; txErr != nil {
								fmt.Println("bsn:", txErr)
								return txErr
							}

							//更新合成进度
							if txErr := tx.Model(TblCompositeLog{}).Where("id=?", tblCompositeLog.ID).
								Updates(map[string]interface{}{
									"status":      1,
									"update_time": time.Now().Unix(),
								}).Error; txErr != nil {
								fmt.Println("bsn:", txErr)
								return txErr
							}

							//更新库存
							if txErr := Eloquent.Model(TblWorks{}).Where("id=?", works.ID).
								Updates(map[string]interface{}{
									"stock_sell":  works.StockSell + 1,
									"update_time": time.Now().Unix(),
								}).Error; txErr != nil {
								fmt.Println("bsn:", txErr)
								return txErr
							}
							return nil
						})

						if err == nil {
							var account TblAccount
							err = Eloquent.Model(&TblAccount{}).Where("id = ?", tblCompositeLog.AccountId).First(&account).Error
							if err == nil {
								var avata utils.Avata
								a, b := avata.CreateNFT(op, account.BsnAddress, works.Name+fmt.Sprintf("#%d",nft.WorkNo))
								fmt.Println("合成:", a, b)
							}
						}
					}
				}
			}
		}
	}()

	log.Printf(" [*] 等待支付成功的订单消息。 To exit press CTRL+C")
	<-forever
}

// 消息推送
func (r *RabbitMQ) Publish(queue, msg string) error {
	ch, err := r.Conn.Channel()
	defer ch.Close()
	if err != nil {
		return err
	}
	return ch.Publish("", queue, false, false, amqp.Publishing{ContentType: "text/plain", Body: []byte(msg)})
}


// claimOnce 用 Redis SETNX 做一次性标记，返回 true 表示第一次处理
func claimOnce(key string) bool {
	if Redis == nil {
		return true
	}
	ok, err := Redis.SetNX(key, 1, 30*24*time.Hour).Result()
	return err == nil && ok
}
