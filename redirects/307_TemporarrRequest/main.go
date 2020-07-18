package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	//Must is a helper that wraps a call to a function returning (*Template, error) and panics if the error is non-nil. It is intended for use in variable initializations such as
	//var t = template.Must(template.New("name").Parse("html"))
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	//HandleFunc(pattern string, handler func(ResponseWriter, *Request)){
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/barred", barred)
	//Handle
	//Handle(pattern string, handler Handler)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	//ListenAndServe listens on the TCP network address addr and then calls Serve with handler to handle requests on incoming connections. Accepted connections are configured to enable TCP keep-alives.
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at foo:", req.Method)
}

func bar(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at bar:", req.Method)
	//func Redirect(w ResponseWriter, r *Request, url string, code int)
	//Redirect replies to the request with a redirect to url, which may be a path relative to the request path.
	//process form submission here
	http.Redirect(w, req, "/", http.StatusTemporaryRedirect)
	//307 (Temporary Redirect)
}

func barred(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at barred:", req.Method)
	//ExecuteTemplate(wr io.Writer, name string, data interface{}) error
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
