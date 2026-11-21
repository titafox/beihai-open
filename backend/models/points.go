package models

type Points struct {
	ID        uint    `json:"id" gorm:"primary_key"`
	BookId    uint    `json:"book_id"`
	Referrer  string  `json:"referrer"`
	Buyer     string  `json:"buyer"`
	Points    float64 `json:"points"`
	CreatedAt int64   `json:"created_at"`
}
