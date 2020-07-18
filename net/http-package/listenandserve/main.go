package main

import (
	"fmt"
	"net/http"
)

type jeffy int

func (j jeffy) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	//j is now an handler because it has the method ServeHTTP

	fmt.Fprintln(w, "Hello Jeff")

}

func main() {
	var l jeffy
	//ListenandServe
	//ListenandServe always return a non-nil error
	http.ListenAndServe(":8080", l)
}
