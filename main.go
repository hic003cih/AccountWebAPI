package main

import (
	"fmt"
	"net/http"
	"os"

	"accountWebAPI/controllers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	//api路由
	/* router.HandleFunc("/api/user/create", controllers.CreateAccount).Methods("POST")
	 router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/chat/listGet", controllers.GetContactsFor).Methods("POST")
	router.HandleFunc("/api/chat/get", controllers.CreateContact).Methods("POST")
	router.HandleFunc("/api/chat/create", controllers.GetContactsFor).Methods("POST")  */
	//router.Use(application.JwtAuthentication)

	//圖表 api
	//router.HandleFunc("/api/chat/createList", controllers.CreateChatInfo).Methods("POST")
	//router.HandleFunc("/api/chat/createData", controllers.CreateChatData).Methods("POST")

	//讀取圖表 api
	//router.HandleFunc("/api/chat/getList", controllers.GetChatList).Methods("POST")
	//router.HandleFunc("/api/chat/getData", controllers.GetChatData).Methods("POST")

	//測試用API
	router.HandleFunc("/api/test/test", controllers.Test).Methods("POST")

	//修改密碼 api
	router.HandleFunc("/api/user/editor", controllers.EditorAccount).Methods("POST")
	//刪除帳號 api
	router.HandleFunc("/api/user/delete", controllers.DeleteAccount).Methods("POST")
	//設定PORT號
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Println(port)
	//
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}
