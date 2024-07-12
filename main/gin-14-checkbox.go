package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type myForm struct {
	Colors []string `form:"colors[]"`
}

func formHeader(c *gin.Context) {
	var fakeForm myForm
	err := c.ShouldBind(&fakeForm)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"colors": fakeForm.Colors,
	})
}

func main() {
	engine := gin.Default()
	engine.LoadHTMLGlob("views/*")
	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "form.html", nil)
	})
	engine.POST("/", formHeader)
	engine.Run()
}
