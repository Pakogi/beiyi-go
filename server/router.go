package server

import (
	"beiyi/api"
	"beiyi/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	// 路由
	v2 := r.Group("/api/v2")
	{
		v2.GET("ping", api.Ping)

		auth := v2.Group("")
		auth.Use(middleware.AuthRequired())
		{
			auth.GET("posts", api.Posts)
		}
	}
	return r
}
