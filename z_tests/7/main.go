package main

import (
	"log"
	"net/http"
	"os"
	"io"
	"html/template"
)


var templates1 *template.Template
var templates2 *template.Template

var fm1 = template.FuncMap{
	"thing1": Thing1,
}

func Thing1(s string) string {
	return s + " this is added with thing 1"
}


type data struct {
	Name string
	Age string
}

func (d data) ChangeData(s string) string {
	return d.Name + s + " was added onto the data"
}

func init()  {

	templates1 = template.Must(template.New("").Funcs(fm1).ParseGlob("templates/index.gohtml"))
	templates2 = template.Must(template.ParseGlob("templates/1.gohtml"))
}

func main()  {

	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))

	http.HandleFunc("/", IndexPage)
	http.HandleFunc("/one", OnePage)

	http.HandleFunc("/1", handle1)
	http.HandleFunc("/2", handle2)
	http.HandleFunc("/3", handle3)


	log.Fatal(http.ListenAndServe(":8080", nil))
}

func IndexPage(w http.ResponseWriter, r *http.Request)  {

	if err := templates1.ExecuteTemplate(w, "index.gohtml", "Joejo"); err != nil{
		http.Error(w, err.Error(), 404)
	}
}

func OnePage(w http.ResponseWriter, r *http.Request)  {

	d := data{"Jim", "tewnty"}

	if err := templates2.ExecuteTemplate(w, "1.gohtml", d); err != nil{
		http.Error(w, err.Error(), 404)
	}
}

func handle1(w http.ResponseWriter, r *http.Request)  {

	//file serve
	http.ServeFile(w, r, "assets/toby.jpg")
}

func handle2(w http.ResponseWriter, r *http.Request)  {

	//File copy
	f, err := os.Open("assets/toby.jpg")
	if err != nil{
		log.Println(err)
		return
	}
	defer f.Close()

	io.Copy(w, f)
}

func handle3(w http.ResponseWriter, r *http.Request)  {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/resources/toby.jpg">`)
}
