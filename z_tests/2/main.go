package main

import (
	"text/template"
	"fmt"
	"net/http"
)

var tpl *template.Template

var fm = template.FuncMap{
	"cs": ChangeString,
}

func ChangeString(s string) string {
	return fmt.Sprintf("%s was the string", s)
}

func init()  {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("templates/index.gohtml"))
}

func main()  {

	http.HandleFunc("/", Homepage)

	fmt.Print(http.ListenAndServe(":8080", nil))
}

func Homepage(w http.ResponseWriter, r *http.Request)  {

	if err := tpl.ExecuteTemplate(w, "index.gohtml", "Hello"); err != nil{
		http.Error(w, err.Error(), 304)
	}
}
