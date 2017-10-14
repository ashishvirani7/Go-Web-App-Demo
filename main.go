package main

import (
	"net/http"
	"html/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request){
	type customData struct {
		Title string
		Members []string
	}

	cd := customData{
		Title: "About Me",
		Members: []string{"Ashish","Virani",},
	}

	tpl.ExecuteTemplate(w, "index.gohtml", cd)
}
