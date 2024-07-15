package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	defer engine.Run()

	engine.GET("/cookie", func(c *gin.Context) {
		cookie, err := c.Cookie("gin_cookiep")

		if err != nil {
			fmt.Println(err)
			cookie = "NotSet"
			c.SetCookie("gin_cookiep", "test", 3600, "/", "localhost", false, true)
		}
		fmt.Printf("Cookie value: %s \n", cookie)
		c.String(200, cookie)
	})

}
