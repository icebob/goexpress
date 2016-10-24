package main

import (
	"net/http"

	gin "gopkg.in/gin-gonic/gin.v1"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	//r := gin.Default()
	r := gin.New()
	//r.Use(gin.Logger())

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
		/*c.JSON(200, gin.H{
			"message": "pong",
		})*/
	})

	//r.Static("/app", "./../static/public")
	//r.StaticFS("/app", http.Dir("./../static/public"))
	r.Run("127.0.0.1:3000") // listen and server on 0.0.0.0:8080
}
