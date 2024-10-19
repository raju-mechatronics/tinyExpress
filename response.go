package te

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type Response struct {
	http.Response
	StatusCode int
	Header     *http.Header
	Body       *[]byte
	resolved   bool
	writer     *http.ResponseWriter
}

func NewResponse() *Response {
	return &Response{
		Header:   &http.Header{},
		resolved: false,
	}
}

func (res *Response) SetStatus(code int) {
	res.StatusCode = code
}

func (res *Response) SetHeader(key, value string) {
	res.Header.Set(key, value)
}

func (res *Response) Write(body *[]byte) {
	res.Body = body
}

func (res *Response) Send(data ...[]byte) error {
	w := *res.writer
	if res.resolved {
		return fmt.Errorf("cannot send response more than once")
	}
	for key, values := range *res.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}
	w.WriteHeader(res.StatusCode)
	if len(data) > 0 {
		_, err := w.Write(data[0])
		if err != nil {
			return err
		}
	} else {
		_, err := w.Write(*res.Body)
		if err != nil {
			return err
		}
	}
	res.resolved = true
	return nil
}

func (res *Response) IsResolved() bool {
	return res.resolved
}

// ServeFile serves a static file to the client
func (res *Response) ServeFile(filePath string) error {
	w := *res.writer
	if res.resolved {
		return fmt.Errorf("cannot send response more than once")
	}
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return nil
	}

	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return nil
	}

	http.ServeContent(w, &http.Request{}, filepath.Base(filePath), fileInfo.ModTime(), file)
	res.resolved = true
	return nil
}

// JSON sends a JSON response to the client
func (res *Response) JSON(data interface{}) error {
	w := *res.writer
	if res.resolved {
		return fmt.Errorf("cannot send response more than once")
	}
	res.Header.Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = w.Write(jsonData)
	if err != nil {
		return err
	}
	res.resolved = true
	return nil
}

// Pipe pipes data from an io.Reader to the response
func (res *Response) Pipe(reader io.Reader) error {
	w := *res.writer
	if res.resolved {
		return fmt.Errorf("cannot send response more than once")
	}
	_, err := io.Copy(w, reader)
	if err != nil {
		return err
	}
	res.resolved = true
	return nil
}
