package main

import (
	"fmt"

	"github.com/icebob/goexpress"
	"github.com/icebob/goexpress/middlewares"
	"github.com/icebob/goexpress/request"
	"github.com/icebob/goexpress/response"
)

func main() {
	//runtime.GOMAXPROCS(8)

	var app = goexpress.Express()

	app.Use(middlewares.Log(middlewares.LOGTYPE_DEV))

	app.Use(middlewares.Static("public"))

	app.Get("/", func(req *request.Request, res *response.Response, next func()) {
		res.Send("Hello World")
	})

	bindAddress, bindPort := "127.0.0.1", 3000
	fmt.Printf("Listening at %s:%d...\n", bindAddress, bindPort)
	err := app.Listen(bindPort, bindAddress)
	if err != nil {
		panic(err)
	}
}
