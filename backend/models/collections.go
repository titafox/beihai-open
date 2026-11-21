package models

import (
	"time"
)

type Collections struct {
	ID             int64     `json:"id" gorm:"primary_key"`
	Uri            string    `json:"uri"`
	Title          string    `json:"title"`
	Cover          string    `json:"cover"`
	Category       uint      `json:"category"`
	Author         string    `json:"author"`
	RoyaltRate     int       `json:"royalt_rate" comment:"版税比例"`
	Description    string    `json:"description"`
	Stock          int       `json:"stock"`
	RemainingStock int       `json:"remaining_stock"`
	Code           string    `json:"code"`
	Status         uint      `json:"status" comment:"0:隐藏,1:显示"`
	CreatedAt      int64     `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (collection *Collections) Get(condition string, uri string) (*Collections, error) {
	err := Eloquent.Model(Collections{}).Select(condition).Where("uri = ?", uri).First(&collection)
	return collection, err.Error
}

// 获取合集信息
func GetCollectionInfo(condition string, uri string) (*Collections, error) {
	var collection Collections
	err := Eloquent.Model(Collections{}).Select(condition).Where("uri = ?", uri).First(&collection)
	return &collection, err.Error
}

// 获取合集信息
func GetCollectionByCode(condition string, code string) (*Collections, error) {
	var collection Collections
	err := Eloquent.Model(Collections{}).Select(condition).Where("code = ?", code).First(&collection)

	return &collection, err.Error
}

func (collection *Collections) Create() error {
	var userDb = Eloquent
	return userDb.Create(collection).Error
}

func (collection *Collections) Update() error {
	return Eloquent.Model(Collections{}).Where("id = ?", collection.ID).Updates(collection).Error
}
