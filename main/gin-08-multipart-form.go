package main

import "github.com/gin-gonic/gin"

func main() {
	engine := gin.Default()
	engine.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	engine.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": message + "_xxx",
			"nick":    nick + "_xxx",
		})
	})

	engine.Run()

	// curl -v --form message=send_msg --form nick=send_nick http://localhost:8080/form_post
	// curl --form message=send_msg --form nick=send_nick http://localhost:8080/form_post
}
