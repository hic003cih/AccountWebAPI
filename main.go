package main

import (
	"fmt"
	"net/http"
	"os"

	/* "golang-api/application"
	"golang-api/controllers" */

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
