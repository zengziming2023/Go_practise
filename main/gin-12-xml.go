package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"net/http"
)

func main() {
	engine := gin.Default()

	// gin.H 是 map[string]interface{} 的一种快捷方式
	engine.GET("/someJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hey", "status": http.StatusOK,
		})
	})

	engine.GET("/moreJSON", func(c *gin.Context) {
		// 你也可以使用一个结构体
		var msg struct {
			Name    string `json:"user"`
			Message string `json:"message"`
			Number  int    `json:"number"`
		}
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123

		// 注意 msg.Name 在 JSON 中变成了 "user"
		// 将输出：{"user": "Lena", "Message": "hey", "Number": 123}
		c.JSON(http.StatusOK, msg)
	})

	engine.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{
			"message": "hey", "status": http.StatusOK,
		})
	})

	engine.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{
			"message": "hey", "status": http.StatusOK,
		})
	})

	engine.GET("/someProtoBuf", func(c *gin.Context) {
		reps := []int64{int64(1), 2, 3}
		label := "test"
		// protobuf 的具体定义写在 testdata/protoexample 文件中。
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		// 请注意，数据在响应中变为二进制数据
		// 将输出被 protoexample.Test protobuf 序列化了的数据
		c.ProtoBuf(http.StatusOK, data)
	})

	engine.Run()
}
