package main

import (
	"text/template"
	"net/http"
	"fmt"
)

var templates *template.Template

func init()  {
	templates = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main()  {

	http.HandleFunc("/", Homepage)

	fmt.Print(http.ListenAndServe(":8080", nil))
}

func Homepage(w http.ResponseWriter, r *http.Request)  {

	if err := templates.ExecuteTemplate(w, "index.gohtml", 12); err != nil{
		http.Error(w, err.Error(), 304)
	}
}
