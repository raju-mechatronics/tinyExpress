package te

import "fmt"

/*
	resolver interface extends the resolve method

	will be implemented by the router struct and similar struct that needs to resolve the request
*/
type resolver interface {
	resolve(req *Request, res *Response)
}

type routeMap map[string]resolver

func (rmap *routeMap) addRoute(path string, resolver resolver) {
	if (*rmap)[path] != nil {
		fmt.Println("warn: the path already have a resolver")
		return
	}
	(*rmap)[path] = resolver
}

func (rmap routeMap) getRoute(path string) resolver {
	var res resolver
	if rmap[path] == nil {
		return nil
	}
	res = rmap[path]
	return res
}

type pathHandler struct {
	requestHandler map[RequestType]routeMap
}

func (ph *pathHandler) resolve(req *Request, res *Response) resolver {
	if res.resolved {
		fmt.Println("warn: the response is resolve => please use return on next and res.send function call")
		return nil
	}

	method := req.Method
	var handler routeMap

	if ph.requestHandler[method] == nil {
		if ph.requestHandler[RequestTypeAny] == nil {
			return nil
		} else {
			handler = ph.requestHandler[RequestTypeAny]
		}
	} else {
		handler = ph.requestHandler[method]
	}

	if handler == nil {
		return nil
	}

	return handler.getRoute(req.Path)

}

type router struct {
	middleware []Middleware
	handlePath pathHandler
}

func (r *router) handleMiddleware(req *Request, res *Response) {
	if res.resolved {
		fmt.Println("warn: the response is resolve => please use return on next and res.send function call")
		return
	}
	if len(r.middleware) > 0 {
		curFunc := r.middleware[0]
		curIndex := 0
		var next func()
		next = func() {
			if res.resolved {
				fmt.Println("warn: the response is resolve => please use return on next and res.send function call")
				return
			}
			curIndex++
			if curIndex < len(r.middleware) {
				nextFunc := r.middleware[curIndex]
				nextFunc(req, res, next)
			} else {
				r.resolve(req, res)
			}
		}
		curFunc(req, res, next)
	}
}

func (r *router) resolve(req *Request, res *Response) {
	r.handleMiddleware(req, res)

	if res.resolved {
		return
	}

}

//func (hwm *handlerWithMiddleware) resolve(req *Request, res *Response) {
//	if res.resolved {
//		fmt.Println("warn: the response is resolve => please use return on next and res.send function call")
//		return
//	}
//	if len(hwm.middleware) > 0 {
//		curFunc := hwm.middleware[0]
//		curIndex := 0
//		var next func()
//		next = func() {
//			if res.resolved {
//				fmt.Println("warn: the response is resolve => please use return on next and res.send function call")
//				return
//			}
//			curIndex++
//			if curIndex < len(hwm.middleware) {
//				nextFunc := hwm.middleware[curIndex]
//				nextFunc(req, res, next)
//			} else {
//				hwm.handler(req, res)
//			}
//		}
//		curFunc(req, res, next)
//	} else {
//		hwm.handler(req, res)
//	}
//}

func (r *router) Use(path string, middleware []Middleware, router router) {
	if r.handlePath[path].routeHandler != nil {
		fmt.Println("warn: the path already have a router")
		return
	}
	r.handlePath[path] = pathHandler{
		routeHandler: &router,
	}
}

func (r *router) Get(path string, middlewares []Middleware, handler Handler) {
	if r.handlePath[path].requestHandler["GET"] != nil {
		fmt.Println("warn: the path already have a handler")
		return
	}
	r.handlePath[path].requestHandler["GET"] = handler
}

func (r *router) Post(path string, middlewares []Middleware, handler Handler) {

}

func (r *router) Delete(path string, middlewares []Middleware, handler Handler) {

}

func (r *router) Put(path string, middlewares []Middleware, handler Handler) {

}
func (r *router) Patch(path string, middlewares []Middleware, handler Handler) {

}
