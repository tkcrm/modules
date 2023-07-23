package templates

import (
	"bufio"
	"bytes"
	"fmt"

	htmlTemplate "html/template"
	textTemplate "text/template"
)

var (
	errExecuteTemplateStr = "execute template error: %w"
	defaultTplName        = "name"
)

func (t *templates) Compile(params CompileParams) ([]byte, error) {
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
			return nil, fmt.Errorf(errExecuteTemplateStr, err)
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
			return nil, fmt.Errorf(errExecuteTemplateStr, err)
		}
	}

	w.Flush()

	return b.Bytes(), nil
}
