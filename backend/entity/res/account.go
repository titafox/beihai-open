package res

type InviteList struct {
	Nums      uint     `json:"nums"`
	TotalPage int      `json:"total_page"`
	List      []Invite `json:"list"`
}

type Invite struct {
	Nickname   string `json:"nickname"`
	CreateTime int64  `json:"create_time"`
}

type InviteTop struct {
	Nickname string `json:"nickname"`
	Nums     int    `json:"nums"`
}
