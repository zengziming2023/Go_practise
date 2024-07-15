package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()
	defer engine.Run()

	// POST /post?ids[a]=1234&ids[b]=hello HTTP/1.1
	//Content-Type: application/x-www-form-urlencoded
	//
	//names[first]=thinkerou&names[second]=tianou
	engine.POST("/post", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")
		fmt.Println(ids, names)
		c.JSON(http.StatusOK, gin.H{
			"ids":   ids,
			"names": names,
		})
	})
}
