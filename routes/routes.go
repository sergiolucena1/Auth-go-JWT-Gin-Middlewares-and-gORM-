package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/sergiolucena1/controllers"
	"github.com/sergiolucena1/middlewares"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		users := main.Group("users") // criando minha rota de usuario
		{
			users.POST("/", controller.CreateUser) // criando usuario
		}
		blogposts:= main.Group("blogposts", middlewares.Auth())
		{
			blogposts.GET("/:id", controller.ShowBlogPost) //  mostrar blogposts por id
			blogposts.GET("/", controller.ShowBlogPosts) // mostrar todos os blogposts
			blogposts.POST("/", controller.CreateBlogPost) //  criar blogposts
			blogposts.PUT("/", controller.UpdateBlogPost) //  atualizar blogposts
			blogposts.DELETE("/:id", controller.DeleteBlogPost) //  deletar blogpost
		}
		main.POST("login",controller.Login)
	}
	return router
}
