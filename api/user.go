package api

import (
	"beiyi/model"
	"beiyi/serializer"

	"github.com/gin-gonic/gin"
)

func User(c *gin.Context) {
	token := c.GetHeader("X-Beiyitoken")
	user, err := model.GetUser(token)

	if err == nil {
		c.JSON(200, serializer.Response{
			Code: 0,
			Msg:  "OK",
			Data: user,
		})
	}

}
