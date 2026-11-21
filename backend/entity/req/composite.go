package req

type StartCompositeData struct {
	CompositeId int64            `json:"composite_id"`
	Data        []StartComposite `json:"data"`
}

type StartComposite struct {
	Wid int64   `json:"wid"`
	Ids []int64 `json:"ids"`
}

type CompositeRes struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}
