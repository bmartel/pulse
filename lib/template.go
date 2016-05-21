package lib

import (
	"html/template"
)

// TemplateFunc ...
type TemplateFunc func() (string, error)

// DefineTemplateFuncs ...
func DefineTemplateFuncs(define func(*Template)) []template.FuncMap {
	tmplFunc := &Template{
		make([]template.FuncMap, 0),
	}
	define(tmplFunc)

	return tmplFunc.funcs
}

// Template ...
type Template struct {
	funcs []template.FuncMap
}

// Add ...
func (t *Template) Add(name string, fn TemplateFunc) {
	t.funcs = append(t.funcs, template.FuncMap{name: fn})
}
