[![GoDoc](https://godoc.org/github.com/icebob/goexpress?status.svg)](https://godoc.org/github.com/icebob/goexpress)
# goexpress
An Express JS Style HTTP server implementation in Golang. The package make use of similar framework convention as they are in express-js. People switching from NodeJS to Golang often end up in a bad learning curve to start building their webapps, this project is meant to ease things up, its a light weight framework which can be extended to do add any number of functionality.

Original code: https://github.com/DronRathore/goexpress

## Hello World
```go
package main

import (
  "github.com/icebob/goexpress"
	"github.com/icebob/goexpress/middlewares"
  "github.com/icebob/goexpress/request"
  "github.com/icebob/goexpress/response"
)

func main() {
  var app = goexpress.Express()

	app.Use(middlewares.Log(middlewares.LOGTYPE_DEV))

  app.Get("/", func(req *request.Request, res *response.Response, next func()){
    res.Write("Hello World")
    // you can skip closing connection
  })
  app.Listen("8080", "0.0.0.0")
}
```

## Router
The router works in the similar way as it does in the express-js. You can have named parameters in the URL or named + regex combo.
```go
func main() {
  var app = goexpress.Express()
  app.Get("/:service/:object([0-9]+)", func(req *request.Request, res *response.Response, next func()){
    res.JSON(req.Params)
  })
  app.Listen("8080", "0.0.0.0")
}
```

## Middleware
You can write custom middlewares, wrappers in the similar fashion. Middlewares can be used to add websocket upgradation lib, session handling lib, static assets server handler
```go
func main() {
  var app = goexpress.Express()
  app.Use(func(req *request.Request, res *response.Response, next func()){
    req.Params["I-Am-Adding-Something"] = "something"
    next()
  })
  app.Get("/:service/:object([0-9]+)", func(req *request.Request, res *response.Response, next func()){
    // json will have the key added
    res.JSON(req.Params)
  })
  app.Listen("8080", "0.0.0.0")
}
```

## Cookies
```go
import (
  express "github.com/icebob/goexpress"
  request "github.com/icebob/goexpress/request"
  response "github.com/icebob/goexpress/response"
  http "net/http"
  Time "time"
)
func main() {
  var app = goexpress.Express()
  app.Use(func(req *request.Request, res *response.Response, next func()){
    var cookie = &http.Cookie{
      Name: "name",
      Value: "value",
      Expires: Time.Unix(0, 0)
    }
    res.Cookie.Add(cookie)
    req.Params["session_id"] = req.Cookies.Get("session_id")
  })
  app.Get("/", func(req *request.Request, res *response.Response, next func()){
    res.Write("Hello World")
  })
  app.Listen("8080", "0.0.0.0")
}
```

## Post Body
```go
func main() {
  var app = goexpress.Express()
  app.Use(func(req *request.Request, res *response.Response, next func()){
    res.Params["I-Am-Adding-Something"] = "something"
    next()
  })
  app.Post("/user/new", func(req *request.Request, res *response.Response, next func()){
    type User struct {
			Name string `json:"name"`
			Email string `json:"email"`
		}
		var list = &User{Name: req.Body["name"], Email: req.Body["email"]}
		res.JSON(list)
  })
  app.Listen("8080", "0.0.0.0")
}
```

## JSON Post
JSON Post data manipulation in golang is slightly different from JS. You have to pass a filler to the decoder, the decoder assumes the data to be in the same format as the filler, if it is not, it throws an error.
```go
func main() {
  var app = goexpress.Express()
  app.Use(func(req *request.Request, res *response.Response, next func()){
    res.Params["I-Am-Adding-Something"] = "something"
    next()
  })
  app.Post("/user/new", func(req *request.Request, res *response.Response, next func()){
    type User struct {
			Name string `json:"name"`
			Email string `json:"email"`
		}
		var list User
		err := req.JSON.Decode(&list) 
		if err != nil {
			res.Error(400, "Invalid JSON")
		} else {
			res.JSON(list)
		}
  })
  app.Listen("8080", "0.0.0.0")
}
```

## File Uploading
The ```response.Response``` struct has a ```GetFile()``` method, which reads a single file at a time, you can make looped calls to retrieve all the files, however this feature is not thoroughly tested, bugs can be reported for the same.


