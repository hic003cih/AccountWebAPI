//創建帳號的功能，包含帳號格式檢查與是否有重複的信箱
package controllers

import (
	"accountWebAPI/models"
	u "accountWebAPI/utils"
	"encoding/json"
	"net/http"
)

//創建帳號,呼叫account的Create邏輯
var CreateAccount = func(w http.ResponseWriter, r *http.Request) {
	//引用 accounts.go
	account := &models.Account{}
	//解析傳入的 json 請求
	err := json.NewDecoder(r.Body).Decode(account)

	//如果輸入的請求錯誤
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}
	//呼叫 accounts.go 透過傳入的資料創建帳號
	resp := account.Create()
	u.Respond(w, resp)
}

//登入驗證,呼叫account的Login邏輯
var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	//如果輸入的請求錯誤
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	//傳入 Email 與 Password
	resp := models.Login(account.Email, account.Password)
	u.Respond(w, resp)
}

//修改密碼的函數
var EditorAccount = func() {
	account := &models.Account{}                   //引用 accounts.go
	err := json.NewDecoder(r.Body).Decode(account) //解析傳入的 json 請求

	//如果輸入的請求錯誤
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := account.Editor() //呼叫 accounts.go 透過傳入的資料刪除帳號
	u.Respond(w, resp)
}

//刪除帳號的函數
//函數先解讀 json 格式的請求，並且驗證正確性，驗證完成後 call models\accounts.go 做事
var DeleteAccount = func(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}                   //引用 accounts.go
	err := json.NewDecoder(r.Body).Decode(account) //解析傳入的 json 請求

	//如果輸入的請求錯誤
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := account.Delete() //呼叫 accounts.go 透過傳入的資料刪除帳號
	u.Respond(w, resp)
}
