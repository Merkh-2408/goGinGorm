package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 再做一次分量，将具体方法分出来
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong3",
	})
}
