package main

import (
	"embed"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/tashima42/textboard/handler"
)

//go:embed public/*
var templatesFS embed.FS

type Template struct {
  templates *template.Template
}

func main() {
  tmpls, err := template.New("").ParseFS( templatesFS, "public/templates/*.tmpl")
  if err != nil {
    log.Fatal(err)
  }

  echo.NotFoundHandler = notFoundHandler
  e := echo.New()
  e.Filesystem = templatesFS
  e.Renderer = NewTemplate(tmpls)
  e.HTTPErrorHandler = customHTTPErrorHandler

  e.GET("/:board", handler.BoardHandler)
  e.GET("/:board/:ref", handler.PostHandler)

  e.Logger.Fatal(e.Start(":3000"))
}

func NewTemplate(templates *template.Template) *Template {
  t := Template{ templates }
  return &t
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	c.Logger().Error(err)
  c.Response().Status = code
	if err := c.Render(code, strconv.Itoa(code), nil); err != nil {
		c.Logger().Error(err)
	}
}

func notFoundHandler(c echo.Context) error {
  return echo.NewHTTPError(http.StatusNotFound)
}

