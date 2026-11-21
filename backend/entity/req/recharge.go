package req

type Recharge struct {
	Money uint `json:"money" binding:"required" comment:"充值金额"`
}
