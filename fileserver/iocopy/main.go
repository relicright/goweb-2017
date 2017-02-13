package main

import (
	"net/http"
	"io"
	"os"
	"log"
)

func main()  {

	http.HandleFunc("/", dog)
	http.HandleFunc("/toby.jpg", dogPic)

	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, r *http.Request)  {

	w.Header().Set("Content-Type", "text/html; charset=urf-8")

	io.WriteString(w, `<img src="/toby.jpg">`)
}

func dogPic(w http.ResponseWriter, r *http.Request)  {

	// Open the file on the server
	f, err := os.Open("toby.jpg")
	if err != nil{
		log.Println(err)
	}
	// Defer the close of the file
	defer f.Close()

	//Copy the byts of the file to the response writer
	io.Copy(w, f)
}
