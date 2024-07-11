package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {
	engine := gin.Default()

	engine.POST("/login", func(c *gin.Context) {
		// 你可以使用显式绑定声明绑定 multipart form：
		// c.ShouldBindWith(&form, binding.Form)
		// 或者简单地使用 ShouldBind 方法自动绑定：
		var form LoginForm
		// 在这种情况下，将自动选择合适的绑定
		if err := c.ShouldBind(&form); err == nil {
			fmt.Println(form)
			if form.User == "user" && form.Password == "password" {
				c.JSON(http.StatusOK, gin.H{
					"status": "ur login.",
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"status": "unauthorized",
				})
			}
		}
	})

	fmt.Println("gin server start.")
	engine.Run()
	fmt.Println("gin server end.")

	// $ curl -v --form user=user --form password=password http://localhost:8080/login
}
