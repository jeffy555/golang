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
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `<img src="tody.jpg">`)
}

func dogpic(res http.ResponseWriter, req *http.Request) {
	f, err := os.Open("tody.jpg")
	if err != nil {
		http.Error(res, "File Not Found", 404)
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		http.Error(res, "File Not Found", 404)
		return
	}
	//func ServeContent(w ResponseWriter, req *Request, name string, modtime time.Time, content io.ReadSeeker)
	http.ServeContent(res, req, f.Name(), fi.ModTime(), f)
}
