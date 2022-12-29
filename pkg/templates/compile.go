package templates

import (
	"bufio"
	"bytes"
	htmlTemplate "html/template"
	textTemplate "text/template"

	"github.com/pkg/errors"
)

var (
	errExecuteTemplateStr = "execute template error"

	defaultTplName = "name"
)

func (t *Templates) Compile(params CompileParams) ([]byte, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	b := bytes.NewBuffer([]byte{})
	w := bufio.NewWriter(b)

	if params.TemplateType == TextTemplateType {
		tpl := textTemplate.Must(
			textTemplate.New(defaultTplName).
				Funcs(t.Funcs).
				ParseFS(
					params.FS,
					params.FSPaths...,
				),
		)

		if err := tpl.ExecuteTemplate(w, params.TemplateName, params.Data); err != nil {
			return nil, errors.Wrap(err, errExecuteTemplateStr)
		}
	} else {
		tpl := htmlTemplate.Must(
			htmlTemplate.New(defaultTplName).
				Funcs(t.Funcs).
				ParseFS(
					params.FS,
					params.FSPaths...,
				),
		)

		if err := tpl.ExecuteTemplate(w, params.TemplateName, params.Data); err != nil {
			return nil, errors.Wrap(err, errExecuteTemplateStr)
		}
	}

	w.Flush()

	return b.Bytes(), nil
}
