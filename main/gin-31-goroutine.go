package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	engine := gin.Default()
	defer engine.Run()

	// 当在中间件或 handler 中启动新的 Goroutine 时，不能使用原始的上下文，必须使用只读副本。
	engine.GET("/long_async", func(c *gin.Context) {
		// 创建在 goroutine 中使用的副本
		cCp := c.Copy()
		go func() {
			// 用 time.Sleep() 模拟一个长任务。
			time.Sleep(5 * time.Second)
			// 请注意您使用的是复制的上下文 "cCp"，这一点很重要
			log.Println("Done! in path :", cCp.Request.URL.Path)
		}()
	})

	engine.GET("/long_sync", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		// 因为没有使用 goroutine，不需要拷贝上下文
		log.Println("Done! in path :", c.Request.URL.Path)
	})
}
