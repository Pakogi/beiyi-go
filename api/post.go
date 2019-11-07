package api

import (
	"beiyi/model"
	"beiyi/serializer"

	"github.com/gin-gonic/gin"
)

func Posts(c *gin.Context) {

	posts, err := model.GetPosts()

	if err == nil {
		c.JSON(200, serializer.Response{
			Code: 0,
			Msg:  "OK",
			Data: posts,
		})
	}

}
