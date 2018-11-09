package handlers 

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context) {
	c.String(http.StatusOK, "success")
}