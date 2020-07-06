package models

import (
	u "golang-api/utils"

	"github.com/jinzhu/gorm"
)

//圖片 url 儲存
type Pic struct {
	gorm.Model
	PicName string `json:"PicName"`
	PicUrl  string `json:"PicUrl"`
}

//圖表清單
type ChatList struct {
	gorm.Model
	ChatName      string `json:"ChatName"`
	ChatDepiction string `json:"ChatDepiction"`
	XName         string `json:"XName"`
	YName         string `json:"YName"`
}

//圖表資料
type ChatData struct {
	gorm.Model
	ChatName string `json:"ChatName"`
	XChat    string `json:"XChat"`
	YChat    string `json:"YChat"`
}

//實作增加圖片 url 的功能
func (pic *Pic) PicCreate() map[string]interface{} {
	GetDB().Create(pic)

	response := u.Message(true, "Pic has been created")
	return response
}

//實作讀取的功能
func PicGet(PicName string) map[string]interface{} {
	pic := &Pic{}
	//去取DB內的資料
	err := GetDB().Table("pic").Where("PicName = ?", PicName).First(PicName).Error

	if err != nil {
		//沒有找到資料回傳找不到
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "pic not found")
		}
		//連接失敗回傳
		return u.Message(false, "Connection error")
	}

	resp := u.Message(true, "pic get")
	resp["pic"] = pic
	return resp
}

//實作圖表邏輯的功能
func (chatlist *ChatList) CreateChatList() map[string]interface{} {
	GetDB().Create(chatlist)

	resp := u.Message(true, "success")
	resp["chatlist"] = chatlist
	return resp
}

//圖表內容寫入
func (chatdata *ChatData) CreateChatData() map[string]interface{} {
	GetDB().Create(chatdata)

	resp := u.Message(true, "success")
	resp["chatdata"] = chatdata
	return resp
}

//驗證傳入訊息
func (chatlist *ChatList) Validate() (map[string]interface{}, bool) {
	//認證圖表是否存在
	temp := &ChatList{}

	err := GetDB().Table("chatlist").Where("chatname = ?", chatlist.ChatName).First(temp).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error"), false
	}
	if temp.ChatName != "" {
		return u.Message(false, "chat has been used"), false
	}
	return u.Message(false, "Requirement passed"), true
}
