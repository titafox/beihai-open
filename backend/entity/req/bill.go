package req

type Bill struct {
	Typ  uint `uri:"typ" comment:"0:收入,1:支出"`
	Page uint `uri:"page" comment:"分页参数"`
}
