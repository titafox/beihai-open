package res

type BuyParameters struct {
	From                  string `json:"from" comment:"nft拥有者"`
	To                    string `json:"to" comment:"买家"`
	NftTokenId            string `json:"nft_token_id" comment:"nft唯一标识"`
	PercentageBasisPoints int    `json:"percentage_basis_points" comment:"版权基础百分比点数"`
	Value                 string `json:"value" comment:"nft价格"`
	MetaData              string `json:"metaData" comment:"元数据"`
}

type SellParameters struct {
	Uri string `json:"uri" comment:"uri"`
}

type GetNftPreSaleRes struct {
	UserName string `json:"user_name"`
	Avatar   string `json:"avatar"`
	Wallet   string `json:"wallet"`
	Uri      string `json:"uri"`
	Title    string `json:"title"`
	Cover    string `json:"cover"`
	Price    string `json:"price"`
	EndTime  int64  `json:"end_time"`
}
