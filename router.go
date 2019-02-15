package goui

import (
	"regexp"
)

type route struct {
	handler func(*Context)
	paras   []string       //named parameters
	regex   *regexp.Regexp //if there are parameters
}

var (
	//optionalPart = regexp.MustCompile(`\((.*?)\)`)
	namedPart = regexp.MustCompile(`:([^/]+)`)
	//splatPart    = regexp.MustCompile(`\*\w+`)
	escapeRegExp = regexp.MustCompile(`([\-{}\[\]+?.,\\\^$|#\s])`)

	routes = make(map[string]*route)
)

func parseRoute(pattern string, route *route) {
	params := namedPart.FindAllStringSubmatch(pattern, -1)
	if params != nil {
		l := len(params)
		route.paras = make([]string, l)

		//fmt.Println(params)
		for i, param := range params {
			route.paras[i] = param[1]
			//fmt.Println(param[1])
		}
		pattern = namedPart.ReplaceAllString(pattern, `([^/]+)`)
		route.regex = regexp.MustCompile(pattern)
	}
	routes[pattern] = route
}

func dispatch(url string) (handler func(*Context), params map[string]string) {
	for key, route := range routes {
		if route.regex == nil {
			if key == url {
				handler = route.handler
				return
			}

		} else {
			matches := route.regex.FindAllStringSubmatch(url, -1)
			if matches != nil && len(matches) == 1 {
				params = make(map[string]string)
				vals := matches[0][1:]
				l, lkey := len(vals), len(route.paras)
				if l > lkey {
					l = lkey
				}
				for i := 0; i < l; i++ {
					params[route.paras[i]] = vals[i]
				}
				handler = route.handler
				return
			}
		}
	}
	return
}
