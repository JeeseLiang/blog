package routers

import (
	"blog/internal/middleware"
	v1 "blog/internal/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Translation())
	apiv1 := r.Group("/api/v1")
	tag := v1.NewTag()
	articles := v1.NewArticle()
	{
		apiv1.POST("/tags", tag.Get)
		apiv1.DELETE("/tags/:id", tag.Delete)
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.GET("/tags", tag.List)
		apiv1.PATCH("/tags/:id/state", tag.Update)

		apiv1.POST("/articles", articles.Create)
		apiv1.DELETE("/articles/:id", articles.Delete)
		apiv1.PUT("/articles/:id", articles.Update)
		apiv1.GET("/articles/:id", articles.Get)
		apiv1.GET("/articles", articles.List)

	}

	return r
}
