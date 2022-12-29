package templates

import (
	"html/template"
	"sync"
)

type ITemplates interface {
	Compile(params CompileParams) ([]byte, error)
}

type Templates struct {
	mu    sync.Mutex
	Funcs template.FuncMap
}

func New() ITemplates {
	return &Templates{
		Funcs: defaultTemplateFuncs,
	}
}
