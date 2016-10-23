package main

import gin "gopkg.in/gin-gonic/gin.v1"

func main() {
	r := gin.Default()
	/*r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})*/
	r.Static("/", "./../static/public")
	//r.StaticFS("/static", http.Dir("./../static/public"))
	r.Run("127.0.0.1:3000") // listen and server on 0.0.0.0:8080
}
