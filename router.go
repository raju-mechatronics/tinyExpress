package te

import (
	"fmt"
)

type Router struct {
	handler []Resolver
}

func convertToResolver[T Resolver](h ...T) []Resolver {
	resolvers := make([]Resolver, len(h))
	for i, handler := range h {
		resolvers[i] = handler
	}
	return resolvers
}

func (r *Router) add(handler ...Resolver) {
	r.handler = append(r.handler, handler...)
}

func (r *Router) handleMiddleware(req *Request, res *Response) {
	if res.resolved {
		fmt.Println("warn: the response is resolve => please use return on Next and res.send function call")
		return
	}
	if len(r.handler) > 0 {
		curFunc := r.handler[0]
		curIndex := 0
		var next func()
		next = func() {
			if res.resolved {
				fmt.Println("warn: the response is resolve => please use return on Next and res.send function call")
				return
			}
			curIndex++
			if curIndex < len(r.handler) {
				nextFunc := r.handler[curIndex]
				req.Next = &next
				nextFunc.Resolve(req, res)
			} else {
				r.Resolve(req, res)
			}
		}
		curFunc.Resolve(req, res)
	}
}

func (r *Router) Resolve(req *Request, res *Response) {
	if res.resolved {
		fmt.Println("warn: the response is resolve => please use return on Next and res.send function call")
		return
	}
	r.handleMiddleware(req, res)
	if res.resolved {
		return
	}
}

func (r *Router) Use(handler ...Resolver) {
	r.handler = append(r.handler, handler...)
}

func (r *Router) UsePath(path string, handler ...Resolver) {
	route := Route(path, RequestTypeAny, handler...)
	r.add(route)
}

func (r *Router) UseMiddleWare(handler ...Handler) {
	r.add(convertToResolver(handler...)...)
}

//
// handle the request with request handler

func (r *Router) Get(path string, handler ...Handler) {
	route := Route(path, RequestTypeGet, convertToResolver(handler...)...)
	r.add(route)
}

func (r *Router) Post(path string, handler ...Handler) {
	route := Route(path, RequestTypePost, convertToResolver(handler...)...)
	r.add(route)
}

func (r *Router) Delete(path string, handler ...Handler) {
	route := Route(path, RequestTypeDelete, convertToResolver(handler...)...)
	r.add(route)
}

func (r *Router) Put(path string, handler ...Handler) {
	route := Route(path, RequestTypePut, convertToResolver(handler...)...)
	r.add(route)
}

func (r *Router) Patch(path string, handler ...Handler) {
	route := Route(path, RequestTypePatch, convertToResolver(handler...)...)
	r.add(route)
}

func (r *Router) Head(path string, handler ...Handler) {
	route := Route(path, RequestTypeHead, convertToResolver(handler...)...)
	r.add(route)
}

func (r *Router) Options(path string, handler ...Handler) {
	route := Route(path, RequestTypeOptions, convertToResolver(handler...)...)
	r.add(route)
}

func (r *Router) Connect(path string, handler ...Handler) {
	route := Route(path, RequestTypeConnect, convertToResolver(handler...)...)
	r.add(route)
}

func (r *Router) Trace(path string, handler ...Handler) {
	route := Route(path, RequestTypeTrace, convertToResolver(handler...)...)
	r.add(route)
}

func (r *Router) Any(path string, handler ...Handler) {
	route := Route(path, RequestTypeAny, convertToResolver(handler...)...)
	r.add(route)
}
