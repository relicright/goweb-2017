package main

import (
	"time"
	"html/template"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"github.com/nu7hatch/gouuid"
	"errors"
	"fmt"
)

type session struct{
	Un string
	SessionTime time.Time
}

type user struct{
	Email string
	UserName string
	Password string
	Role string
}

var tpl *template.Template
var userdb = map[string]user{}
var sessiondb = map[string]session{}
var cleansession time.Time

func init()  {
	tpl = template.Must(template.ParseGlob("templates/*gohtml"))

	bs, err := bcrypt.GenerateFromPassword([]byte("pass"), bcrypt.MinCost)
	if err != nil{
		log.Println(err)
		return
	}
	u := user{"test@aol.com", "relic", string(bs), "admin"}
	u2 := user{"test1@aol.com", "relic1", string(bs), "admin1"}
	userdb[u.Email] = u
	userdb[u2.Email] = u

	cleansession = time.Now()
}

func main()  {

	http.HandleFunc("/login", LoginPage)
	http.HandleFunc("/loggedin", LoggedinPage)

	http.ListenAndServe(":8080", nil)

}

func LoginPage(w http.ResponseWriter, r *http.Request)  {

	//Check for cookie
	c, err := r.Cookie("ssid")
	if err != nil{
		sID, err := uuid.NewV4()
		if err != nil{
			log.Println(err)
			return
		}
		c = &http.Cookie{
			Name: "ssid",
			Value: sID.String(),
			HttpOnly: true,
		}

		http.SetCookie(w, c)
	}

	if r.Method == http.MethodPost{

		// Pull loging info from the forms
		em := r.FormValue("email")
		pw := r.FormValue("password")

		// Range over the member database
		for k, v := range userdb{
			// If the key equals the login email address
			if k == em{
				// Compare the password for that user
				err := bcrypt.CompareHashAndPassword([]byte(v.Password), []byte(pw))
				if err != nil{
					log.Println("Username/password not found")
					http.Redirect(w, r, "/login", http.StatusSeeOther)
					return
				}
				// Add a session to the DB
				sessiondb[c.Value] = session{v.Email, time.Now()}
				// Redirect to logged in page
				http.Redirect(w, r, "/loggedin", http.StatusSeeOther)
				fmt.Print(sessiondb)
				CleanSessions()
			}
		}
	}

	tpl.ExecuteTemplate(w, "login.gohtml", nil)
}

func LoggedinPage(w http.ResponseWriter, r *http.Request)  {

	c, err := IsLoggedIn(w, r)
	if err != nil{
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

	var u user
	if v, ok := sessiondb[c.Value]; ok{
		u = userdb[v.Un]
		sessiondb[c.Value] = session{v.Un, time.Now()}
	}

	tpl.ExecuteTemplate(w, "loggedin.gohtml", u)
}

func IsLoggedIn(w http.ResponseWriter, r *http.Request) (*http.Cookie, error) {

	//check cookie
	c, err := r.Cookie("ssid")
	if err != nil{
		c = &http.Cookie{
			Name: "ssid",
			Value: "",
			HttpOnly: true,
			MaxAge: -1,
		}
		http.SetCookie(w, c)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return c, errors.New("Not logged in")
	}

	//Check if user exist
	if _, ok := sessiondb[c.Value]; !ok{
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return c, errors.New("member does not exists")
	}

	return c, nil
}

func CleanSessions()  {

	for k, v := range sessiondb{
		if time.Now().Sub(v.SessionTime) > (time.Second * 5){
			delete(sessiondb, k)
			fmt.Print("deleted", k)
		}
	}
}

