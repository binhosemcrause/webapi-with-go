package routes

import (
	"github.com/binhosemcrause/webapi-with-go/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		books := main.Group("books")
		{
			books.GET("/", controllers.AllBooks)
			books.GET("/:id", controllers.GetBook)
			/*books.POST("/", controllers.NewBook)
			books.PUT("/", controllers.EditBook)
			books.DELETE("/:id", controllers.DeleteBook)*/
		}
	}

	return router
}
