package res

type RecommendAd struct {
	Title    string `json:"title" comment:"标题"`
	Cover    string `json:"cover" comment:"封面图"`
	Typ      uint   `json:"typ" comment:"转跳类型,1:链接2:作品"`
	Url      string `json:"url" comment:"转跳链接"`
	WorksUri string `json:"works_uri" comment:"作品标识"`
}
