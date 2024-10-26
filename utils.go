package te

import (
	"regexp"
	"strings"
)

func extractParamsFromStr(re *regexp.Regexp, str string) (string, map[string]string) {
	var params map[string]string = nil
	matches := re.FindStringSubmatch(str)
	names := re.SubexpNames()
	for i, name := range names {
		if i != 0 && name != "" {
			if params == nil {
				params = make(map[string]string)
			}
			params[name] = matches[i]
		}
	}
	return matches[0], params
}

/*
a function that takes the path and returns the regular expression pattern

Example:

	/user/:id{int}/ => ^/user/(?P<id>\d+)/
	/user/:id{string} => ^/user/(?P<id>\w+)/
	/user/:id{float} => ^/user/(?P<id>\d+\.\d+)/
	/user/{regexp} => ^/user/(regexp)/ [Not yet completed]
*/
func makeRegExpPattern(path string) *regexp.Regexp {
	parts := strings.Split(path, "/")

	stringRegExp := `.*`
	intRegExp := `\d+`
	floatRegExp := `\d+\.\d+`

	regExpString := `^`

	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		if part[0] != ':' {
			regExpString += "(" + part + ")" + "/"
		} else if part[0] == ':' {
			paramType := ""
			paramName := ""
			leftBraceIndex := strings.Index(part, "{")
			rightBraceIndex := strings.Index(part, "}")
			if leftBraceIndex == -1 || rightBraceIndex == -1 {
				paramType = "string"
			} else {
				paramType = part[leftBraceIndex+1 : rightBraceIndex]
			}

			if leftBraceIndex != -1 {
				paramName = part[1:leftBraceIndex]
			} else {
				paramName = part[1:]
			}

			if paramType == "int" {
				regExpString += "(?P<" + paramName + ">" + intRegExp + ")/"
			} else if paramType == "float" {
				regExpString += "(?P<" + paramName + ">" + floatRegExp + ")/"
			} else if paramType == "string" {
				regExpString += "(?P<" + paramName + ">" + stringRegExp + ")/"
			} else if paramType == "regexp" {
				regExpString += "(" + paramName + ")/"
			}
		}
	}

	regExpString = strings.TrimSuffix(regExpString, "/")
	rexp := regexp.MustCompile(regExpString)
	return rexp
}
