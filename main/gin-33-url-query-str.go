package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Person3 struct {
	Name string `form:"name"`
	Addr string `form:"addr"`
}

func main() {
	engine := gin.Default()
	defer engine.Run()

	engine.Any("/testing", func(c *gin.Context) {
		var person Person3
		if err := c.ShouldBindQuery(&person); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"person": person})
	})
	// curl "localhost:8080/testing?name=nick&addr=xxxx"
}
