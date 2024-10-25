package te

/*
	ParamPath: /user/:id{int}/
	ParamPath: /user/:id{string}
	ParamPath: /user/:id{float}
	ParamPath: /user/{regexp}
	ParamPath: /user/:id{url}
*/

type PathHandler struct {
	normalPath map[string]Resolver
	paramPath  map[string]map[string]Resolver
}

/*
func (ph PathHandler) Add(path string, resolver Resolver) error {
	if !strings.Contains(path, ":") || !strings.Contains(path, "{") {
		if ph.normalPath == nil {
			ph.normalPath = make(map[string]Resolver)
			ph.normalPath[path] = resolver
		} else {
			if _, ok := ph.normalPath[path]; ok {
				return fmt.Errorf("path %s already exist", path)
			} else {
				ph.normalPath[path] = resolver
			}
		}
	} else {
		if ph.paramPath == nil {
			ph.paramPath = make(map[string]map[string]Resolver)
			parts := strings.Split(path, "/")
			for index, path := range parts {
				if strings.Contains(path, ":") {
					var param, paramType string
					openCurl := strings.Index(path, "{")
					closeCurl := strings.Index(path, "}")
					if openCurl != -1 && closeCurl != -1 {
						paramType = path[openCurl+1 : closeCurl]
					} else {
						paramType = "string"
					}
					param = path[1:openCurl]
				}
			}
		}

	}
}

func (ph PathHandler) Get(path string) Resolver {
	if resolver, ok := ph.normalPath[path]; ok {
		return resolver
	}

	//

	return nil
}

type Router struct {
	handler  []Resolver
	routeMap map[RequestType]PathHandler
}

func (r Router) Resolve(req *Request, res *Response) {

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

func (r *Router) Use(path string, handler ...Resolver) {

}

func (r *Router) Get(path string, handler ...Handler) {

}

func (r *Router) Post(path string, handler ...Handler) {

}

func (r *Router) Delete(path string, handler ...Handler) {

}

func (r *Router) Put(path string, handler ...Handler) {

}
func (r *Router) Patch(path string, handler ...Handler) {

}


*/
