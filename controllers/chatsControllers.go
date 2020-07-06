package controllers

import (
	"encoding/json"
	"golang-api/models"
	u "golang-api/utils"
	"net/http"
)

//先將圖表的名稱與 url 寫入 db
var CreatePic = func(w http.ResponseWriter, r *http.Request) {
	pic := &models.Pic{}
	err := json.NesDecoder(r.Body).Decode(pic)

	//如果輸入的請求錯誤
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := pic.PicCreare()
	u.Respoond(w, resp)
}

//作讀取，透過傳入圖片名稱即可得到 url

var GetPic = func(w http.ResponseWriter, r *http.Request) {
	chat := &models.Pic{}
	err := json.NesDecoder(r.Body).Decode(chat)
	//如果輸入的請求錯誤
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	resp := models.PicGet(chat.PicName)
	u.Respond(w, resp)
}

//寫入圖表資訊的 api
var CreateChatList = func(w http.ResponseWriter, r *http.Request) {
	ChatList := &models.ChatList{}
	err := json.NewDecoder(r.Body).Decode(ChatList)

	//如果輸入的請求錯誤
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := ChatList.CreateChatList()
	u.Respond(w, resp)
}

//寫入圖表內容的 api
var CreateChatDatta = func(w http.ResponseWriter, r *http.Request) {
	ChatData := &models.ChatData{}
	err := json.NewDecoder(r.Body).Decode(ChatData)

	//如果輸入的請求錯誤
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := ChatData.CreateChatData()
	u.Respond(w, resp)
}

//讀取圖表
var GetChatList = func(w http.ResponseWriter, r *http.Request) {

	ChatList := &models.ChatList{}
	err := json.NewDecoder(r.Body).Decode(ChatList)
	//如果輸入的請求錯誤
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := ChatList.GetChatData()
	u.Respond(w, resp)
}
