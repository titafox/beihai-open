package res

type Transaction struct {
	CastingUri   string `json:"uri" comment:"铸件uri"`
	Transaction  uint   `json:"transaction" comment:"交易类型1铸造2销售"`
	OrderNo      string `json:"orderNo" comment:"关联订单号"`
	Price        string `json:"price" comment:"价格ETH扩大100000000"`
	Avatar       string `json:"avatar" comment:"用户图像"`
	BuyerWallet  string `json:"buyerWallet" comment:"买家钱包"`
	SellerWallet string `json:"sellerWallet" comment:"卖家钱包"`
	CreatedAt    uint   `json:"createdTime" comment:"创建时间"`
}

type HistoryPrice struct {
	CastingUri string `json:"uri" comment:"铸件uri"`
	Price      string `json:"price" comment:"价格ETH扩大100000000"`
	CreatedAt  uint   `json:"createdTime" comment:"创建时间"`
}

type TransactionData struct {
	Bill  []Transaction  `json:"bill"`
	Price []HistoryPrice `json:"price"`
}

// 我拥有的nft
type MyTokensResponse struct {
	Page  int        `json:"page" comment:"下一页参数的值"`
	IsEnd bool       `json:"isEnd" comment:"是否还有下一页"`
	Total int64      `json:"total"`
	List  []MyTokens `json:"list"`
}

type MyTokens struct {
	Owner Owner `json:"owner"`
	TokenInfo
}

type TokenInfo struct {
	Id          int    `json:"id"`
	Uri         string `json:"uri"`
	Category    uint   `json:"category"`
	Cover       string `json:"cover"`
	Title       string `json:"title"`
	Description string `json:"description"`
	SerialNo    string `json:"serial_no"`
	Author      string `json:"author"`
	Owner       string `json:"owner"`
	//IsSell      uint   `json:"isSell"`
	IsLock           uint               `json:"isLock" comment:"是否锁定"`
	IsOwner          bool               `json:"isOwner"`
	NftTokenId       string             `json:"nftTokenId"`
	CollectionUri    string             `json:"collectionUri"  comment:"合集的uri"`
	CollectionImages []CollectionImages `json:"collectionImages"  comment:"合集图片top3"`
}

// 铸件合集作品
type CreateCastingSubResp struct {
	CollectionUri string `json:"collectionUri" comment:"合集的uri"`
}
