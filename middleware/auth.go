package middleware

import (
	"beiyi/model"
	"beiyi/serializer"

	"github.com/gin-gonic/gin"
)

// AuthRequired 檢查一定要把 token 帶在 header
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("X-Beiyitoken")

		if _, err := model.GetUser(token); err == nil {
			c.Next()
			return
		}

		c.JSON(200, serializer.CheckLogin())
		c.Abort()
	}
}
