package handlers

import (
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	c.JSON(200, Response{
		Code: 0,
		Message: "success registered",
	})
}