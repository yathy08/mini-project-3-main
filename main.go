// main.go
package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yathy08/mini-project3/internal/handler"
)

func main() {
	r := gin.Default()
	r.GET("/users", handler.GetAll)
	r.GET("/users/:id", handler.GetByID)
	r.POST("/users", handler.Create)
	r.PUT("/users/:id", handler.Update)
	r.DELETE("/users/:id", handler.Delete)

	log.Println("Server is running on http://localhost:3000")
	r.Run(":3000")
}
