package server

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func server1() {
	// 初始化空的服务器
	app := gin.New()
	log.Printf("启动服务器在 http address: %s", ":8080")
	log.Printf(http.ListenAndServe(":8080", app).Error())
}

func server2() {
	// 初始化空的服务器
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World")
	})
	r.Run()
}

func server3() {
	// 初始化空的服务器
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	r.Run() // listen and server on 0.0.0.0:8080
}
