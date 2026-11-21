package res

type Category struct {
	ID           uint   `json:"id" gorm:"primary_key"`
	CategoryName string `json:"categoryName"`
	CategoryEn   string `json:"categoryEn"`
	CategoryVi   string `json:"categoryVi"`
	ItemCount    int64  `json:"itemCount"`
}
