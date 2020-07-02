//如果出現路由內沒有的名稱，就會透過 utils\utils.go 標示狀態為失敗，並且顯示 not found
//透過 errors.go 就能在錯誤時顯示訊息
package application

import (
	u "accountWebAPI/utils"
	"net/http"
)

//傳入http.Handler,最後回傳一個處理過的http.Handler
var NotFoundHandler = func(next http.Handler) http.Handler {
	//當找不到路由名稱時
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)

		//透過 u.Message 加上狀態與訊息，與透過 u.Respond 回傳 JSON 格式的訊息
		u.Respond(w, u.Message(false, "not found"))
		next.ServeHTTP(w, r)
	})
}
