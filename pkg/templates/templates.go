package templates

import (
	"html/template"
	"sync"
)

type ITemplates interface {
	Compile(params CompileParams) ([]byte, error)
	AddFunc(key string, value any)
}

type templates struct {
	mu    sync.Mutex
	Funcs template.FuncMap
}

func New() ITemplates {
	return &templates{
		Funcs: DefaultTemplateFuncs,
	}
}
