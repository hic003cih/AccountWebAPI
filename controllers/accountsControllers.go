//創建帳號的功能，包含帳號格式檢查與是否有重複的信箱
package controllers

import (
	"accountWebAPI/models"
	u "accountWebAPI/utils"
	"encoding/json"
	"net/http"
)

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