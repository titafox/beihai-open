package req

// 账户实名认证
type Authentication struct {
	Name   string `json:"name"`
	IdCard string `json:"id_card"`
}

type UpdateAccountInfo struct {
	Nickname string `json:"nickname" validate:"required(m=昵称不能为空)"`
	Avatar   string `json:"avatar" validate:"required(m=头像不能为空);url(m=请上传正确的头像)"`
	Desc     string `json:"desc" validate:"required(m=个人简介不能为空)"`
}
