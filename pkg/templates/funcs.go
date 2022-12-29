package templates

import (
	"html/template"
	"strings"
	"time"

	"github.com/gobeam/stringy"
)

func (t *Templates) AddFunc(key string, value any) {
	t.mu.Lock()
	t.Funcs[key] = value
	t.mu.Unlock()
}

var defaultTemplateFuncs = template.FuncMap{
	"lower": strings.ToLower,
	"snake_case": func(str string) string {
		return stringy.New(str).SnakeCase().ToLower()
	},
	"camel_case": func(str string) string {
		return stringy.New(str).CamelCase()
	},
	"lcfirst": func(str string) string {
		if str == "ID" {
			return "id"
		}
		return stringy.New(str).LcFirst()
	},
	"join": strings.Join,
	"date_format": func(t time.Time, layout string) string {
		return t.Format(layout)
	},
	"increment": func(v int) int {
		return v + 1
	},
}
