package tests

import (
	"coltstail-server/utils"
	"encoding/json"
	"fmt"
	"testing"
)

func TestCreateWallet(t *testing.T) {
	var p utils.CreateWalletParam
	p.OrderId = "CC123"
	p.UserId = "346"
	p.UserName = "张三"
	p.IdCard = "110101199001011234"
	a, b := p.CreateWallet()
	t.Log(a, b)
}

func TestSelectWallet(t *testing.T) {
	var p utils.SelectWalletParam
	p.UserId = "346"
	p.OrderId = "CC1241"
	a, b := p.SelectWallet()
	t.Log(a, b)
}

func TestSelectWalletStatus(t *testing.T) {
	var p utils.SelectWalletParam
	p.UserId = "346"
	a, b, c := p.SelectWalletStatus()
	t.Log(a, b, c)
}

type DivDetails struct {
	RiskDivType string `json:"riskDivType"`
	DivAmt      string `json:"divAmt"`
	DivAcctId   string `json:"divAcctId"`
	DivCustId   string `json:"div_cust_id"`
}

func TestBalancePay(t *testing.T) {
	divDetails := []DivDetails{
		{RiskDivType: "1", DivAmt: "96.00", DivAcctId: "B00000001", DivCustId: "6666000000000001"},
		{RiskDivType: "1", DivAmt: "4.00", DivAcctId: "B00000002", DivCustId: "6666000000000002"},
	}

	divDetailsJson, _ := json.Marshal(divDetails)
	fmt.Println("divDetailsJson=", string(divDetailsJson))

	var p utils.BalancePayParam
	//p.UserId = "10775"
	p.OrderId = "t23456"
	p.TransAmt = "100.00"
	p.GoodsDesc = "test02"
	p.UserCustId = "6666000000000003"
	p.DivDetails = string(divDetailsJson)
	a, b := p.BalancePay()
	t.Log(a, b)
}
