package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()
	defer engine.Run()

	engine.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})

	engine.GET("/test2", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/foo")
	})

	engine.GET("/test3", func(c *gin.Context) {
		c.Request.URL.Path = "/foo"
		engine.HandleContext(c)
	})

	engine.GET("/foo", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "foo",
		})
	})
}
