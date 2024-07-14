package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	engine := gin.Default()
	defer engine.Run()

	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	engine.POST("/foo", func(c *gin.Context) {
		c.JSON(http.StatusOK, "foo")
	})

	engine.GET("/bar", func(c *gin.Context) {
		c.JSON(http.StatusOK, "bar")
	})

	engine.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, "status")
	})
}
