package res

type PageResult struct {
	Page  uint        `json:"page"`
	IsEnd bool        `json:"isEnd"`
	Total int64       `json:"total"`
	List  interface{} `json:"list"`
}

type UploadToken struct {
	SessionToken string `json:"sessionToken"`
	TmpSecretID  string `json:"tmpSecretID"`
	TmpSecretKey string `json:"tmpSecretKey"`
	StartTime    int64  `json:"startTime"`
	ExpiredTime  int64  `json:"expiredTime"`
}
