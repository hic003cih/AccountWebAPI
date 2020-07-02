//定義回傳的格式，由於這是每一個功能都會用到的功能，因此我們把它獨立抽出來模組化，方便開發使用
package utils

import (
	"encoding/json"
	"net/http"
)

//將要回傳的錯誤訊息打包成Message
func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

//以 json 格式回應,將打包好的Message回傳的map[string]interface{}接收
//然後用json包中的Encode將資料變成json做反應
func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
