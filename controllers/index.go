package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"cotton/controllers/account"
)

func IndexHandle(c *gin.Context){

	c.JSON(http.StatusOK, gin.H{
		"message":"你好",
	})

	return
}


func AuthRequired(c *gin.Context){

	 const redirect =  `/#/admin`
	 var tokenString string
	 tokenArray, ok  := c.Request.Header["Authentication"]

	 if !ok || len(tokenArray)  == 0 {

		 c.Redirect(http.StatusMovedPermanently, redirect)
		 c.Abort()
		 return
	 }else {
		 tokenString = tokenArray[0]
	 }

	if len(tokenString) == 0 {
		c.Redirect(http.StatusMovedPermanently, redirect)
		c.Abort()
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(account.HmacSampleSecret), nil
	})

	if err != nil {
		c.Redirect(http.StatusMovedPermanently, redirect)
		c.Abort()
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		c.Set("user", &map[string]interface{}{
			"userId": claims["userId"],
			"username": claims["username"],
			"nickname": claims["nickname"],
		})
		fmt.Println(c.Get("user"))
	} else {
		fmt.Println(err)
		c.Redirect(http.StatusMovedPermanently, redirect)
		c.Abort()
		return
	}

}