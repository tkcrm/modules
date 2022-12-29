package templates

import (
	"io/fs"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type TemplateType int

const (
	TextTemplateType TemplateType = iota + 1
	HtmlTemplateType
)

type CompileParams struct {
	TemplateName string
	TemplateType TemplateType
	FS           fs.FS
	FSPaths      []string
	Data         any
}

func (s *CompileParams) Validate() error {
	return validation.ValidateStruct(
		s,
		validation.Field(&s.TemplateName, validation.Required),
		validation.Field(&s.TemplateType, validation.Required),
		validation.Field(&s.FS, validation.Required),
		validation.Field(&s.FSPaths, validation.Required),
		validation.Field(&s.Data, validation.Required),
	)
}
