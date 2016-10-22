package main

import (
	"log"

	express "github.com/icebob/goexpress"
	"github.com/icebob/goexpress/request"
	"github.com/icebob/goexpress/response"

	logger "github.com/icebob/goexpress/middlewares/log"
)

func main() {
	var app = express.Express()

	app.Use(logger.Simple)

	app.Get("/test", func(req *request.Request, res *response.Response, next func()) {
		res.Write("Hello Test")
		// you can skip closing connection
	})

	app.Get("/", func(req *request.Request, res *response.Response, next func()) {
		res.Write("Hello World")
		// you can skip closing connection
	})

	bindAddress, bindPort := "127.0.0.1", 3000
	log.Printf("Listening at %s:%d...\n", bindAddress, bindPort)
	err := app.Listen(bindPort, bindAddress)
	if err != nil {
		panic(err)
	}
}
