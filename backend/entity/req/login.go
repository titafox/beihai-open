package req

type Login struct {
	Phone      string `json:"phone" binding:"required"`
	Code       string `json:"code" form:"code" xml:"code" binding:"required"`
	Invitation string `json:"invitation"`
}
