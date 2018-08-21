package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"runtime/debug"
)

type ListInt interface {
	GetInfo() string
}

func (li List) GetInfo() string {
	return li.Subject
}

func (msg Message) GetInfo() string {
	return msg.Subject
}

func (app *App) RenderForm(w http.ResponseWriter, r *http.Request, tmpl string, li ListInt) {

	pli := &li //pointer to a instance of interface
	files := []string{
		filepath.Join(app.HTMLDir, tmpl+".html"),
		//app.HTMLDir + "/" + tmpl + ".html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Printf("%s\n%s", err.Error(), debug.Stack())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, tmpl, pli)

	if err != nil {
		log.Printf("%s\n%s", err.Error(), debug.Stack())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
