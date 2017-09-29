package handlers

//Created By: Ricky Clevinger
//Updated On: 8/17/2017
//Last Updated By: Ricky Clevinger

import (
	"html/template"
	"net/http"
	"regexp"
	"modify"
)

var validPath = regexp.MustCompile("^/(index.html)$")
var templates = template.Must(template.ParseFiles("views/index.html"))

type Page struct {
}

func LoadPage() *Page {

	return &Page{}
}

//Renders HTML page
func RenderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//Validates path and calls handler
func MakeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r)
	}
}

//Handles the index page
func IndexHandler(w http.ResponseWriter, r *http.Request) {

	p := LoadPage()
	RenderTemplate(w, "index", p)
}


//Redirect to login.html
func Redirect(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "/index.html", 301)
}


//Handles
func Handles() {

	http.HandleFunc("/index.html", MakeHandler(IndexHandler))
	http.HandleFunc("/modify", modify.Modify)
	http.HandleFunc("/", Redirect)
}
