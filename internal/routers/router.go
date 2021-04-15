package routers

import (
	"github.com/gin-gonic/gin"
	v1 "go_blog_service/internal/routers/api/v1"
)

/*
@Time    : 2021/4/15 8:28 上午
@Author  : Ziks
@Email   : zhangzhaoxiang7@163.com
@File    : router.go
@Software: GoLand
*/

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	tag := v1.NewTag()
	article := v1.NewArticle()

	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/tags", tag.Create)
		apiv1.DELETE("/tags/:id", tag.Delete)
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.GET("/tags", tag.List)

		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("/articles/:id", article.Delete)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.PATCH("/articles/:id/state", article.Update)
		apiv1.GET("/articles/:id", article.Get)
		apiv1.GET("/articles", article.List)
	}

	return r
}
