package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template
var tpl2 *template.Template

func init() {

	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main() {

	http.HandleFunc("/", Builtin)

	fmt.Print(http.ListenAndServe(":8080", nil))

}

func Builtin(w http.ResponseWriter, r *http.Request) {

	list := []string{"one", "two", "three", "four", "five",}

	err := tpl.ExecuteTemplate(w, "index.gohtml", list)
	if err != nil {
		log.Fatalln(err)
	}
}


