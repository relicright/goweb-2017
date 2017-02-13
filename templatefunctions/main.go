package main

import (
	"text/template"
	"strings"
	"fmt"
	"net/http"
)

var tpl *template.Template

// Create a FuncMap to register functions.
// 'uc' is what the func will be called in the template
// 'uc' is the ToUpper func from package strings
// 'ft' is a func I declared
// 'ft' slices a string, returning the first three characters
var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}






func init()  {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("templates/index.gohtml"))
}

type sage struct{
	Name string
	Age int
}

type sages struct{
	List []sage
}

func main()  {

	http.HandleFunc("/", Homepage)

	fmt.Print(http.ListenAndServe(":8080", nil))
}

func Homepage(w http.ResponseWriter, r *http.Request)  {

	b := sage{
		"gandhi",
		55,
	}

	c := sage{
		"MLK",
		32,
	}

	d := sage{
		"Jesus",
		31,
	}

	l := []sage{b, c, d}

	if err := tpl.ExecuteTemplate(w, "index.gohtml", l); err != nil{
		http.Error(w, err.Error(), http.StatusNotFound)
	}
}
