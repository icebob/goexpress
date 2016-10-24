// Header Package, handles the Response & Request Header
// The package is responsible for setting Response headers
// and pushing the same on the transport buffer
package header

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Header Struct
type Header struct {
	response       http.ResponseWriter
	request        *http.Request
	writer         *bufio.ReadWriter
	bodySent       bool
	basicSent      bool
	transferChunks bool
	StatusCode     int
	ProtoMajor     int
	ProtoMinor     int
	Length         int
	listeners      []func()
}

var statusCodeMap = map[int]string{
	200: "OK",
	201: "Created",
	202: "Accepted",
	204: "No Content",
	205: "Reset Content",
	206: "Partial Content",
	301: "Moved Permanently",
	302: "Found",
	304: "Not Modified",
	305: "Use Proxy",
	306: "Switch Proxy",
	307: "Temporary Redirect",
	308: "Permanent Redirect",
	400: "Bad Request",
	401: "Unauthorized",
	403: "Forbidden",
	404: "NOT FOUND",
	405: "Method Not Allowed",
	413: "Payload Too Large",
	414: "URI Too Long",
	500: "Internal Server Error",
	501: "Not Implemented",
	502: "Bad Gateway",
	503: "Service Unavailaible",
	504: "Gateway Timeout",
	505: "HTTP Version Not Supported",
}

// Initialise with response, request and io buffer
func (h *Header) Init(response http.ResponseWriter, request *http.Request, writer *bufio.ReadWriter) *Header {
	h.response = response
	h.request = request
	h.writer = writer
	h.bodySent = false
	h.basicSent = false
	h.transferChunks = false
	h.ProtoMinor = 1
	h.ProtoMajor = 1
	h.Length = 0
	return h
}

// Sets a header
func (h *Header) Set(key string, value string) *Header {
	h.response.Header().Set(key, value)
	return h
}

// Returns the header
func (h *Header) Get(key string) string {
	return h.response.Header().Get(key)
}

// Deletes a Header
func (h *Header) Del(key string) *Header {
	h.response.Header().Del(key)
	return h
}

func (h *Header) AddListener(callback func()) {
	h.listeners = append(h.listeners, callback)
}

// Add non-chunk response functionality
func (h *Header) SetLength(length int) {
	h.Length = length
}

// Add chunk response functionality
func (h *Header) TranferChunks() {
	h.transferChunks = true
}

func (h *Header) IsTranferChunks() bool {
	return h.transferChunks
}

// Flushes Headers
func (h *Header) FlushHeaders() bool {
	if h.bodySent == true {
		log.Panic("Cannot send headers in middle of body")
		return false
	} else {

		// Call finishedListeners
		for _, cb := range h.listeners {
			cb()
		}

		// Clear listeners
		h.listeners = h.listeners[:0]

		// Send basic headers
		if h.basicSent == false {
			h.sendBasics()
		}
		// write the latest headers
		if h.Get("Content-Type") == "" {
			h.Set("Content-Type", "text/html; charset=utf-8")
		}
		if err := h.response.Header().Write(h.writer); err != nil {
			return false
		} else {
			if h.IsTranferChunks() {
				var chunkSize = fmt.Sprintf("%x", 0)
				h.writer.WriteString(chunkSize + "\r\n")
			}
			h.writer.WriteString("\r\n")
			h.writer.Writer.Flush()
			return true
		}
	}
}

// An internal helper function to set Cookie Header
func (h *Header) AppendCookie(key string, value string) {
	if h.Get(key) != "" {
		h.Set(key, h.Get(key)+";"+value)
	} else {
		h.Set(key, value)
	}
}

// Returns the state of Headers whether they are sent or not
func (h *Header) BasicSent() bool {
	return h.basicSent
}

// Returns the state of response
func (h *Header) CanSendHeader() bool {
	if h.basicSent == true {
		if h.bodySent == false {
			return true
		} else {
			return false
		}
	}
	return true
}

// Sets the HTTP Status of the Request
func (h *Header) SetStatus(code int) {
	h.StatusCode = code
}

func (h *Header) sendBasics() {
	if h.StatusCode == 0 {
		h.StatusCode = 200
	}
	fmt.Fprintf(h.writer, "HTTP/%d.%d %03d %s\r\n", h.ProtoMajor, h.ProtoMinor, h.StatusCode, statusCodeMap[h.StatusCode])
	h.Set("Date", time.Now().UTC().Format(time.RFC1123))
	h.Set("Connection", "keep-alive")
	if h.transferChunks {
		h.Set("Transfer-Encoding", "chunked")
	} else {
		h.Set("Content-Length", strconv.Itoa(h.Length))
	}
	h.basicSent = true
}
