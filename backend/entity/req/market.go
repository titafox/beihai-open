package req

type Page struct {
	Id   int64 `form:"id"`
	Page uint  `form:"page"`
}

// 流通
type GetWorksNft struct {
	Page uint   `form:"page"`
	Wid  int64  `form:"wid"`
	Sort string `form:"sort"`
}

// 二级市场
type Market struct {
	Page uint   `form:"page"`
	Typ  uint   `form:"typ"`
	Wid  uint   `form:"wid"`
	Sort string `form:"sort"`
	Keyword string `form:"keyword"`
}

// 二级市场寄售详情
type MarketInfo struct {
	Id uint `uri:"id" binding:"required"`
}

// 寄售
type MarketOnTheNFT struct {
	Id    int64   `json:"id" binding:"required"`
	Price float64 `json:"price" binding:"required,gt=0,lte=1000000"`
}

// 下架
type MarketOffTheNFT struct {
	Id int64 `json:"id" binding:"required"`
}

// 购买下单
type MarketBuyTheNFT struct {
	Id int64 `json:"id" binding:"required"`
}
