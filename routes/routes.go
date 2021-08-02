package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/sergiolucena1/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		blogposts:= main.Group("blogposts")
		{
			blogposts.GET("/:id", controller.ShowBlogPost) //  mostrar blogposts por id
			blogposts.GET("/", controller.ShowBlogPosts) // mostrar todos os blogposts
			blogposts.POST("/", controller.CreateBlogPost) //  criar blogposts
			blogposts.PUT("/", controller.UpdateBlogPost) //  atualizar blogposts
			blogposts.DELETE("/:id", controller.DeleteBlogPost) //  deletar blogpost

		}
	}
	return router
}