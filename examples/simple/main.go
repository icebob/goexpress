package main

import (
	"log"

	express "github.com/icebob/goexpress"
	logger "github.com/icebob/goexpress/middlewares/log"
	"github.com/icebob/goexpress/request"
	"github.com/icebob/goexpress/response"
)

func main() {
	//runtime.GOMAXPROCS(8)

	var app = express.Express()

	app.Use(logger.Simple)

	app.Get("/test", func(req *request.Request, res *response.Response, next func()) {
		res.Send("Hello Test")
	})

	app.Get("/chunked", func(req *request.Request, res *response.Response, next func()) {
		res.WriteChunk("Hello World\n")
		//time.Sleep(5 * time.Second)
		res.WriteChunk("Hello World2\n")
	})

	app.Get("/", func(req *request.Request, res *response.Response, next func()) {
		res.Send("Hello World")
	})

	bindAddress, bindPort := "127.0.0.1", 3000
	//bindAddress, bindPort := "0.0.0.0", 3000
	log.Printf("Listening at %s:%d...\n", bindAddress, bindPort)
	err := app.Listen(bindPort, bindAddress)
	if err != nil {
		panic(err)
	}
}
