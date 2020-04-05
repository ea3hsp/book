package render

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/oxtoacart/bpool"
)

var (
	mainTmpl  = `{{define "main" }} {{ template "base" . }} {{ end }}`
	bufpool   *bpool.BufferPool
	templates map[string]*template.Template
)

// Render renderizes html templates
type Render struct {
	templateLayoutPath  string
	templateIncludePath string
}

// New creates new render
func New(layoutPath, includePath string) *Render {
	bufpool = bpool.NewBufferPool(64)
	return &Render{
		templateLayoutPath:  layoutPath,
		templateIncludePath: includePath,
	}
}

// LoadTemplates loads html templating files
func (r *Render) LoadTemplates() (err error) {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}
	// load layout templates
	layoutFiles, err := filepath.Glob(r.templateLayoutPath + "*.html")
	if err != nil {
		return err
	}
	// load include templates
	includeFiles, err := filepath.Glob(r.templateIncludePath + "*.html")
	if err != nil {
		return err
	}
	// main template generator
	mainTemplate := template.New("main")
	// parsing main template
	mainTemplate, err = mainTemplate.Parse(mainTmpl)
	if err != nil {
		return err
	}
	for _, file := range includeFiles {
		fileName := filepath.Base(file)
		files := append(layoutFiles, file)
		templates[fileName], err = mainTemplate.Clone()
		if err != nil {
			return err
		}
		templates[fileName] = template.Must(templates[fileName].ParseFiles(files...))
	}
	return nil
}

// RenderTemplate renders template
func (r *Render) RenderTemplate(w http.ResponseWriter, name string, data interface{}) error {
	tmpl, ok := templates[name]
	if !ok {
		http.Error(w, fmt.Sprintf("The template %s does not exist.", name),
			http.StatusInternalServerError)
		err := errors.New("Template doesn't exist")
		return err
	}
	buf := bufpool.Get()
	defer bufpool.Put(buf)

	err := tmpl.Execute(buf, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		err := errors.New("Template execution failed")
		return err
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	buf.WriteTo(w)
	return nil
}
