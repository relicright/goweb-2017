package main

import (
	"net/http"
	"io"
	"time"
)

func main()  {

	http.HandleFunc("/", dog)
	http.HandleFunc("/1", dogPic)

	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, r *http.Request)  {

	w.Header().Set("Content-Type", "text/html; charset=urf-8")

	io.WriteString(w, `<img src="/1">`)
}

func dogPic(w http.ResponseWriter, r *http.Request)  {

	time.Sleep(5 * time.Second)
	http.ServeFile(w, r, "1.doc")
}
