package main

import (
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/acme/autocert"
	"log"
	"net/http"
)

func main() {
	engine := gin.Default()

	engine.GET("/encrypt", func(c *gin.Context) {
		c.String(http.StatusOK, "hello gin")
	})

	//log.Fatal(autotls.Run(engine, "example1.com", "example2.com"))

	manager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("example1.com", "example2.com"),
		Cache:      autocert.DirCache("/var/www/.cache"),
	}

	log.Fatal(autotls.RunWithManager(engine, &manager))
}
