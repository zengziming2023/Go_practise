package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Person2 struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func main() {
	engine := gin.Default()
	engine.GET("/testing", func(c *gin.Context) {
		var person Person2
		// 如果是 `GET` 请求，只使用 `Form` 绑定引擎（`query`）。
		// 如果是 `POST` 请求，首先检查 `content-type` 是否为 `JSON` 或 `XML`，然后再使用 `Form`（`form-data`）。
		// 查看更多：https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L88
		err := c.ShouldBind(&person)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"person": person})
	})

	engine.Run()

	//curl -X GET "localhost:8080/testing?name=appleboy&address=xyz&birthday=1992-03-15"
}
