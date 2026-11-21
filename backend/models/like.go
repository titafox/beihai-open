package models

import (
	"time"
)

type Like struct {
	ID         uint      `json:"id" gorm:"primary_key"`
	BoId       int64     `json:"bo_id"`
	BoUri      string    `json:"bo_uri"`
	BoType     uint      `json:"bo_type"`
	UserWallet string    `json:"user_wallet"`
	UserUri    string    `json:"user_uri"`
	IsDel      int       `json:"is_del"`
	CreatedAt  int64     `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (like *Like) Insert() error {
	var followDb = Eloquent
	return followDb.Create(like).Error
}

func UpdateLikeStatus(id uint, status int) error {
	var followDb = Eloquent
	return followDb.Model(Like{}).Where("id = ?", id).Updates(map[string]interface{}{
		"is_del": status,
	}).Error
}

// CheckIsLike 根据wallet判断是否已点赞
func CheckIsLike(boUri, wallet string, boType int) (*Like, error) {
	var db = Eloquent
	var like = new(Like)
	err := db.Table("like").Model(Like{}).
		Where("bo_uri = ? and user_wallet = ? and bo_type = ?", boUri, wallet, boType).
		Select([]string{"id", "is_del"}).Find(&like).Error

	return like, err
}

func CheckIsLike2(boUri, wallet string, boType int) bool {
	var db = Eloquent
	var like = new(Like)
	err := db.Table("like").Model(Like{}).
		Where("bo_uri = ? and user_wallet = ? and bo_type = ?", boUri, wallet, boType).
		Select([]string{"id", "is_del"}).Find(&like).Error
	bol := true
	if err != nil || like.IsDel == 1 {
		bol = false
	}
	return bol
}

func ListBoByUriBoId(wallet string, boType uint, boIds []int64) []Like {
	var db = Eloquent
	var likeList []Like

	if wallet == "" {
		return likeList
	}

	db.Table("like").Model(&Like{}).Where("user_wallet = ? and bo_type = ? and bo_id in (?) and is_del = ?", wallet, boType, boIds, 0).
		Select([]string{"bo_id"}).Find(&likeList)
	return likeList
}
