// Package tmpl contains the logic for creating and parsing html templates
package tmpl

import (
	"html/template"
	"io"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

// tmpls struct for parsed html templates
type tmpls struct {
	templates *template.Template
}

// Render renders html templates
func (t *tmpls) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// newTemplate creates a new collection of parsed templates
func CreateTemplate() *tmpls {
	tmpl := template.New("")

	// walk through all html files in the view folder
	err := filepath.Walk("views", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// skip directories
		if info.IsDir() {
			return nil
		}

		// parse html files
		if filepath.Ext(path) == ".html" {
			_, err = tmpl.ParseFiles(path)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	return &tmpls{
		templates: tmpl,
	}
}
