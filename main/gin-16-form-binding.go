package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type StructA struct {
	FieldA string `form:"field_a"`
}

type StructB struct {
	NestedStruct StructA
	FieldB       string `form:"field_b"`
}

type StructC struct {
	NestedStructPointer *StructA
	FieldC              string `form:"field_c"`
}

type StructD struct {
	NestedAnonyStruct struct {
		FieldX string `form:"field_x"`
	}
	FieldD string `form:"field_d"`
}

func GetDataB(c *gin.Context) {
	var data StructB
	c.Bind(&data)
	c.JSON(http.StatusOK, gin.H{
		"a": data.NestedStruct,
		"b": data.FieldB,
	})
}

func GetDataC(c *gin.Context) {
	var data StructC
	c.Bind(&data)
	c.JSON(http.StatusOK, gin.H{
		"a": data.NestedStructPointer,
		"c": data.FieldC,
	})
}

func GetDataD(c *gin.Context) {
	var data StructD
	c.Bind(&data)
	c.JSON(http.StatusOK, gin.H{
		"a": data.NestedAnonyStruct,
		"d": data.FieldD,
	})
}

func main() {
	engine := gin.Default()
	engine.GET("/getb", GetDataB)
	engine.GET("/getc", GetDataC)
	engine.GET("/getd", GetDataD)
	engine.Run()

	//$ curl "http://localhost:8080/getb?field_a=hello&field_b=world"
	//{"a":{"FieldA":"hello"},"b":"world"}
	//$ curl "http://localhost:8080/getc?field_a=hello&field_c=world"
	//{"a":{"FieldA":"hello"},"c":"world"}
	//$ curl "http://localhost:8080/getd?field_x=hello&field_d=world"
	//{"d":"world","x":{"FieldX":"hello"}}
}
