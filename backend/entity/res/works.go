package res

import "time"

type WorksList struct {
	ID        int64  `json:"id"`
	Uri       string `json:"uri"`
	Title     string `json:"title"`
	Cover     string `json:"cover"`
	Category  uint   `json:"category"`
	Author    string `json:"author"`
	Owner     string `json:"owner"`
	Price     string `json:"price"`
	NewResell int    `json:"newResell"`
	LikeNum   uint   `json:"likeNum" comment:"点赞数量"`
}

type WorkMap struct {
	ID        uint     `json:"id"`
	Uri       string   `json:"uri"`
	Title     string   `json:"title"`
	Cover     string   `json:"cover"`
	Category  uint     `json:"category"`
	Owner     uint     `json:"owner"`
	Price     string   `json:"price"`
	NewResell int      `json:"newResell"`
	IsLike    bool     `json:"isLike" comment:"是否点赞"`
	LikeNum   uint     `json:"likeNum" comment:"点赞数量"`
	Author    UserBody `json:"author"`
}

type UserBody struct {
	Avatar   string `json:"avatar"`
	UserName string `json:"userName"`
	Uri      string `json:"uri"`
	Wallet   string `json:"wallet"`
	IsFollow bool   `json:"isFollow" comment:"是否关注"`
}

type Works struct {
	Uri            string   `json:"uri"`
	CollectionUri  string   `json:"collection_uri"  comment:"合集的uri"`
	Title          string   `json:"title"`
	SerialNo       string   `json:"serialNo"`
	Cover          string   `json:"cover"`
	Category       uint     `json:"category"`
	Images         string   `json:"images"`
	Author         UserBody `json:"author"`
	Owner          UserBody `json:"owner"`
	IsCollection   bool     `json:"isCollection"`
	CollectionInfo struct {
		ServiesName  string `json:"serviesName"`
		SerialNumber string `json:"serialNumber"`
	} `json:"collectionInfo"`
	TokenInfo struct {
		ContractAddress string `json:"contractAddress"`
		TokenId         string `json:"tokenId"`
		TokenStandard   string `json:"tokenStandard"`
		BlockChain      string `json:"blockChain"`
	}
	Price        string `json:"price"`
	Stock        uint   `json:"stock"`
	TotalStock   uint   `json:"totalStock"`
	GoodsStatus  uint   `json:"goodsStatus"`
	NewResell    int    `json:"newResell"`
	Description  string `json:"description"`
	NftTokenId   string `json:"nftTokenId"`
	NftContract  string `json:"nftContract"`
	IsLike       bool   `json:"isLike" comment:"是否点赞"`
	LikeNum      uint   `json:"likeNum" comment:"点赞数量"`
	IsLock       uint   `json:"isLock" comment:"是否锁定"`
	RoyaltRate   int    `json:"royalt_rate" comment:"版税比例"`
	HtmlText     string `json:"htmlText"`
	ContractType uint   `json:"contract_type"`
}

// 作品列表
type WorksListResponse struct {
	Page  int          `json:"page" comment:"下一页参数的值"`
	IsEnd bool         `json:"isEnd" comment:"是否结束分页"`
	Total int64        `json:"total"`
	List  []WorksList2 `json:"list"`
}

type WorksList2 struct {
	Owner Owner `json:"owner"`
	WorksInfo
}

type Owner struct {
	Avatar   string `json:"avatar" comment:"图像"`
	IsFollow bool   `json:"isFollow" comment:"是否已关注"`
	UserName string `json:"userName" comment:"用户名"`
	Wallet   string `json:"wallet" comment:"钱包地址"`
}

type CollectionImages struct {
	Cover string `json:"cover"`
}

type WorksInfo struct {
	Id               int                `json:"id"`
	Uri              string             `json:"uri"`
	Category         uint               `json:"category"`
	Title            string             `json:"title"`
	SerialNo         string             `json:"serial_no"`
	Cover            string             `json:"cover"`
	Description      string             `json:"description"`
	NftTokenId       string             `json:"nftTokenId"`
	Author           string             `json:"author"`
	Owner            string             `json:"owner"`
	Price            string             `json:"price" comment:"价格Wei"`
	LikeNum          int                `json:"like_num" comment:"喜欢人数"`
	IsLike           bool               `json:"isLike" comment:"是否点赞"`
	IsOwner          bool               `json:"isOwner" comment:"是否为拥有者"`
	ContractType     uint               `json:"contract_type"`
	CollectionUri    string             `json:"collectionUri"  comment:"合集的uri"`
	CollectionImages []CollectionImages `json:"collectionImages" comment:"合集作品前三张"`
	//拥有者信息
	Avatar    string    `json:"avatar" comment:"图像"`
	IsFollow  bool      `json:"isFollow" comment:"是否已关注"`
	UserName  string    `json:"userName" comment:"用户名"`
	UpdatedAt time.Time `json:"updated_at"`
}
