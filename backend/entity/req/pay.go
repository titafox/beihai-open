package req

type DivDetails struct {
	RiskDivType string `json:"riskDivType"`
	DivAmt      string `json:"divAmt"`
	DivAcctId   string `json:"divAcctId"`
	DivCustId   string `json:"div_cust_id"`
}
