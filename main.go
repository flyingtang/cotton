package main

import (
	"cotton/controllers"
	"cotton/controllers/account"
	"cotton/models"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
)

const version = "/api/v1"
const model = "/admin"

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{TimestampFormat: "2014-03-10 19:57:38.562264131"})
}

func main() {

	db := models.NewMySQL()
	defer db.Close()

	//db.AutoMigrate(&models.Account{})

	r := gin.Default()

	// router no auth
	noAuth := r.Group(version)
	{

		noAuth.POST("/signup", account.SignUp)
		noAuth.POST("/login", account.Login)

	}

	// router must auth
	withAuth := noAuth.Group(model)
	//withAuth.Use(controllers.AuthRequired)
	//withAuth.Use(controllers.AuthRequired)
	{
		withAuth.GET("/index", controllers.IndexHandle)
		withAuth.GET("/accounts", account.GetAccounts)
	}

	r.Run(":4000")

}
