package main

import (
	"fmt"
	"time"

	"github.com/icebob/goexpress"

	"github.com/icebob/goexpress/middlewares"
	"github.com/icebob/goexpress/request"
	"github.com/icebob/goexpress/response"
)

func main() {
	//runtime.GOMAXPROCS(8)

	var app = goexpress.Express()

	app.Use(middlewares.Log(middlewares.LOGTYPE_TINY))

	app.Get("/test", func(req *request.Request, res *response.Response, next func()) {
		res.Send("Hello Test")
	})

	app.Get("/", func(req *request.Request, res *response.Response, next func()) {
		time.Sleep(200 * time.Millisecond)
		res.Send("Hello World")
	})

	//bindAddress, bindPort := "127.0.0.1", 3000
	bindAddress, bindPort := "0.0.0.0", 3000
	fmt.Printf("Listening at %s:%d...\n", bindAddress, bindPort)
	err := app.Listen(bindPort, bindAddress)
	if err != nil {
		panic(err)
	}
}
