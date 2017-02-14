package main

import (
	"net/http"
	"log"
	"github.com/satori/go.uuid"
	"fmt"
)

func main()  {

	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request)  {

	cookie, err := r.Cookie("session")
	if err != nil{
		log.Println(err)

		// Make a UUID
		id := uuid.NewV4()
		cookie = &http.Cookie{
			Name: "session",
			Value: id.String(),
			Secure: true,
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
	}
	fmt.Print(cookie)
}