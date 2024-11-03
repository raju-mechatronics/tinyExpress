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

func (r *Router) Resolve(req *Request, res *Response) {
	if res.resolved {
		fmt.Println("warn: the response is resolve => please use return on Next and res.send function call")
		return
	}

	for _, handler := range r.handler {
		if res.resolved {
			return
		}
		handler.Resolve(req, res)
	}
}

func (r *Router) Use(handler ...Resolver) {
	r.handler = append(r.handler, handler...)
}

func (r *Router) UsePath(path string, handler ...Resolver) {
	route := Route(path, RequestMethodAny, handler...)
	r.add(route)
}

func (r *Router) UseMiddleWare(handler ...Handler) {
	r.add(convertToResolver(handler...)...)
}

//
// handle the request with request handler

func (r *Router) Get(path string, handler ...Handler) {
	route := Route(path, RequestMethodGet, convertToResolver(handler...)...)
	r.add(route)
}

func (r *Router) Post(path string, handler ...Handler) {
	route := Route(path, RequestMethodPost, convertToResolver(handler...)...)
	r.add(route)
}

func (r *Router) Delete(path string, handler ...Handler) {
	route := Route(path, RequestMethodDelete, convertToResolver(handler...)...)
	r.add(route)
}

func (r *Router) Put(path string, handler ...Handler) {
	route := Route(path, RequestMethodPut, convertToResolver(handler...)...)
	r.add(route)
}

func (r *Router) Patch(path string, handler ...Handler) {
	route := Route(path, RequestMethodPatch, convertToResolver(handler...)...)
	r.add(route)
}

func (r *Router) Head(path string, handler ...Handler) {
	route := Route(path, RequestMethodHead, convertToResolver(handler...)...)
	r.add(route)
}

func (r *Router) Options(path string, handler ...Handler) {
	route := Route(path, RequestMethodOptions, convertToResolver(handler...)...)
	r.add(route)
}

func (r *Router) Connect(path string, handler ...Handler) {
	route := Route(path, RequestMethodConnect, convertToResolver(handler...)...)
	r.add(route)
}

func (r *Router) Trace(path string, handler ...Handler) {
	route := Route(path, RequestMethodTrace, convertToResolver(handler...)...)
	r.add(route)
}

func (r *Router) Any(path string, handler ...Handler) {
	route := Route(path, RequestMethodAny, convertToResolver(handler...)...)
	r.add(route)
}
