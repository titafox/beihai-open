package models

type WorksExtend struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	WorksId     int64  `json:"works_id"`
	Description string `json:"description"`
}

func (w *WorksExtend) Insert() error {
	var extendDb = Eloquent
	return extendDb.Create(w).Error
}

func (w *WorksExtend) Update(worksId int64) error {
	var extendDb = Eloquent
	return extendDb.Model(&w).Where("works_id=?", worksId).Updates(w).Error
}

func GetDescById(id int64) (WorksExtend, error) {
	var works WorksExtend
	var extendDb = Eloquent

	//err := extendDb.Model(&works).Where("works_id = ?", id).First(&works).Error
	extendDb.Where("works_id = ?", id).First(&works)
	err := extendDb.Model(&works).Error
	return works, err
}
