package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthRequire() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("AuthRequire")
		c.Next()
	}
}

func main() {
	// 新建一个没有任何默认中间件的路由
	engine := gin.New()
	defer engine.Run()

	// 全局中间件
	// Logger 中间件将日志写入 gin.DefaultWriter，即使你将 GIN_MODE 设置为 release。
	// By default gin.DefaultWriter = os.Stdout
	engine.Use(gin.Logger())
	// Recovery 中间件会 recover 任何 panic。如果有 panic 的话，会写入 500。
	engine.Use(gin.Recovery())

	// 你可以为每个路由添加任意数量的中间件。
	engine.GET("/benchmark", func(c *gin.Context) {
		fmt.Println("middleWare 1, ", *(c.Request))
		c.Next()
	}, func(c *gin.Context) {
		fmt.Println("middleWare 2, ", *(c.Request))
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// 认证路由组
	// authorized := r.Group("/", AuthRequired())
	// 和使用以下两行代码的效果完全一样:
	group := engine.Group("/")
	// 路由组中间件! 在此例中，我们在 "authorized" 路由组中使用自定义创建的
	// AuthRequired() 中间件
	group.Use(AuthRequire())
	{
		group.GET("/login", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status": "login ok",
			})
		})

		// 嵌套路由组
		testing := group.Group("/testing")
		testing.GET("/analytics", func(c *gin.Context) {
			fmt.Println("analytics ok")
			c.String(http.StatusOK, "analytics ok")
		})
	}

}
