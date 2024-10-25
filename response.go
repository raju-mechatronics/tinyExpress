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

func NewResponse(w *http.ResponseWriter) *Response {
	return &Response{
		resolved: false,
		writer:   w,
	}
}
