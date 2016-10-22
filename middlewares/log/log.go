package log

import (
	"log"
	"strings"

	"time"

	"github.com/icebob/goexpress/request"
	"github.com/icebob/goexpress/response"
)

func Simple(req *request.Request, res *response.Response, next func()) {
	res.SetProp("startAt", time.Now())
	res.AddFinishedListener(func() {
		elapsed := time.Since(res.GetProp("startAt").(time.Time))
		log.Printf("%s %s %d %d - %s", strings.ToUpper(req.Method), req.URL, res.Header.StatusCode, res.Header.Length, elapsed)
	})

	next()
}
