package res

type CastingUser struct {
	Id          uint   `json:"id"`
	Uri         string `json:"uri"`
	Title       string `json:"title"`
	Cover       string `json:"cover"`
	Category    string `json:"category"`
	NftTokenId  string `json:"nftTokenId"`
	Author      string `json:"author" comment:"作者钱包"`
	Description string `json:"description"`
	IsSell      int    `json:"isSell" comment:"是否可以挂售0否1是"`
}

type CastingDetail struct {
	Id          uint     `json:"id"`
	Uri         string   `json:"uri"`
	Title       string   `json:"title"`
	Cover       string   `json:"cover"`
	Category    string   `json:"category"`
	NftTokenId  string   `json:"nftTokenId"`
	Description string   `json:"description"`
	IsSell      int      `json:"isSell" comment:"是否可以挂售0否1是"`
	Author      UserBody `json:"author" comment:"作者信息"`
}

type NftTokenId struct {
	NftTokenId string
	Author     string
}
