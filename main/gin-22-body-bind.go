package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

type formA struct {
	Foo string `json:"foo" xml:"foo" binding:"required"`
}

type formB struct {
	Bar string `json:"bar" xml:"bar" binding:"required"`
}

// SomeHandler 一般通过调用 c.Request.Body 方法绑定数据，但不能多次调用这个方法。
func SomeHandler(c *gin.Context) {
	objA := formA{}
	objB := formB{}
	// c.ShouldBind 使用了 c.Request.Body，不可重用。
	//if err := c.ShouldBind(&objA); err == nil {
	//	c.String(http.StatusOK, `the body should be formA`)
	//} else if err := c.ShouldBind(&objB); err == nil {
	//	c.String(http.StatusOK, `the body should be formB`)
	//}

	if err := c.ShouldBindWith(&objA, binding.JSON); err == nil {
		c.String(http.StatusOK, `the body should be formA`)
	} else if err := c.ShouldBindWith(&objB, binding.XML); err == nil {
		c.String(http.StatusOK, `the body should be formB`)
	} else {
		c.String(http.StatusBadRequest, "body should be formA or formB")
	}
}

func main() {
	engine := gin.Default()
	defer engine.Run()

	engine.POST("/", SomeHandler)

	// curl -X POST http://localhost:8080/ -H "Content-Type: application/json" -d '{ "bar": "bar bar"}'
	// curl -X POST http://localhost:8080/ \
	//-H "Content-Type: application/xml" \
	// -d '<xml>
	//        <data>
	//            <bar>bar bar</name>
	//        </data>
	//    </xml>'
}
