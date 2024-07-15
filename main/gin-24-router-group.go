package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()
	defer engine.Run()

	v1 := engine.Group("/v1")
	{
		v1.POST("/login", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status": "v1 login",
			})
		})
		v1.POST("/submit", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status": "v1 submit",
			})
		})
	}

	v2 := engine.Group("/v2")
	{
		v2.POST("/login", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status": "v2 login",
			})
		})
		v2.POST("/submit", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status": "v2 submit",
			})
		})
	}

}
