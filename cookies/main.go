package main

import (
	"log"
	"net/http"
	"fmt"
)

func handler(w http.ResponseWriter, req *http.Request) {

	c, err := req.Cookie("ssid")
	if err != nil{
		c = &http.Cookie{
			Name: "ssid",
			Value: "1234565",
			MaxAge: 60,
		}

		http.SetCookie(w, c)
	}

	fmt.Print(c.Value)

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
}

func main() {
	http.HandleFunc("/", handler)
	log.Printf("About to listen on 10443. Go to https://127.0.0.1:10443/")
	err := http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil)
	log.Fatal(err)
}
