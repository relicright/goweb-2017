package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"
	"goweb/templatepipelines/pipeline"
)

var tpl *template.Template
var tpl2 *template.Template

func monthDayYear(t time.Time) string {
	return t.Format("01-02-2006")
}

var fm = template.FuncMap{
	"fdayMDY": monthDayYear,
}

func init() {
	//You can have different functions for different templates
	// This this example we have a second template that draws it's functions from a
	// FuncMap located inside of the 'pipeline' package
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("templates/index.gohtml"))
	tpl2 = template.Must(template.New("").Funcs(pipeline.Fm2).ParseFiles("templates/index2.gohtml"))
}

func main() {

	http.HandleFunc("/", TimePage)
	http.HandleFunc("/2", PipelinePage)

	fmt.Print(http.ListenAndServe(":8080", nil))

}

func TimePage(w http.ResponseWriter, r *http.Request) {

	err := tpl.ExecuteTemplate(w, "index.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}

func PipelinePage(w http.ResponseWriter, r *http.Request) {

	err := tpl2.ExecuteTemplate(w, "index2.gohtml", 22)
	if err != nil {
		log.Fatalln(err)
	}
}
