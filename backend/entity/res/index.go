package res

//type RecommendWorksListResponse struct {
//	RecommendAuthorList
//	Items []RecommendWorksInfo
//}

type RecommendAuthorList struct {
	ID              uint                 `json:"id" gorm:"primary_key"`
	UserName        string               `json:"user_name"`
	Wallet          string               `json:"wallet"`
	Avatar          string               `json:"avatar"`
	BackgroundImage string               `json:"background_image" comment:"个人中心背景图"`
	Signature       string               `json:"signature"`
	FollowNum       uint                 `json:"follow_num"`
	WorksCount      int64                `json:"works_count"`
	Items           []RecommendWorksInfo `json:"items"`
}

type RecommendWorksInfo struct {
	Id           int    `json:"id"`
	Uri          string `json:"uri"`
	Category     uint   `json:"category"`
	Title        string `json:"title"`
	SerialNo     string `json:"serial_no"`
	Cover        string `json:"cover"`
	Description  string `json:"description"`
	NftTokenId   string `json:"nft_token_id"`
	Price        string `json:"price" comment:"价格Wei"`
	LikeNum      int    `json:"like_num" comment:"喜欢人数"`
	IsLike       bool   `json:"isLike" comment:"是否点赞"`
	ContractType uint   `json:"contract_type"`
	Owner        string `json:"owner"`
	UserInfo     Owner  `json:"user_info"`
}
