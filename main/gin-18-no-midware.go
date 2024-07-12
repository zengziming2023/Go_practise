package main

import "github.com/gin-gonic/gin"

func main() {
	engine := gin.New()
	// Default 使用 Logger 和 Recovery 中间件
	//r := gin.Default()
	engine.Run()
}
