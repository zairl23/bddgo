package handlers

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int `json:"code"`
	Message string `json:"message"`
}

func Json(c *gin.Context) {
	c.JSON(200, Response{
		Code: 0,
		Message: "success",
	})
}