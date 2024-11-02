package te

import (
	"encoding/json"
	"io"
	"net/http"
)

type ResponseEventType string

const (
	BeforeSend    ResponseEventType = "beforeSend"
	End           ResponseEventType = "finish"
	ErrorResponse ResponseEventType = "error"
)

type Response struct {
	writer         http.ResponseWriter
	status         int
	headers        http.Header
	body           []byte
	file           string
	cookies        []*http.Cookie
	resolved       bool
	eventListeners map[ResponseEventType][]func(req *Request, res *Response)
}

// NewResponse creates a new Response instance.
func NewResponse(w http.ResponseWriter) *Response {
	return &Response{
		writer:  w,
		headers: make(http.Header),
	}
}

func (res *Response) SetHeader(key, value string) *Response {
	res.headers.Set(key, value)
	return res
}

func (res *Response) SetBody(body []byte) *Response {
	res.body = body
	return res
}

func (res *Response) AppendBody(body []byte) *Response {
	res.body = append(res.body, body...)
	return res
}

func (res *Response) SetContentType(contentType string) *Response {
	res.SetHeader("Content-Type", contentType)
	return res
}

func (res *Response) SetCookie(cookie ...*http.Cookie) *Response {
	res.cookies = append(res.cookies, cookie...)
	return res
}

func (res *Response) SetStatusCode(code int) *Response {
	res.status = code
	return res
}

func (res *Response) Pipe(reader *io.Reader) {
	io.Copy(res.writer, *reader)
	res.resolved = true
}

func (res *Response) GetContentType() string {
	return res.headers.Get("Content-Type")
}

// Helper methods for common response types
func (res *Response) SendString(str string) {
	if res.GetContentType() == "" {
		res.SetContentType("text/plain").SetBody([]byte(str)).send()
	} else {
		res.SetBody([]byte(str)).send()
	}
}

func (res *Response) SendBytes(bytes []byte) {
	res.SetBody(bytes).send()
}

func (res *Response) SendHTML(html string) {
	res.SetContentType("text/html").SetBody([]byte(html)).send()
}

func (res *Response) SendText(text string) {
	res.SetContentType("text/plain").SetBody([]byte(text)).send()
}

func (res *Response) SendJSON(data interface{}) error {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		// Send an internal server error with the error message
		res.SetStatusCode(StatusInternalServerError).SetContentType("text/plain").SetBody([]byte(err.Error()))
		res.send()
		return err
	}
	res.SetContentType("application/json").SetBody(dataBytes).send()
	return nil
}

func (res *Response) SendFile(path string) error {

}

// Send the response to the client
func (res *Response) send() {
	// Set status code
	if res.status != 0 {
		res.writer.WriteHeader(res.status)
	}

	// Apply headers
	for key, values := range res.headers {
		for _, value := range values {
			res.writer.Header().Add(key, value)
		}
	}

	// Apply cookies
	for _, cookie := range res.cookies {
		http.SetCookie(res.writer, cookie)
	}

	// Write body
	if len(res.body) > 0 {
		res.writer.Write(res.body)
	}

	res.resolved = true
}

// Check if the response has been resolved
func (res *Response) IsResolved() bool {
	return res.resolved
}

// New method for redirecting
func (res *Response) Redirect(url string, statusCode int) {
	res.SetStatusCode(statusCode)
	http.Redirect(res.writer, &http.Request{}, url, statusCode)
	res.resolved = true
}

// AddEventListener adds a listener for a specific event type
func (res *Response) On(event ResponseEventType, listener func(req *Request, res *Response)) {
	if res.eventListeners == nil {
		res.eventListeners = make(map[ResponseEventType][]func(req *Request, res *Response))
	}
	res.eventListeners[event] = append(res.eventListeners[event], listener)
}
