package te

import (
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Request struct {
	*http.Request
	Method       string
	URL          *url.URL
	OriginalURL  string
	Path         string
	CurrentPath  string
	OriginalPath string
	Header       http.Header
	Body         []byte
	RemoteAddr   string
	Params       map[string]string
	Query        url.Values
	Cookies      []*http.Cookie
	Session      map[string]interface{}
	Host         string
	IP           string
	Protocol     string
	Secure       bool
	Next         *func()
}

// NewRequest initializes a new Request object
func NewRequest(r *http.Request) (*Request, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	cookies := r.Cookies()
	protocol := "http"
	if r.TLS != nil {
		protocol = "https"
	}

	return &Request{
		Method:     r.Method,
		URL:        r.URL,
		Header:     r.Header,
		Body:       body,
		RemoteAddr: r.RemoteAddr,
		Params:     make(map[string]string), // This should be populated based on your routing logic
		Query:      r.URL.Query(),
		Cookies:    cookies,
		Session:    make(map[string]interface{}), // This should be populated based on your session management logic
		Host:       r.Host,
		IP:         strings.Split(r.RemoteAddr, ":")[0],
		Protocol:   protocol,
		Secure:     r.TLS != nil,
	}, nil
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

// GetSession retrieves a session value by key
func (req *Request) GetSession(key string) interface{} {
	return req.Session[key]
}

func (req *Request) GetBody() interface{} {
	return req.Body
}
