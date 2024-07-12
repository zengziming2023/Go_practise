package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Person struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

func main() {
	engine := gin.Default()

	engine.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "TEST"})
	})

	engine.GET("/:name/:id", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBindUri(&person); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"person": person})
	})
	engine.Run()

	// curl -v localhost:8080/thinkerou/987fbc97-4bed-5078-9f07-9141ba07c9f3
	// curl -v localhost:8080/thinkerou/not-uuid
}
