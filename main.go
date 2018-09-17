package main

import (
	"github.com/gin-gonic/gin"
	"cotton/controllers/account"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"cotton/models"
	"cotton/controllers"
	"github.com/sirupsen/logrus"
)

const version = "/api/v1"
const model = "/admin"


func init(){
	logrus.SetFormatter(&logrus.TextFormatter{TimestampFormat:"2014-03-10 19:57:38.562264131"})
}

func main(){

	db := models.NewMySQL()
	defer db.Close()

	db.AutoMigrate(&models.Account{})
	r := gin.Default()

	// router no auth
	noAuth := r.Group(version)
	{

		noAuth.POST("/signup", account.SignUp)
		noAuth.POST("/login", account.Login)


	}

	// router must auth
	withAuth := noAuth.Group(model)
	withAuth.Use(controllers.AuthRequired)
	//withAuth.GET("/accounts", controllers.)

	r.Run(":4000")

}
