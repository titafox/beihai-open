package res

type ScParameters struct {
	ContractAddress string `json:"contract_address" comment:"合约地址"`
	SpenderAddress  string `json:"spender_address"`
	ABI             string `json:"abi" comment:"ABI 参数"`
}
