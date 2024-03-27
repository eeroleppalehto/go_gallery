package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
	tmpl *template.Template
}

func newTemplate() *Template {
	return &Template{
		tmpl: template.Must(template.ParseGlob("views/*.html")),
	}
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.tmpl.ExecuteTemplate(w, name, data)
}

type Page struct {
	Data string
}

func main() {

	e := echo.New()

	e.Static("/static", "static")

	e.Renderer = newTemplate()
	e.Use(middleware.Logger())

	page := Page{
		Data: "mooi",
	}

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", page)
	})

	e.Logger.Fatal(e.Start(":8081"))
}
