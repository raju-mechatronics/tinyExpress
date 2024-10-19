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
func teLog(option *LogOption) te.Middleware {
	return func(req *te.Request, res *te.Response, next func()) {
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

		next()
	}
}
