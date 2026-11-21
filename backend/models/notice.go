package models

import "fmt"

type TblNoticeType struct {
	ID      int64  `json:"id" gorm:"primary_key"`
	Name    string `json:"name"`
	Display bool   `json:"display"`
}

type TblNotice struct {
	ID         int64  `json:"id" gorm:"primary_key"`
	Nid        int64  `json:"nid"`
	Title      string `json:"title"`
	Summary    string `json:"summary"`
	Content    string `json:"content"`
	Display    bool   `json:"display"`
	CreateTime int64  `json:"create_time"`
}

// 公告类型
func GetNoticeType() []TblNoticeType {
	var typeList []TblNoticeType
	err := Eloquent.Model(&TblNoticeType{}).Find(&typeList).Error
	if err != nil {
		return make([]TblNoticeType, 0)
	}
	return typeList
}

// 获取单个公告详情
func GetNoticeDetail(id int64) *TblNotice {
	var notice TblNotice
	err := Eloquent.Model(&TblNotice{}).Where("id = ?", id).First(&notice).Error
	if err != nil {
		return nil
	}
	return &notice
}

// 小喇叭
func GetNoticeLaba() []TblNotice {
	var notice []TblNotice
	err := Eloquent.Model(&TblNotice{}).Where("display=1").Find(&notice).Error
	if err != nil {
		return make([]TblNotice, 0)
	}
	return notice
}

// 获取某个类型下的所有公告
func GetNoticeById(id int64) []TblNotice {
	var noticeList []TblNotice
	err := Eloquent.Model(&TblNotice{}).Where("nid = ?", id).Find(&noticeList).Error
	if err != nil {
		return make([]TblNotice, 0)
	}
	return noticeList
}

type NoticeListResp struct {
	Count  int         `json:"count"`
	Result []TblNotice `json:"result"`
}

// 获取某个类型下的所有公告2
func GetNoticeList(id int64) *NoticeListResp {
	fmt.Println("iiiiiii=", id)
	var resp NoticeListResp
	list := GetNoticeById(id)
	resp.Count = len(list)
	resp.Result = list
	return &resp
}

// 综合公告 每个三条
func GetNoticeTypeAll() map[string][]TblNotice {
	all := make(map[string][]TblNotice)
	typeList := GetNoticeType()

	if len(typeList) == 0 {
		return nil
	}

	for _, v := range typeList {
		all[v.Name] = GetNoticeById(v.ID)
	}

	return all
}
