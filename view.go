package main

import (
	"html/template"

	"github.com/eknkc/amber"
	"github.com/gin-gonic/gin/render"
)

// NewAmberRenderer ...
func NewAmberRenderer(dir string, ext string, funcs *template.FuncMap) *AmberRender {
	options := amber.DirOptions{}
	options.Recursive = true
	options.Ext = ext

	if funcs != nil {
		tmplFuncs := *funcs
		for k, v := range amber.FuncMap {
			tmplFuncs[k] = v
		}
		amber.FuncMap = tmplFuncs
	}

	templates, err := amber.CompileDir(dir, options, amber.DefaultOptions)
	if err != nil {
		panic(err)
	}

	return &AmberRender{
		templates: templates,
	}
}

// AmberRender ...
type AmberRender struct {
	templates map[string]*template.Template
}

// Add ...
func (r *AmberRender) Add(name string, tmpl *template.Template) {
	if r.templates == nil {
		r.templates = make(map[string]*template.Template)
	}
	r.templates[name] = tmpl
}

// Instance ...
func (r *AmberRender) Instance(name string, data interface{}) render.Render {
	return render.HTML{
		Template: r.templates[name],
		Data:     data,
	}
}
