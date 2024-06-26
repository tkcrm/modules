package templates

import (
	"fmt"
	"html/template"
	"strings"
	"time"

	"github.com/gobeam/stringy"
)

func (t *templates) AddFunc(key string, value any) {
	t.mu.Lock()
	t.Funcs[key] = value
	t.mu.Unlock()
}

var DefaultTemplateFuncs = template.FuncMap{
	"lower": strings.ToLower,
	"snakeCase": func(str string) string {
		return stringy.New(str).SnakeCase().ToLower()
	},
	"camelCase": func(str string) string {
		return stringy.New(str).CamelCase().Get()
	},
	"lowerCaseFirstLetter": func(str string) string {
		return stringy.New(str).LcFirst()
	},
	"upperCaseFirstLetter": func(str string) string {
		return stringy.New(str).UcFirst()
	},
	"join": strings.Join,
	"dateFormat": func(t time.Time, layout string) string {
		return t.Format(layout)
	},
	"increment": func(v int) int {
		return v + 1
	},
	"roundFloat": func(count string, number float64) string {
		return fmt.Sprintf("%."+count+"f", number)
	},
}
