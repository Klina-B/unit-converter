package main

import (
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles("templates/body.html"))

type PageData struct{
	Title string
	Message string
}

func handler(writer http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		req.ParseForm()
		name := req.FormValue("name")

		data := PageData{
			Title: "Result",
			Message: "Hello, " + name,
		}
		tmpl.Execute(writer, data)
		return
	}

	data := PageData{
		Title: "Main page",
		Message: "",
	}
	tmpl.Execute(writer, data)
}

func main(){
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}