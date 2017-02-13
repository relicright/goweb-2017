package main

import (
	"net/http"
	"io"
)

type Hotdog int

func (h Hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This is for dogs")
}

type Hotcat int

func (h Hotcat) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This is for cats")
}

func main()  {

	var d Hotdog
	var c Hotcat

	mux := http.NewServeMux()
	mux.Handle("/dogs", d)
	mux.Handle("/cats", c)

	http.ListenAndServe(":8080", mux)
}
