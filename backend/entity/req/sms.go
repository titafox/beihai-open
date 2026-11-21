package req

type Send struct {
	Phone string `json:"phone" binding:"required"`
}
