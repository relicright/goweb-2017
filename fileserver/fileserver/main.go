package main

import (
	"net/http"
	"io"
)

func main()  {

	http.HandleFunc("/", dog)

	// "Resources" will be the folder you designate where the file server will look when you ask
	// to serve a file.  The path "/resources" will be stripped once 'http.StripPrefix' is called.
	// Then, the 'http.FileServer' will use the directory folder "assets" to serve all files from.
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Using the path "/resources" will automatically look in the
	// designated asset folder
	io.WriteString(w, `<img src="/resources/toby.jpg">`)
}