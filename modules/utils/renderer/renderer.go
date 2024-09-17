package renderer

import (
	"io"
	"text/template"

	"github.com/labstack/echo/v4"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	println(t.templates.DefinedTemplates())
	return t.templates.ExecuteTemplate(w, name, data)
}

func Renderer() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("modules/templates/**/*.html")),
	}
}
