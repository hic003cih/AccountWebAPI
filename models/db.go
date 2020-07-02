package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"                     //使用 gorm 作為 orm 工具
	_ "github.com/jinzhu/gorm/dialects/postgres" //使用 gorm 的 postgresql 連結功能
	"github.com/joho/godotenv"                   //使用 godotenv 取得 .env 內的資料
)

var db *gorm.DB

//初始化時直接執行
func init() {

	//執行godotenv套件
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	//透過 Getenv 來讀取 .env
	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	//連結 db,用Sprintf將字都連接起來
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	fmt.Println(dbUri)

	//錯誤攔截與使用gorm套件建立連接
	conn, err := gorm.Open("postgres", dbUri)

	if err != nil {
		fmt.Print(err)
	}
	//將連結存到var db變數
	db = conn

	//db.Debug().AutoMigrate(&Account{}, &Contact{})
	db.Debug().AutoMigrate(&Account{})
}

//建立一個GetDB()讓其他package可以直接使用,取得DB
func GetDB() *gorm.DB {
	return db
}
