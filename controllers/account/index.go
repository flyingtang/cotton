package account

import (
	"github.com/gin-gonic/gin"
	"cotton/models"
	"strings"
	"cotton/utils"
	"errors"
	"crypto/sha256"
	"encoding/hex"
	"github.com/dgrijalva/jwt-go"
)


const minLength = 6 // 用户名和密码的最小长度
const hmacSampleSecret = "this is my cotton"

var ReturnFailed, ReturnSucceed utils.ReturnFunc

func init(){
	 ReturnFailed =  utils.ReturnFailed("account")
	 ReturnSucceed =  utils.ReturnSuccess("account")
}

type accountForm struct {
	Username string `form:"username" binding:"required" `
	Password string `form:"password" binding:"required"`
	Nickname string `form:"nickname"`
}

// login api
func Login(c *gin.Context){

	var af accountForm
	if err := c.ShouldBind(&af); err != nil {

		ReturnFailed(c, "c.ShouldBind", gin.H{ "error": err, "message" :"登录失败: 用户名和密码为必传参数",})
		return
	}

	username := strings.TrimSpace(af.Username)
	password := strings.TrimSpace(af.Password)
	if len(username) < minLength {

		ReturnFailed(c, "verify username", gin.H{ "error": errors.New("username too short, "), "message" :"用户名太短,至少6位",})
		return
	}
	if len(password) < minLength {

		ReturnFailed(c, "verify password", gin.H{ "error": errors.New("password too short "), "message" :"密码太短，至少6位",})
		return
	}

	var user models.Account
	db := models.DB.Where("username=?", username).First(&user)
	if db.Error != nil{

		ReturnFailed(c, "models.DB.Where", gin.H{ "error": db.Error, "message" :"用户名无效",})
		return
	}

	h := sha256.New()
	p := h.Sum([]byte(password))
	password = hex.EncodeToString(p)
	if password != user.Password {

		ReturnFailed(c, "models.DB.Where", gin.H{ "error": errors.New("invalid password"), "message" :"密码错误",})
		return
	}
	// jwt-go 办法token
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID,
	})
	tokenString, err := t.SignedString([]byte(hmacSampleSecret))
	if err != nil {
		ReturnFailed(c, "t.SignedString", gin.H{ "error": err, "message" :"登录失败",})
		return
	}

	ReturnSucceed(c, "models.DB.Create", gin.H{"message" :"登陆成功", "token": tokenString})
	return
}


// sign up api
func SignUp(c *gin.Context){

	var af accountForm
	if err := c.ShouldBind(&af); err != nil {

		ReturnFailed(c, "c.ShouldBind", gin.H{ "error": err, "message" :"注册失败: 用户名和密码为必传参数",})
		return
	}

	data := models.Account{
		Username: strings.TrimSpace(af.Username),
		Password: strings.TrimSpace(af.Password),
		Nickname: strings.TrimSpace(af.Nickname),
	}

	if len(data.Username) < minLength {

		ReturnFailed(c, "verify username", gin.H{ "error": errors.New("username too short, "), "message" :"用户名太短,至少6位",})
		return
	}
	if len(data.Password) < minLength {

		ReturnFailed(c, "verify password", gin.H{ "error": errors.New("password too short "), "message" :"密码太短，至少6位",})
		return
	}

	h := sha256.New()
	password := h.Sum([]byte(data.Password))
	pass := hex.EncodeToString(password)
	data.Password =pass

	res  :=  models.DB.Create(&data)
	if res.Error != nil {

		ReturnFailed(c, "models.DB.Create", gin.H{ "error": res.Error, "message" :"注册失败",})
		return
	}

	ReturnSucceed(c, "models.DB.Create", gin.H{"message" :"注册成功",})
	return
}

