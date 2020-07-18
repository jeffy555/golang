package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/tody", dogpic)
	http.ListenAndServe(":8080", nil)
}

func dog(res http.ResponseWriter, req *http.Request) {
	f, err := os.Open("new.jpg")
	if err != nil {
		//http.Error takes Error(w ResponseWriter, error string, code int)
		http.Error(res, "File not Found", 404)
		return
	}
	defer f.Close()
	//io.Copy : Copy has Copy(dst Writer, src Reader) returns --> (written int64, err error)
	io.Copy(res, f)
}

func dogpic(res http.ResponseWriter, req *http.Request) {
	f, err := os.Open("tody.jpg")
	if err != nil {
		//http.Error takes Error(w ResponseWriter, error string, code int)
		http.Error(res, "File not Found", 404)
		return
	}
	defer f.Close()
	//io.Copy : Copy has Copy(dst Writer, src Reader) returns --> (written int64, err error)
	io.Copy(res, f)
}
