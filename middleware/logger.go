package middleware

import (
	"time"
	te "tinyExpress"
)

type LogOption struct {
	TimeStamp  bool
	FullUrl    bool
	Path       bool
	Method     bool
	StatusCode bool
}

// Log the details of request
func TeLog(option LogOption) te.Handler {
	return func(req *te.Request, res *te.Response) {
		//log the request
		log := ""

		if option.TimeStamp {
			log += "Time: " + time.Now().String() + " "
		}

		if option.FullUrl {
			log += "Url: " + req.URL.String() + " "
		}

		if option.Path {
			log += "Path: " + req.URL.Path + " "
		}

		if option.Method {
			log += "Method: " + req.Method + " "
		}

		if req.Next != nil {
			next := *req.Next
			next()
		}
	}
}
