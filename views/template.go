package views

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"log"
	"net/http"

	"github.com/gorilla/csrf"
)

func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return t
}

func ParseFS(fs fs.FS, patterns ...string) (Template, error) {
	tpl := template.New(patterns[0])
	tpl = tpl.Funcs(
		template.FuncMap{
			"csrfField": func() (template.HTML, error) {
				return "", fmt.Errorf("csrf field not implemented")
			},
		},
	)
	tpl, err := tpl.ParseFS(fs, patterns...)
	if err != nil {
		return Template{}, fmt.Errorf("parseFS template: %w", err)
	}
	return Template{
		htmlTpl: tpl,
	}, nil
}

// func Parse(filepath string) (Template, error) {
// 	//----- OS agnostic way, this to set / or \ correctly -----
// 	// tplPath := filepath.Join("templates", "home.gohtml")
// 	// tpl, err := template.ParseFiles(tplPath)
// 	tpl, err := template.ParseFiles(filepath)
// 	if err != nil {
// 		return Template{}, fmt.Errorf("parsing template: %w", err)
// 	}
// 	return Template{
// 		htmlTpl: tpl,
// 	}, nil
// }

type Template struct {
	htmlTpl *template.Template
}

func (t Template) Execute(w http.ResponseWriter, r *http.Request, data interface{}) {
	tpl, err := t.htmlTpl.Clone()
	if err != nil {
		log.Printf("cloning tempalte: %v", err)
		http.Error(w, "There was an error rendering the page:", http.StatusInternalServerError)
	}
	tpl = tpl.Funcs(
		template.FuncMap{
			"csrfField": func() template.HTML {
				return csrf.TemplateField(r)
			},
		},
	)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var buf bytes.Buffer
	err = tpl.Execute(&buf, data)
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "There was an error executing the template.", http.StatusInternalServerError)
		return
	}
	io.Copy(w, &buf)
}
