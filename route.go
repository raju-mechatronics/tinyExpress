package te

import (
	"regexp"
)

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

func (r *RouteUnit) addResolver(resolver Resolver) {
	r.resolver = append(r.resolver, resolver)
}

func (r *RouteUnit) Resolve(req *Request, res *Response) {

	if res.resolved {
		return
	}

	if r.method != RequestTypeAny && r.method != req.Method {
		return
	}

	matched, matchedPart, rest, params := extractParamsFromStr(r.path, req.CurrentPath)

	if !matched {
		return
	}

	//merge the params
	if req.Params == nil {
		req.Params = make(map[string]string)
	}
	for k, v := range params {
		req.Params[k] = v
	}

	req.CurrentPath = matchedPart
	req.NextPath = rest

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
