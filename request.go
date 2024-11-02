package te

import (
	"io"
	"net/http"
	"net/url"
)

type Request struct {
	*http.Request

	//Original URL string
	OriginalURL string

	// full Path
	Path string

	// Next path contains the rest of the path that is not matched.
	NextPath string

	//Current Path That is matched
	CurrentPath string

	// Body data.
	Body []byte

	// Parameters from the matched route
	Params map[string]string

	// Query string parameters
	Query url.Values

	// Cookies sent with the request
	Cookies []*http.Cookie

	// for middle ware support
	Next               *func()
	applicationContext map[string]interface{}
}

// NewRequest initializes a new Request object
func NewRequest(r *http.Request) (*Request, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	cookies := r.Cookies()

	req := &Request{
		Request:            r,
		OriginalURL:        r.URL.String(),
		Path:               r.URL.Path,
		NextPath:           r.URL.Path,
		CurrentPath:        "",
		Body:               body,
		Params:             nil,
		Query:              r.URL.Query(),
		Cookies:            cookies,
		Next:               nil,
		applicationContext: nil,
	}

	return req, nil

}

// GetParam retrieves a route parameter by name
func (req *Request) GetParam(name string) string {
	return req.Params[name]
}

// GetQuery retrieves a query string parameter by name
func (req *Request) GetQuery(name string) string {
	return req.Query.Get(name)
}

// GetCookie retrieves a cookie by name
func (req *Request) GetCookie(name string) (*http.Cookie, error) {
	for _, cookie := range req.Cookies {
		if cookie.Name == name {
			return cookie, nil
		}
	}
	return nil, http.ErrNoCookie
}

func (req *Request) GetBody() interface{} {
	return req.Body
}
