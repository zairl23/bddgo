package handlers

import (
	"github.com/gin-gonic/gin"
)


func Login(c *gin.Context) {
	c.JSON(200, Response{
		Code: 1,
		Message: "wrong account",
	})
}