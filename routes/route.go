package routes

import (
	"book-management-api/controllers"

	"github.com/gin-gonic/gin"
)

func Server() *gin.Engine {
	route := gin.Default()

	books := route.Group("/api")
	{
		books.GET("/books", controllers.GetBooks)
		books.GET("/book/:id", controllers.GetBookByID)
		books.POST("/book", controllers.CreateBook)
		books.PUT("/book/:id", controllers.EditBook)
		books.DELETE("/book/:id", controllers.DeleteBook)
	}

	category := route.Group("/api")
	{
		category.GET("/categories", controllers.GetCategories)
		category.POST("/category", controllers.CreateCategory)
		category.PUT("/category/:id", controllers.EditCategory)
		category.DELETE("/category/:id", controllers.DeleteCategory)
	}

	return route
}
