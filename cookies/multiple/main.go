package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/lot", lot)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie",
		Value: "some value",
		Path:  "/",
	})
	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(w, "in chrome go to: dev tools / application / cookies")
}

func read(w http.ResponseWriter, req *http.Request) {

	c1, err := req.Cookie("my-cookie")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "Your COOKIE is:", c1)
	}

	c2, err := req.Cookie("good")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "Your COOKIE is:", c2)
	}

	c3, err := req.Cookie("bad")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "Your COOKIE is:", c3)
	}
}

func lot(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "good",
		Value: "Jehovah",
	})

	http.SetCookie(w, &http.Cookie{
		Name:  "bad",
		Value: "Lucifer",
	})

	fmt.Fprintln(w, "COOKIES WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(w, "in chrome go to: dev tools / application / cookies")
}
