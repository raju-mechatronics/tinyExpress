package te

type AppConfig struct {
	Port          int
	Host          string
	AllowedMethod []string
}

type Handler func(req *Request, res *Response)
type Middleware func(req *Request, res *Response, next func())

type Router interface {
	Resolve(req *Request, res *Response)
}
