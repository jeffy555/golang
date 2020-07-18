package main

import (
	"fmt"
	"html/template"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

//Fields inside the Struct User (Contains signup details which will be fetch from the html pages)
type user struct {
	UserName string
	Password string
	First    string
	Last     string
}

var tpl *template.Template

//Map --> key value pair
//dbUsers has userId as key and value will be the users
//dbSessions has sessionID has key and user ID as value --> dbSessions value has --> Composite literal of strings

var dbUsers = map[string]user{}      // user ID, user
var dbSessions = map[string]string{} // session ID, user ID

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/hello", hello)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	//getUser method is used in sessions.go to check and validate the user exist already
	u := getUser(req)
	tpl.ExecuteTemplate(w, "index.gohtml", u)
	fmt.Println("I am inside Index")
}

func bar(w http.ResponseWriter, req *http.Request) {
	u := getUser(req)
	//Check if the user is already logged in
	if !alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
	fmt.Println("I am inside Bar")
}

func signup(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(w, req, "/hello", http.StatusSeeOther)
		return
	}

	// process form submission
	if req.Method == http.MethodPost {

		// get form values
		un := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")

		// username taken?
		if _, ok := dbUsers[un]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}

		// create session
		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
		dbSessions[c.Value] = un

		// store user in dbUsers
		u := user{un, p, f, l}
		dbUsers[un] = u

		// redirect
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
	fmt.Println("Form Processed")
}

func hello(w http.ResponseWriter, req *http.Request) {
	u := getUser(req)
	//Check if the user is already logged in
	if !alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "home.gohtml", u)
	fmt.Println("I am inside home")
}
