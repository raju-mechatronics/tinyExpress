package te

import (
	"fmt"
	"net/http"
	"time"
)

var server *http.Server

type CustomHandler struct {
}

func (CH *CustomHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	fmt.Println(r.URL.Path)
	fmt.Println(r.URL.Query())
	fmt.Println(r.URL.Host, r.URL)
	fmt.Println(r.Response)
	time.Sleep(time.Second * 5)
	rw.Header().Set("Status", "Failed to load")
	rw.WriteHeader(400)
	rw.Write([]byte("hello failed"))
}

func main() {
	server = &(http.Server{
		Addr:    "127.0.0.1:4000",
		Handler: &CustomHandler{},
	})

	server.ListenAndServe()

	router := Router{}

	handler := func(req *Request, res *Response, next ...func()) {
		fmt.Println(req, res)
	}

	router.Get("/hello", handler)

	//tinyexpress.Main()
}
