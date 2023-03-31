package auth

import (
	`github.com/gookit/goutil/arrutil`
	`github.com/kataras/iris/v12/context`
	`os`
	`strings`
)

type List struct {
	Name    string `json:"name"`
	Path    string `json:"path"`
	Allowed bool   `json:"allowed"`
}

type Acl struct {
	// Lists 访问列表
	Lists map[string]List
	// Exceptional 不检查权限
	Exceptional []string `json:"exceptional"`
}

func route(value string) string {
	return strings.ToLower(strings.Trim(value, string(os.PathSeparator)))
}

func paths(values []string) (data []string) {
	for _, value := range values {
		data = append(data, route(value))
	}
	return
}

func NewAcl(routers []context.RouteReadOnly, exceptional []string, rules []string) (a Acl) {
	exceptional = paths(exceptional)
	rules = paths(rules)
	var lists = make(map[string]List)
	for _, router := range routers {
		var url = route(router.Path())
		lists[url] = List{
			Name:    router.Name(),
			Path:    url,
			Allowed: arrutil.StringsHas(rules, url),
		}
	}
	a = Acl{
		Lists:       lists,
		Exceptional: exceptional,
	}
	return
}

func (a Acl) Checker(url string) bool {
	url = route(url)
	rule, ok := a.Lists[url]
	if ok {
		return rule.Allowed || arrutil.StringsHas(a.Exceptional, url)
	}
	return ok
}
