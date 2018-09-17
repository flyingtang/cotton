package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHandle(c *gin.Context){

	c.JSON(http.StatusOK, gin.H{
		"message":"你好",
	})

	return
}


func AuthRequired(c *gin.Context){
	
}