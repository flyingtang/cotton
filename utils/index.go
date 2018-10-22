package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/sirupsen/logrus"
)

type ReturnFunc func(*gin.Context, string, gin.H)

// 返回失败
func ReturnFailed(model string) ReturnFunc {

	return func(c *gin.Context, action string, h gin.H){

		logrus.WithFields(logrus.Fields{"model ": model, "action " : action,}).Error(h["error"])
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message" : h["message"], })
		return
	}
}

// 返回成功
func ReturnSuccess(model string) ReturnFunc {

	return func(c *gin.Context, action string, h gin.H){
		logrus.WithFields(logrus.Fields{"model ": model, "action " : action,}).Debug(h["message"])
		c.JSON(http.StatusOK, h)
		return
	}

}