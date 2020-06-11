package main

import (
	"book-gorm-postgres-gin/controllers" // new
	"book-gorm-postgres-gin/setup"       // new

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func main() {
	r := gin.Default()
	db := setup.SetupModels() // new
	// Provide db variable to controllers
	r.Use(CORS(db))
	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)       // create
	r.GET("/books/:id", controllers.FindBook)      // find by id
	r.PATCH("/books/:id", controllers.UpdateBook)  // update by id
	r.DELETE("/books/:id", controllers.DeleteBook) // delete by id
	r.Run(":8118")
}

func CORS(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Set("db", db)
		c.Next()
	}
}
