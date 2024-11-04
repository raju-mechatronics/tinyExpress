package te

import (
	"fmt"
	"testing"
)

func TestApp(t *testing.T) {
	app := App()

	if app == nil {
		t.Error("Expected app to be created")
	}

	//app.UseMiddleWare(func(req *Request, res *Response) {
	//	log := ""
	//
	//	log += req.Method + " " + req.URL.Path + "\n"
	//	next := *req.Next
	//	if next != nil {
	//		next()
	//	}
	//})

	app.Get("/h", func(req *Request, res *Response) {
		res.SendString("Hello World")
	})

	err := app.Listen()
	fmt.Println("err=>", err)

}
