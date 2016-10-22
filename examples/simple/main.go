package main

import (
	express "github.com/icebob/goexpress"
	"github.com/icebob/goexpress/request"
	"github.com/icebob/goexpress/response"

	"github.com/icebob/goexpress/middlewares/log"
)

func main() {
	var app = express.Express()

	app.Use(log.Simple)

	app.Get("/test", func(req *request.Request, res *response.Response, next func()) {
		res.Write("Hello Test")
		// you can skip closing connection
	})

	app.Get("/", func(req *request.Request, res *response.Response, next func()) {
		res.Write("Hello World")
		// you can skip closing connection
	})

	app.Listen("3000", "127.0.0.1")
}
