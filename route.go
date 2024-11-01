package te

import "regexp"

type RouteUnit struct {
	path     *regexp.Regexp
	resolver []Resolver
	method   RequestMethod
}

func Route(path string, method RequestMethod, resolvers ...Resolver) *RouteUnit {
	pattern := makeRegExpPattern(path)
	return &RouteUnit{
		path:     pattern,
		resolver: resolvers,
		method:   method,
	}
}

func (r *RouteUnit) AddResolver(resolver Resolver) {
	r.resolver = append(r.resolver, resolver)
}

func (r *RouteUnit) Resolve(req *Request, res *Response) {

	//handle middleware
	if res.resolved {
		return
	}
	if len(r.resolver) > 0 {
		curFunc := r.resolver[0]
		curIndex := 0
		var next func()
		next = func() {
			if res.resolved {
				return
			}
			curIndex++
			if curIndex < len(r.resolver) {
				nextFunc := r.resolver[curIndex]
				req.Next = &next
				nextFunc.Resolve(req, res)
			}
		}
		req.Next = &next
		curFunc.Resolve(req, res)
	}

}

// return if the path is matched with the route path, the matched part, the rest of the path, and the parameters
func (r *RouteUnit) match(path string) (bool, string, string, map[string]string) {

	return extractParamsFromStr(r.path, path)
}
