package te

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Response struct {
	writer   http.ResponseWriter
	status   int
	headers  http.Header
	body     []byte
	file     string
	cookies  []*http.Cookie
	resolved bool
}

// NewResponse creates a new Response instance.
func NewResponse(w *http.ResponseWriter) *Response {
	return &Response{
		writer:  *w,
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
	if res.resolved {
		fmt.Println("Response already resolved")
		return
	}
	_, err := io.Copy(res.writer, *reader)
	if err != nil {
		fmt.Println(err)
		// send error response
		res.SetStatusCode(StatusInternalServerError).SetBody([]byte(err.Error())).send()
	}
	res.resolved = true
}

func (res *Response) GetContentType() string {
	return res.headers.Get("Content-Type")
}

var i = 0

// Helper methods for common response types
func (res *Response) SendString(str string) {
	fmt.Println("SendString", i)
	i++
	if res.resolved {
		fmt.Println("Response already resolved")
		return
	}
	if res.GetContentType() == "" {
		res.SetContentType("text/plain").SetBody([]byte(str)).send()
		return
	} else {
		res.SetBody([]byte(str)).send()
		return
	}
}

func (res *Response) SendBytes(bytes []byte) {
	if res.resolved {
		fmt.Println("Response already resolved")
		return
	}
	//if content type is not set, then detect it
	if res.GetContentType() == "" {
		res.SetContentType(http.DetectContentType(bytes))
	}
	res.SetBody(bytes).send()
}

func (res *Response) SendHTML(html string) {
	if res.resolved {
		fmt.Println("Response already resolved")
		return
	}
	res.SetContentType("text/html").SetBody([]byte(html)).send()
}

func (res *Response) SendText(text string) {
	if res.resolved {
		fmt.Println("Response already resolved")
		return
	}
	res.SetContentType("text/plain").SetBody([]byte(text)).send()
}

func (res *Response) SendJSON(data interface{}) {
	if res.resolved {
		fmt.Println("Response already resolved")
	}
	//check if the data is string
	if _, ok := data.(string); ok {
		res.SetContentType("application/json").SetBody([]byte(data.(string))).send()
		return
	}

	dataBytes, err := json.Marshal(data)
	if err != nil {
		// Send an internal server error with the error message
		res.SetStatusCode(StatusInternalServerError).SetContentType("text/plain").SetBody([]byte(err.Error()))
		res.send()

	}
	res.SetContentType("application/json").SetBody(dataBytes).send()
}

func (res *Response) SendFile(path string) {
	if res.resolved {
		fmt.Println("Response already resolved")
		return
	}
	// read the file bytes, set it to the body and send it
	file, err := os.Open(path)
	if err != nil {
		res.SetStatusCode(StatusNotFound).send()
		return
	}
	defer file.Close()
	fileInfo, _ := file.Stat()
	fileSize := fileInfo.Size()
	bytes := make([]byte, fileSize)
	_, err = file.Read(bytes)
	if err != nil {
		res.SetStatusCode(StatusInternalServerError).send()
		return
	}
	// get content type
	contentType := http.DetectContentType(bytes)
	res.SetContentType(contentType).SetBody(bytes).send()
}

// Send the response to the client
func (res *Response) send() {
	if res.resolved {
		fmt.Println("Response already resolved")
		return
	}
	res.resolved = true
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
