package res

type Nfts struct {
	Wid int64 `json:"wid,omitempty"`
	Id           int64  `json:"id"`
	Name         string `json:"name" comment:"藏品名称"`
	Image        string `json:"image" comment:"藏品图片"`
	AuthorName   string `json:"author_name" commnet:"创造者名称"`
	AuthorAvatar string `json:"author_avatar" comment:"创作者图像"`
	IsSyntheticMaterial uint `json:"is_synthetic_material"`
	Nums         uint   `json:"nums,omitempty"`
}

type SeriesNfts struct {
	Nfts
	WorkNo      uint    `json:"work_no"`
	Price       float64 `json:"price"`
	ChainStatus uint    `json:"chain_status"`
	Status      uint    `json:"status"`
	IsFailCard  bool `json:"is_fail_card"`
}

type SeriesNftsList struct {
	TotalPage int          `json:"total_page"`
	Name      string       `json:"name,omitempty"`
	FailCard      []SeriesNfts `json:"fail_card"`
	List      []SeriesNfts `json:"list"`
}
