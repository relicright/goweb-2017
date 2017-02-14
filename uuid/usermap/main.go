package main

import (
	"html/template"
	"net/http"
	"github.com/satori/go.uuid"
)

type user struct{
	UserName string
	First string
	Last string
}

var tpl *template.Template
var dbUsers = map[string]user{}	//user ID, user
var dbSessions = map[string]string{} // session ID, user ID

func init()  {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main()  {

	http.HandleFunc("/", Index)
	http.HandleFunc("/bar", Bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func Index(w http.ResponseWriter, r *http.Request)  {

	//Get Cookie
	c, err := r.Cookie("session")
	if err != nil {
		sID := uuid.NewV4()
		c = &http.Cookie{
			Name: "session",
			Value: sID.String(),
			//Secure: true,
			HttpOnly: true,
		}
		http.SetCookie(w, c)
	}

	// if the user exist already, get user
	var u user
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}

	//process form submission
	if r.Method == http.MethodPost{
		un := r.FormValue("username")
		f := r.FormValue("firstname")
		l := r.FormValue("lastname")
		u = user{un, f, l}
		dbSessions[c.Value] = un
		dbUsers[un] = u
	}

	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func Bar(w http.ResponseWriter, r *http.Request)  {

	c, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	un, ok := dbSessions[c.Value]
	if !ok{
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	u := dbUsers[un]
	tpl.ExecuteTemplate(w, "bar.gohtml", u)

}