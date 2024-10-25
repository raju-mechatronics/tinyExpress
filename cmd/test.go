package main

import (
	"fmt"
	"regexp"
)

func main() {
	path := []string{
		"/",
		"/hello",
		"/hello/name",
		"/hello/name/18",
		"/hello/name/age/any",
		"/hello/hi",
		"/hello/hi/18",
		"/hello/hi/18/abc",
		"/hello/hi/18/abc/def",
	}

	re := regexp.MustCompile(`/hello/name/(?P<age>\d+)`)

	for _, p := range path {
		if re.MatchString(p) {
			fmt.Println(re.FindStringSubmatch(p))
			fmt.Println(re.SubexpNames())
		}
	}
}
