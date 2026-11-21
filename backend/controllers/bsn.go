package controllers

import (
	"coltstail-server/models"
	"coltstail-server/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
)

type Tx struct {
	Message     string `json:"message"`
	BlockHeight int    `json:"block_height"`
	Nft         string `json:"nft"`
	Status      int    `json:"status"`
	//Timestamp time.Time `json:"timestamp"`
	TxHash      string `json:"tx_hash"`
	Type        string `json:"type"`
	OperationId string `json:"operation_id"`
}

type NftInfo struct {
	NftId string `json:"nft_id"`
}

func BsnNotify(ctx *gin.Context) {
	data, err := io.ReadAll(io.LimitReader(ctx.Request.Body, 1<<20))
	fmt.Println("收到BSN回调：", err, string(data))
	_, _ = ctx.Writer.Write([]byte("SUCCESS"))

	var tx Tx
	var nftInfo NftInfo
	if err = json.Unmarshal(data, &tx); err != nil || tx.OperationId == "" {
		fmt.Println("解析Bsn回调失败1：", err)
		return
	}

	// 回调本身不可信（任何人都能 POST），用 operation_id 向 AVATA 回查真实结果
	real, err := utils.QueryTx(tx.OperationId)
	if err != nil {
		fmt.Println("回查AVATA交易失败：", tx.OperationId, err)
		return
	}
	if real.Status != 1 {
		fmt.Println("AVATA交易未成功：", tx.OperationId, real.Status, real.Message)
		return
	}

	if err = json.Unmarshal([]byte(real.Nft), &nftInfo); err != nil {
		fmt.Println("解析Bsn回调失败2：", err)
		return
	}

	models.TxHook(tx.OperationId, real.BlockHeight, real.TxHash, real.Type, nftInfo.NftId)
}
