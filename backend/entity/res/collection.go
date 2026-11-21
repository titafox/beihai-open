package res

type ImportPreCollectionRes struct {
	Title       string `json:"title"`
	Cover       string `json:"cover"`
	Category    uint   `json:"category"`
	RoyaltRate  uint   `json:"royaltRate" comment:"版税比例"`
	Description string `json:"description"`
	Price       string `json:"price"`
}
