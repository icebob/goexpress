package middlewares

import (
	"fmt"
	"net/http"

	"github.com/icebob/goexpress/request"
	"github.com/icebob/goexpress/response"
	"github.com/icebob/goexpress/router"
)

func Static(staticFolder string) router.Middleware {

	handler := http.FileServer(http.Dir(staticFolder))

	return func(req *request.Request, res *response.Response, next func()) {
		fmt.Println("URL: " + req.URL)

		handler.ServeHTTP(res.GetRaw(), req.GetRaw())

		//next()
	}

}
