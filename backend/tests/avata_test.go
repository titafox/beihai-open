package tests

import (
	"coltstail-server/utils"
	"encoding/json"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	var avata utils.Avata
	bsnAddress, err3 := avata.CreateAccount(utils.GenerateUri("n"), utils.GenerateUri("op"))
	t.Log(bsnAddress, err3)
}

func TestCreateNFT(t *testing.T) {
	var avata utils.Avata
	bol, err := avata.CreateNFT("sq3", "iaa1sal7lvy8l7lycf4s4dp4e87fu4rt2nldpg2qc6", "s44")
	t.Log(bol, err)
}

func TestTransferNFT(t *testing.T) {
	var avata utils.Avata
	bol, err := avata.NftTransfer("iaa1xfhfweu9k0txe3wenz7ynadt5tf2xvlt7sce4m", "avata2tudxguexipnzx6mo60lueoeoan", "iaa1nhkj2n2efd5r95ap56ldp5lp6pc06h3askergm", "jfdle")
	t.Log(bol, err)
}

type Tx struct {
	Message     string `json:"message"`
	BlockHeight int    `json:"block_height"`
	Nft         string `json:"nft"`

	Status int `json:"status"`
	//Timestamp time.Time `json:"timestamp"`
	TxHash string `json:"tx_hash"`
	Type   string `json:"type"`
}

func TestTx(t *testing.T) {
	str := []byte(`{"block_height":2283675,"message":"SUCCESS","module":"nft","mt":"","operation_id":"operation1579779358925692928","status":1,"timestamp":"2022-10-11 10:22:15 +0000 UTC","tx_hash":"61F428E59C73E6B4C3C57E9B12C5D9399D9D4398C8A6908003D62688FF185285","type":"mint_nft"`)
	var tx Tx
	err := json.Unmarshal(str, &tx)
	t.Log(err, tx)
}
