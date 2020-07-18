package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/tody.jpg", dogPic)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="tody.jpg">	`)
}

func dogPic(w http.ResponseWriter, req *http.Request) {
	//ServeFile : func ServeFile(w ResponseWriter, r *Request, name string)
	http.ServeFile(w, req, "tody.jpg")
}
