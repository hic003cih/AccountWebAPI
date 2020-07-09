//主要的邏輯放置在這
package models

import (
	u "accountWebAPI/utils" // 開頭的 u 代表用 u 來引用他
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

//JWT 結構
type Token struct {
	UserId uint
	jwt.StandardClaims
}
//帳號結構
type Account struct {
	gorm.Model
	Email    string `json:email`
	Password string `json:"password"`
	Token    string `json:"token";sql:"-"`
}

//驗證傳入訊息
func (account *Account) Validate() (map[string]interface{}, bool) {

	if !strings.Contains(account.Email, "@") {
		return u.Message(false, "Email format error"), false
	}
	//如果密碼小於8位數，回傳帳號格式錯誤
	if len(account.Password) < 8 {
		return u.Message(false, "Password format error"), false
	}
	//認證 Email 是否存在
	temp := &Account{}

	//去db中比對email,如果存在回傳err
	err := GetDB().Table("accounts").Where("email=?", account.Email).First(temp).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error"), false
	}
	//如果 Email 已存在，回傳已被使用
	if temp.Email != "" {
		return u.Message(false, "Email has been used"), false
	}
	return u.Message(false, "Requirement passed"), true
}

//帳號創立
func (account *Account) Create() map[string]interface{} {
	//先做Validate比對
	if resp, ok := account.Validate(); !ok {
		return resp
	}
	//生成hash密碼
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)
	GetDB().Create(account)

	//如果取到的值小於
	if account.ID <= 0 {
		return u.Message(false, "Failed to create account, connection error.")
	}

	////準備聲明內容
	tk := &Token{UserId: account.ID}
	//宣告使用 HS256 與加入Payload 的聲明內容
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	//將 token_pwd 設定為 secret 並產生 jwt
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	//將產生好的 jwt 值給予 account.Token
	account.Token = tokenString

	account.Password = ""

	//創建完成
	response := u.Message(true, "Account has been created")
	response["account"] = account
	return response
}
//登入
func Login(email, password string) map[string]interface{} {
	account := &Account{}
	//帳號驗證
	err := GetDB().Table("accounts").Where("email = ?", email).First(account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Email address not found")
		}
		return u.Message(false, "Connection error")
	}
	//密碼驗證
	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return u.Message(false, "Invalid login credentials")
	}
	account.Password = ""
	//創建 token
	tk :=&Token(UserId:account.ID)
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString

	resp:=u.Message(true, "Logged In")
	resp["account"] = account
	return resp
}
func parseToken(token string) (*jwt.StandardClaims, error) { 
    jwtToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (i interface{}, e error) {
         return []byte(config.Secret), nil 
         }) 
    if err == nil && jwtToken != nil { 
     if claim, ok := jwtToken.Claims.(*jwt.StandardClaims); ok && jwtToken.Valid { return claim, nil } 
    } 
    return nil, err 
}

//密碼修改
func Editor(email, oldPassword, newPassword string) (map[string]interface{}) {

	//如果密碼小於8位數，回傳帳號格式錯誤
	  if len(oldPassword) < 8 && len(newPassword) < 8 {
		  return u.Message(false, "Password format error"), false
	  }
  
	//帳號驗證
	  err := GetDB().Table("accounts").Where("email = ?", email).First(account).Error
	  if err != nil {
		  if err == gorm.ErrRecordNotFound {
			  return u.Message(false, "Email address not found")
		  }
		  return u.Message(false, "Connection error")
	  }
	
	//舊密碼驗證
	  err = bcrypt.CompareHashAndPassword([]byte(oldPassword), []byte(password))
	  if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { 
		  return u.Message(false, "pwd error")
	  }  
	
	  hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	  account.Password = string(hashedPassword)  	
	
	//保存
	GetDB().Save(account)
  
	//清空 pwd 的數值以便回傳
	account.Password = ""
	
	//修改完成
	  response := u.Message(true, "Account has been editor")
	  response["account"] = account
	  return response
  }
  //帳號刪除
  //包含驗證輸入的密碼是否正確
func (account *Account) Editor() (map[string]interface{}) {  

	account := &Account{}
	
	err := GetDB().Table("accounts").Where("email = ?", email).First(account).Error
	  if err != nil {
		  if err == gorm.ErrRecordNotFound {
			  return u.Message(false, "Email address not found")
		  }
		  return u.Message(false, "Connection error. Please retry")
	  }
  
	  err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	  if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		  return u.Message(false, "pwd error")
	  }
	
	//刪除
	GetDB().Delete(Account{}, "email = ?", email)
  
	//清空 pwd 的數值以便回傳
	account.Password = ""
	
	//刪除完成
	  response := u.Message(true, "Account has been delete")
	  response["account"] = account
	  return response
  }