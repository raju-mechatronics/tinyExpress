package te

type AppConfig struct {
	Port          int
	Host          string
	AllowedMethod []string
}

type Resolver interface {
	Resolve(req *Request, res *Response)
}

type Handler func(req *Request, res *Response)

func (h Handler) Resolve(req *Request, res *Response) {
	h(req, res)
}
