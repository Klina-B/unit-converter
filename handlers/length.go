package handlers

import (
	"html/template"
	"net/http"
)

func LengthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet{
		tmpl := template.Must(template.ParseFiles("templates/length.html"))
		tmpl.Execute(w, nil)
	}
}