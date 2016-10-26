// Response package provides the core functionality of handling
// the client connection, chunked response and other features
package response

import (
	"context"
	"encoding/json"
	"net/http"
)

// Response Structure extends basic http.ResponseWriter interface
// It encapsulates Header and Cookie class for direct access
type Response struct {
	request  *http.Request
	response http.ResponseWriter
	header   http.Header
	//Cookie            *cookie.Cookie
	Context           context.Context
	Props             map[string]interface{}
	StatusCode        int
	ContentLength     int
	headerSent        bool
	ended             bool
	headerListeners   []func()
	finishedListeners []func()
}

// Intialise the Response Struct, requires the Hijacked buffer,
// connection and Response interface
func (res *Response) Init(rsp http.ResponseWriter, req *http.Request) *Response {
	res.request = req
	res.response = rsp
	res.header = rsp.Header()
	res.Context = req.Context()
	res.Props = make(map[string]interface{})
	//res.Cookie = &cookie.Cookie{}
	//res.Cookie.Init(res, req)
	res.StatusCode = 0
	res.ContentLength = 0
	res.headerSent = false
	res.ended = false

	return res
}

/*
// This function is for internal Use by Cookie Struct
func (res *Response) AddCookie(key string, value string) {
	res.Header.AppendCookie(key, value)
}*/

func (res *Response) AddHeaderListener(callback func()) {
	res.headerListeners = append(res.headerListeners, callback)
}

func (res *Response) SetProp(key string, value interface{}) {
	//res.request.WithContext(context.WithValue(res.request.Context(), key, value))
	res.Props[key] = value
}

func (res *Response) GetProp(key string) interface{} {
	//return res.Context.Value(key)
	return res.Props[key]
}

func (res *Response) Send(content string) *Response {
	res.sendContent(200, "text/html; charset=utf-8", []byte(content))
	return res
}

func (res *Response) sendContent(status int, content_type string, content []byte) {
	if !res.headerSent {
		res.header["Content-Type"] = []string{content_type}
		res.response.WriteHeader(status)
		res.StatusCode = status

		// Call headerListeners
		for _, cb := range res.headerListeners {
			cb()
		}

		// Clear listeners
		res.headerListeners = res.headerListeners[:0]

		res.headerSent = true
	}
	n, err := res.response.Write(content)
	if err == nil {
		res.ContentLength += n
		res.ended = true
	}
}

func (res *Response) HasEnded() bool {
	return res.ended
}

// Redirects a request, takes the url as the Location
func (res *Response) Redirect(url string) *Response {
	http.Redirect(res.response, res.request, url, 301)
	return res
}

// A helper for middlewares to get the original http.ResponseWriter
func (res *Response) GetRaw() http.ResponseWriter {
	return res.response
}

// Send Error, takes HTTP status and a string content
func (res *Response) Error(status int, str string) {
	res.sendContent(status, "text/html", []byte(str))
}

// Send JSON response, takes interface as input
func (res *Response) JSON(content interface{}) {
	output, err := json.Marshal(content)
	if err != nil {
		res.sendContent(500, "application/json", []byte(""))
	} else {
		res.sendContent(200, "application/json", output)
	}
}
