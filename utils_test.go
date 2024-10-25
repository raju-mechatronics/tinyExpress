package te

import (
	"fmt"
	"regexp"
	"testing"
)

func TestExtractParamsFromRoute(t *testing.T) {
	path := []string{
		"/",
		"/hello",
		"/hello/na-me",
		"/hello/name/18",
		"/hello/name/age/any/many",
	}

	regExpList := []string{
		"^/",
		"^/hello",
		`^/hello/(?P<name>\w+)`,
		`^/hello/(?P<name>\w+)/(?P<age>\d+)`,
		`^/hello/(?P<name>\w+)/(?P<age>\w+)`,
	}

	for i, p := range path {
		re := regexp.MustCompile(regExpList[i])
		fmt.Println(re.FindStringSubmatch(p))
		match, params := extractParamsFromRoute(re, p)
		fmt.Println(match, params)
	}
}
