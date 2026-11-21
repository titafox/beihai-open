package models

import (
	"coltstail-server/entity/res"
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	ID              uint      `json:"id" gorm:"primary_key"`
	Uri             string    `json:"uri" sql:"index"`
	UserName        string    `json:"userName"`
	UserPassword    string    `json:"userPassword"`
	Wallet          string    `json:"wallet"`
	Avatar          string    `json:"avatar"`
	BackgroundImage string    `json:"background_image" comment:"个人中心背景图"`
	Nonce           string    `json:"nonce" comment:"签名随机字符串"`
	Signature       string    `json:"signature"`
	RoyaltyAddress  string    `json:"royalty_address"`
	FollowNum       uint      `json:"follow_num"`
	IsDel           uint      `json:"isDel"`
	Referrer        string    `json:"referrer"`
	ReferrerNumber  int       `json:"referrer_number"`
	Points          float64   `json:"points"`
	CreatedAt       int64     `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

func (user *User) Insert() error {
	var userDb = Eloquent
	return userDb.Create(user).Error
}

// UpdateByWallet 根据钱包去更新
func (user *User) UpdateByWallet(wallet string) error {
	var userDb = Eloquent
	return userDb.Model(&user).Where("wallet = ?", wallet).Updates(user).Error
}

// UpdateUserById 根据ID 修改用户信息
func UpdateUserById(id uint, data map[string]interface{}) error {
	var userDb = Eloquent
	return userDb.Model(User{}).Where("id = ?", id).Updates(data).Error
}

func GetUserCount() (int, error) {
	var count int
	var userDb = Eloquent
	err := userDb.Model(User{}).Count(&count).Error
	return count, err
}

func GetUserList(offset, limit int) ([]User, error) {
	var users []User
	var userDb = Eloquent
	err := userDb.Where("is_del = 0", true).Offset(offset).Limit(limit).Find(&users).Error
	return users, err
}

func AddReferrerNum(id uint) error {
	//works,err := GetWorks ById(id)
	return Eloquent.Model(&User{}).Where("id = ?", id).UpdateColumn("referrer_number", gorm.Expr("referrer_number + ?", 1)).Error
	//return db.Model(&Works{}).Where("id = ?", id).Updates(map[string]interface{}{"like_num":3}).Error
}

// GetUserByUri 根据uri查询用户信息
func GetUserByUri(uri string) (User, error) {
	var user User
	var userDb = Eloquent

	userDb.Where("uri = ? and is_del = 0", uri).Select([]string{"id", "name", "avatar", "wallet", "signature"}).First(&user, uri)
	err := userDb.Model(&user).Error
	return user, err
}

// GetUserIdByUri 根据uri查询用户ID
func GetUserIdByUri(uri string) (User, error) {
	var user User
	var userDb = Eloquent

	err := userDb.Model(&user).Where("uri = ? and is_del = 0", uri).Select([]string{"id"}).First(&user).Error
	return user, err
}

// GetUserByWallet 通过wallet查询用户
func GetUserByWallet(wallet string) (User, error) {
	var user User
	var userDb = Eloquent

	err := userDb.Model(User{}).Where("wallet = ?  and is_del = 0", wallet).
		Select([]string{"id", "user_name", "uri", "avatar", "wallet", "avatar", "background_image", "nonce", "signature", "royalty_address", "follow_num", "referrer_number", "points", "referrer"}).First(&user).Error
	return user, err
}

func GetUserByID(id int) (User, error) {
	var user User
	var userDb = Eloquent

	userDb.Where("id = ?", id).First(&user, id)
	err := userDb.Model(&user).Error
	return user, err
}

// ListUserByIds 根据用户ID查找用户
func ListUserByIds(ids []uint) []User {
	var userList []User
	var userDb = Eloquent

	userDb.Model(&User{}).Select([]string{"id", "user_name", "uri", "avatar", "wallet"}).Find(&userList, ids)
	return userList
}

// ListUserByWallets 根据用户钱包查找用户
func ListUserByWallets(wallets []string) []User {
	var userList []User
	var userDb = Eloquent

	userDb.Model(&User{}).Select([]string{"id", "user_name", "uri", "avatar", "wallet"}).Where("wallet in (?) ", wallets).Find(&userList)
	return userList
}

func GetUserByUsername(username string) (User, error) {
	var user User
	var userDb = Eloquent
	userDb.Where("user_name = ?", username).First(&user)

	err := Eloquent.Model(&user).Error
	return user, err
}

func GetUserByRecommend() ([]res.RecommendAuthorList, error) {
	var list []res.RecommendAuthorList
	//err := Eloquent.Model(&User{}).Select([]string{"id", "user_name", "wallet", "avatar", "background_image","signature","follow_num"}).Where("index_recommend = 1").Find(&userList).Error
	//return userList, err

	sql := "SELECT id,user_name,wallet,avatar,background_image,signature,follow_num FROM user WHERE index_recommend = 1 ORDER BY sort DESC"
	err := Eloquent.Raw(sql).Scan(&list).Error
	///fmt.Println(err,list)
	return list, err

}
