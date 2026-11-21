package req

type SelectOrder struct {
	Typ  uint `uri:"typ" example:"0"`
	Page uint `uri:"page"`
}
