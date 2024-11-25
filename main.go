package main

import (
	"gin_gorm/controllers"
	"gin_gorm/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	productRoutes := r.Group("/api/products")
	{
		productRoutes.GET("/", controllers.Index)     // GET /api/products
		productRoutes.GET("/:id", controllers.Show)   // GET /api/products/:id
		productRoutes.POST("/", controllers.Create)   // POST /api/products
		productRoutes.PUT("/:id", controllers.Update) // PUT /api/products/:id
		productRoutes.DELETE("/", controllers.Delete) // DELETE /api/products
	}

	r.Run() // listen and serve on 0.0.0.0:8080
}
