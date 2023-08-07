package utils

import (
	"container/list"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string, data list.List) {
	tmplPath := "templates/" + tmpl + ".html"
	t, err := template.ParseFiles("templates/base.html", tmplPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, data)
	if err != nil {
		return
	}
}
