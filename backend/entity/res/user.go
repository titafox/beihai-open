package res

type User struct {
	Uri      string `json:"uri" comment:"用户Uri"`
	UserName string `json:"userName" comment:"用户名"`
	Avatar   string `json:"avatar" comment:"头像"`
	Wallet   string `json:"wallet" comment:"钱包地址"`
	//RoyaltyAddress string `json:"royaltyAddress" comment:"版税地址"`
	Signature string               `json:"signature" comment:"个人签名"`
	FollowNum int32                `json:"followNum" comment:"被关注数量"`
	Items     []RecommendWorksInfo `json:"items"`
}

type JwtInfo struct {
	Token string `json:"token" comment:"jwt生成的token"`
}

type IsUpload struct {
	Bol bool `json:"bol" comment:""`
}
