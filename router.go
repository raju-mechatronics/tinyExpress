package te

import "fmt"

/*
	router logic
	every router will have some default middleware
	router.use() will add middleware to the middleware
	if and only if nextRouter is nil & pathRouter is nil
	else router.use will add a router that has that middleware
	and link it to the nextRouter
*/

type router struct {
	middleware []Middleware
	nextRouter *router
	handlePath map[string]pathHandler
}

type pathHandler struct {
	requestHandler map[string]Handler
	routeHandler   *router
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
