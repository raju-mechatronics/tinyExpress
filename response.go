package te

import (
	"net/http"
)

type Response struct {
	http.Response
	statusCode int
	header     *http.Header
	body       *[]byte
	resolved   bool
	writer     *http.ResponseWriter
}

func NewResponse() *Response {
	return &Response{
		resolved: false,
	}
}
