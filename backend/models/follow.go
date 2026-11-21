package models

import (
	"errors"
	"time"
)

type Follow struct {
	ID             uint      `json:"id" gorm:"primary_key"`
	FollowerWallet string    `json:"followerWallet" comment:"关注人的钱包"`
	FollowerUri    string    `json:"follower_uri" comment:"关注人的uri"`
	IdolWallet     string    `json:"idolWallet" comment:"爱豆被关注人的钱包"`
	IdolUri        string    `json:"idol_uri" comment:"爱豆被关注人的uri"`
	IsDel          int       `json:"is_del" comment:"是否删除"`
	CreatedAt      int64     `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (follow *Follow) Insert() error {
	var followDb = Eloquent
	return followDb.Create(follow).Error
}

func (follow *Follow) Update() error {
	var followDb = Eloquent
	return followDb.Model(&follow).Updates(follow).Error
}

func UpdateFollowStatus(id uint, status int) error {
	var followDb = Eloquent
	return followDb.Model(Follow{}).Where("id = ?", id).Updates(map[string]interface{}{
		"is_del": status,
	}).Error
}

func ListIdolByFollow(follow string, idolWallet []string) []Follow {
	var db = Eloquent
	var followList []Follow

	if follow == "" {
		return followList
	}

	db.Table("follow").Model(&Follow{}).Where("follower_uri = ? and idol_wallet in (?) and is_del = ?", follow, idolWallet, 0).
		Select([]string{"idol"}).Find(&followList)
	return followList
}

// GetIsFollow 根据A用户的钱包和B用户的钱包 判断A是否关注B
func GetIsFollow(followerWallet string, idolWallet string) *Follow {
	var db = Eloquent
	var follow = new(Follow)

	if followerWallet == "" {
		return follow
	}
	db.Table("follow").Model(&Follow{}).Where("follower_wallet = ? and idol_wallet = ? and is_del = ?", followerWallet, idolWallet, 0).
		Select([]string{"id"}).Find(&follow)

	return follow
}

// CheckIsExists 根据A用户的uri和B用户的ID 判断数据是否存在
func CheckIsExists(follower string, idolWallet string) (*Follow, error) {
	var db = Eloquent
	var follow = new(Follow)

	if follower == "" || idolWallet == "" {
		return follow, errors.New("param error")
	}
	err := db.Table("follow").Model(&Follow{}).
		Where("follower_wallet = ? and idol_wallet = ?", follower, idolWallet).
		Select([]string{"id", "is_del", "idol"}).Find(&follow).Error

	return follow, err
}
