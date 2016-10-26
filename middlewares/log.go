package middlewares

/*
	Inspired by Morgan
	https://github.com/expressjs/morgan/blob/master/index.js
*/

import (
	"log"
	"strings"

	"time"

	"github.com/icebob/goexpress/request"
	"github.com/icebob/goexpress/response"
	"github.com/icebob/goexpress/router"
)

const (
	LOGTYPE_TINY byte = iota
	LOGTYPE_DEV
	LOGTYPE_DEFAULT
	LOGTYPE_COMMON
	LOGTYPE_COMBINED
)

func Log(logType byte) router.Middleware {

	return func(req *request.Request, res *response.Response, next func()) {
		res.SetProp("startTime", time.Now())
		res.AddHeaderListener(func() {
			elapsed := time.Since(res.GetProp("startTime").(time.Time))
			switch logType {
			case LOGTYPE_TINY:
				log.Printf("%s %s %d %d - %s", strings.ToUpper(req.Method), req.URL, res.StatusCode, res.ContentLength, elapsed)
			case LOGTYPE_DEV:
				var color int
				status := res.StatusCode
				switch {
				case status >= 500:
					color = 31
				case status >= 400:
					color = 33
				case status >= 300:
					color = 36
				case status >= 200:
					color = 32
				default:
					color = 0
				}
				log.Printf("\x1b[0m%s %s \x1b[%dm%d\x1b[0m %d - %s", strings.ToUpper(req.Method), req.URL, color, res.StatusCode, res.ContentLength, elapsed)
			}
		})

		next()
	}

}
