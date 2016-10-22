package log

import (
	"log"

	"github.com/icebob/goexpress/request"
	"github.com/icebob/goexpress/response"
)

func Simple(req *request.Request, res *response.Response, next func()) {
	log.Printf("%s %s", req.Method, req.URL)
	next()
}
