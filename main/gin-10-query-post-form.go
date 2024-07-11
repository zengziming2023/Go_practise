package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// POST /post?id=1234&page=1 HTTP/1.1
// Content-Type: application/x-www-form-urlencoded
//
// name=manu&message=this_is_great
func main() {
	engine := gin.Default()
	engine.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")
		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
		})
	})

	engine.Run()

	// curl --form message=send_msg --form name=send_nick http://localhost:8080/post?id=1234&page=1
}
