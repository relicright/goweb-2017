package main

import (
	"net/http"
	"io"
	"os"
	"log"
	"io/ioutil"
)

func main()  {

	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/dogcopy", dogcopy)
	http.HandleFunc("/dogserv", dogserver)
	http.HandleFunc("/dogcreate", dogcreate)
	http.HandleFunc("/doghtml", doghtml)

	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, r *http.Request)  {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/resources/toby.jpg">`)
}

func dogcopy(w http.ResponseWriter, r *http.Request)  {

	f, err := os.Open("assets/toby.jpg")
	if err != nil{
		log.Println(err)
	}
	defer f.Close()

	io.Copy(w, f)
}

func dogserver(w http.ResponseWriter, r *http.Request)  {

	http.ServeFile(w, r, "assets/1.doc")
}

func dogcreate(w http.ResponseWriter, r *http.Request)  {

	f, err := os.Create("assets/newone.doc")
	if err != nil{
		log.Println()
		return
	}

	defer f.Close()

	f1, err := os.Open("assets/1.doc")
	if err != nil{
		log.Println()
		return
	}

	bs, err := ioutil.ReadAll(f1)
	if err != nil{
		log.Println(err)
	}
	f.Write(bs)
}

func doghtml(w http.ResponseWriter, r *http.Request)  {

	f, err := os.Create("assets/1.html")
	if err != nil{
		log.Println()
		return
	}

	defer f.Close()

	f.WriteString(`<img src="/resources/toby.jpg">`)
}
