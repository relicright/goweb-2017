package main

import (
	"net/http"
	"fmt"
	"html/template"
	"log"
)

type Hotdog int

var fm = template.FuncMap{
	"changeAge": ChangeAge,
}

func ChangeAge(a int) int {
	return a * 3
}

func (h Hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	test := make(map[string]int)
	test["Janet"] = 22
	test["James"] = 33
	test["john"] = 44

	fmt.Print(test)

	str := []string{"one", "two", "Three", "Four"}

	if err := templates.ExecuteTemplate(w, "index.gohtml", str); err != nil{
		http.Error(w, err.Error(), 392)
	}
}

var templates *template.Template
func init()  {
	templates = template.Must(template.New("").Funcs(fm).ParseGlob("templates/*.gohtml"))
}

func main()  {

	//var h Hotdog

	http.HandleFunc("/", Homepage)

	fmt.Print(http.ListenAndServe(":8080", nil))
}

func Homepage(w http.ResponseWriter, r *http.Request)  {

	err := r.ParseForm()
	if err != nil{
		log.Println(err)
	}

	fmt.Print(r.Form)
	fmt.Print(r.PostForm)

	data := struct{
		ContentLen int64
		Method string
		Header http.Header
		URL string
	}{
		r.ContentLength,
		r.Method,
		r.Header,
		r.URL.Path,
	}

	if err := templates.ExecuteTemplate(w, "index.gohtml", data); err != nil{
		http.Error(w, err.Error(), 304)
	}
}
