package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		// 设置 example 变量
		c.Set("example", "12345")
		// 请求前...

		c.Next()

		// 请求后
		latency := time.Since(t)
		log.Print("cost: ", latency)

		// 获取发送的 status
		status := c.Writer.Status()
		log.Println("status: ", status)

	}
}

func main() {
	engine := gin.Default() //New()
	defer engine.Run()

	engine.Use(Logger())
	engine.GET("/", func(c *gin.Context) {
		example := c.MustGet("example").(string)
		log.Println(example)
		c.JSON(http.StatusOK, gin.H{"example": example})
	})
}
