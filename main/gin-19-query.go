package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()
	defer engine.Run()

	engine.GET("/welcome", func(c *gin.Context) {
		firstName := c.DefaultQuery("firstname", "default")
		lastName := c.Query("lastname") // // c.Request.URL.Query().Get("lastname") 的一种快捷方式
		c.String(http.StatusOK, "Hello %s %s", firstName, lastName)
	})

	// /welcome?firstname=Jane&lastname=Doe
}
