package main

import (
	"io"
	"net/http"
)

//type Handler is an interface that has ServeHttp method(responsewriter, *request), It has following functions that returns this handler
//func FileServer(root FileSystem) Handler
//func NotFoundHandler() Handler
//func RedirectHandler(url string, code int) Handler
//func StripPrefix(prefix string, h Handler) Handler
//func TimeoutHandler(h Handler, dt time.Duration, msg string) Handler

func main() {
	http.HandleFunc("/", dog)
	//func StripPrefix(prefix string, h Handler) Handler : Returning a Handler http.FileServer
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/resources/tody.jpg">`)
}

/*

./assets/tody.jpg

*/
