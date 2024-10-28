package te

import (
	"fmt"
)

type Router struct {
	routes  []*RouteUnit
	handler []Resolver
}

func (r *Router) addRoute(route *RouteUnit) {
	r.routes = append(r.routes, route)
}

func (r *Router) addHandler(handler Resolver) {
	r.handler = append(r.handler, handler)
}

func (r *Router) Resolve(req *Request, res *Response) {
	if res.resolved {
		fmt.Println("warn: the response is resolve => please use return on next and res.send function call")
		return
	}
	r.handleMiddleware(req, res)
	if res.resolved {
		return
	}

	currentPath := req.Path
	for _, route := range r.routes {
		matched, matchedPath, restPath, params := route.match(currentPath)
		if !matched {
			continue
		} else {
			if req.Params == nil {
				req.Params = make(map[string]string)
			}
			for k, v := range params {
				req.Params[k] = v
			}
			req.CurrentPath = matchedPath
			req.Path = restPath
		}

	}
}

func (r *Router) handleMiddleware(req *Request, res *Response) {
	if res.resolved {
		fmt.Println("warn: the response is resolve => please use return on next and res.send function call")
		return
	}
	if len(r.handler) > 0 {
		curFunc := r.handler[0]
		curIndex := 0
		var next func()
		next = func() {
			if res.resolved {
				fmt.Println("warn: the response is resolve => please use return on next and res.send function call")
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

func (r *Router) Use(handler ...Resolver) {
	r.handler = append(r.handler, handler...)
}

func (r *Router) UsePath(path string, handler ...Resolver) {
	route := Route(path, RequestTypeAny, handler...)
	r.addRoute(route)
}

func (r *Router) Get(path string, handler ...Handler) {
	//convert handler to resolver
	resolvers := make([]Resolver, len(handler))
	for i, h := range handler {
		resolvers[i] = h
	}

	route := Route(path, RequestTypeGet, resolvers...)
	r.addRoute(route)
}

func (r *Router) Post(path string, handler ...Handler) {
	resolvers := make([]Resolver, len(handler))
	for i, h := range handler {
		resolvers[i] = h
	}

	route := Route(path, RequestTypePost, resolvers...)
	r.addRoute(route)
}

func (r *Router) Delete(path string, handler ...Handler) {
	resolvers := make([]Resolver, len(handler))
	for i, h := range handler {
		resolvers[i] = h
	}

	route := Route(path, RequestTypeDelete, resolvers...)
	r.addRoute(route)

}

func (r *Router) Put(path string, handler ...Handler) {
	resolvers := make([]Resolver, len(handler))
	for i, h := range handler {
		resolvers[i] = h
	}

	route := Route(path, RequestTypePut, resolvers...)
	r.addRoute(route)
}
func (r *Router) Patch(path string, handler ...Handler) {
	resolvers := make([]Resolver, len(handler))
	for i, h := range handler {
		resolvers[i] = h
	}

	route := Route(path, RequestTypePatch, resolvers...)
	r.addRoute(route)
}

func (r *Router) Head(path string, handler ...Handler) {
	resolvers := make([]Resolver, len(handler))
	for i, h := range handler {
		resolvers[i] = h
	}

	route := Route(path, RequestTypeHead, resolvers...)
	r.addRoute(route)
}

func (r *Router) Options(path string, handler ...Handler) {
	resolvers := make([]Resolver, len(handler))
	for i, h := range handler {
		resolvers[i] = h
	}

	route := Route(path, RequestTypeOptions, resolvers...)
	r.addRoute(route)
}

func (r *Router) Connect(path string, handler ...Handler) {
	resolvers := make([]Resolver, len(handler))
	for i, h := range handler {
		resolvers[i] = h
	}

	route := Route(path, RequestTypeConnect, resolvers...)
	r.addRoute(route)
}

func (r *Router) Trace(path string, handler ...Handler) {
	resolvers := make([]Resolver, len(handler))
	for i, h := range handler {
		resolvers[i] = h
	}

	route := Route(path, RequestTypeTrace, resolvers...)
	r.addRoute(route)
}

func (r *Router) Any(path string, handler ...Handler) {
	resolvers := make([]Resolver, len(handler))
	for i, h := range handler {
		resolvers[i] = h
	}

	route := Route(path, RequestTypeAny, resolvers...)
	r.addRoute(route)
}
