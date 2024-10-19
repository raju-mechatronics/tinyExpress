package te

import "testing"

func TestApp(t *testing.T) {
	app := App()
	app.Get("", []Middleware{}, func(req *Request, res *Response) {
		res.Send("Hello")
	})

	if app == nil {
		t.Error("Expected app to be created")
	}
}
