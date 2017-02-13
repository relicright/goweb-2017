package main

import (
	"net/http"
	"fmt"
	"html/template"
)

var templates *template.Template
var templates2 *template.Template
func init()  {
	templates = template.Must(template.New("").ParseFiles("templates/index.gohtml"))
	templates2 = template.Must(template.New("").Funcs(fm).ParseFiles("templates/index2.gohtml"))
}

func main()  {

	http.HandleFunc("/1", Homepage)
	http.HandleFunc("/2", NextPage)

	fmt.Print(http.ListenAndServe(":8080", nil))
}

type Person struct{
	Name string
	Age int
}

func (p Person) DoubleAge() int {
	return p.Age * 2
}

func (p Person) AddToName() string {
	return p.Name + " This is now added"
}

var fm = template.FuncMap{
	"adder": AddToName2,
}

func AddToName2(s string) string {
	return s + "This is added #2"
}

func Homepage(w http.ResponseWriter, r *http.Request)  {

	p := Person{
		"James",
		22,
	}
	if err := templates.ExecuteTemplate(w, "index.gohtml", p); err != nil{
		http.Error(w, err.Error(), http.StatusNotFound)
	}
}

func NextPage(w http.ResponseWriter, r *http.Request)  {

	p := "Simon"

	if err := templates2.ExecuteTemplate(w, "index2.gohtml", p); err != nil{
		http.Error(w, err.Error(), http.StatusNotFound)
	}
}