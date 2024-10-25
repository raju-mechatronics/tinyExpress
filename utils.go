package te

import "regexp"

func extractParamsFromRoute(re *regexp.Regexp, path string) (string, map[string]string) {
	var params map[string]string = nil
	matches := re.FindStringSubmatch(path)
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
