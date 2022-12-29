package templates_test

import (
	"embed"
	"fmt"
	"testing"
	"time"

	"github.com/tkcrm/modules/pkg/templates"
)

//go:embed testtpls/*.go.tmpl
var tplsFS embed.FS

func Test_Compile(t *testing.T) {
	tpl := templates.New()

	res, err := tpl.Compile(templates.CompileParams{
		TemplateName: "testTextTpl",
		TemplateType: templates.TextTemplateType,
		FS:           tplsFS,
		FSPaths:      []string{"testtpls/test.go.tmpl"},
		Data: struct {
			Date    time.Time
			Title   string
			Version float32
		}{
			Date:    time.Now(),
			Title:   "Some perfect title",
			Version: 0.1,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	res2, err := tpl.Compile(templates.CompileParams{
		TemplateName: "testHtmlTpl",
		TemplateType: templates.TextTemplateType,
		FS:           tplsFS,
		FSPaths:      []string{"testtpls/test.go.tmpl"},
		Data: struct {
			Date    time.Time
			Title   string
			Version float32
		}{
			Date:    time.Now(),
			Title:   "Some perfect title",
			Version: 0.1,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(res))
	fmt.Println(string(res2))
}
