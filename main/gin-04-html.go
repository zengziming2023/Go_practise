package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"reflect"
)

func main() {
	engine := gin.Default()
	engine.LoadHTMLGlob("templates/**/*")

	//engine.LoadHTMLFiles("template/index.tmpl", "template/posts/index.tmpl", "template/users/index.tmpl")
	engine.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	engine.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "Posts",
		})
	})

	engine.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "Users",
		})
	})

	engine.Run()
}

func getLoadedTemplates(engine *gin.Engine) []string {
	var templates []string

	// 使用反射访问 gin.Engine 的私有字段
	renderValue := reflect.ValueOf(engine.HTMLRender)
	if renderValue.Kind() == reflect.Ptr {
		renderValue = renderValue.Elem()
	}

	templatesField := renderValue.FieldByName("Template")
	if templatesField.IsValid() {
		tmpl, ok := templatesField.Interface().(*template.Template)
		if ok && tmpl != nil {
			for _, t := range tmpl.Templates() {
				templates = append(templates, t.Name())
			}
		}
	}

	return templates
}
