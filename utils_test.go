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
		match, params := extractParamsFromStr(re, p)
		fmt.Println(match, params)
	}
}

func TestMakeRegExpPattern(t *testing.T) {
	path := []string{
		"/",
		"/p1/:id{int}/",
		"/p2/:id{string}/",
		"/p3/:id{float}/",
		"/p5/:id{int}/p6/:name{string}/",
		"/p7/:id{int}/p8/:name{string}/p9/:age{float}/",
		//"/p4/{regexp}/",
	}

	regExpList := []string{
		"^/",
		`^(p1)/(?P<id>\d+)`,
		`^(p2)/(?P<id>.*)`,
		`^(p3)/(?P<id>\d+\.\d+)`,
		`^(p5)/(?P<id>\d+)/(p6)/(?P<name>.*)`,
		`^(p7)/(?P<id>\d+)/(p8)/(?P<name>.*)/(p9)/(?P<age>\d+\.\d+)`,
		//`^/p4/(regexp)/`,
	}

	for i, p := range path {
		re := makeRegExpPattern(p)
		fmt.Println(re.String())
		fmt.Println(regExpList[i])
		if re.String() != regExpList[i] {
			t.Error("Error")
		}
	}
}
