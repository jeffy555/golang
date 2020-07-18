package main

import (
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func getUser(w http.ResponseWriter, req *http.Request) user {
	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		fmt.Println("Value from C is", c.Value)

	}
	http.SetCookie(w, c)

	// if the user exists already, get user
	var u user
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
		fmt.Println("Lets see what's inside u:", u)
	}
	return u
}

func alreadyLoggedIn(req *http.Request) bool {
	fmt.Println("I am inside AlreadyLoggedIN validation function")
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	un := dbSessions[c.Value]
	fmt.Println("Whats inside un:", un)
	_, ok := dbUsers[un]
	return ok
}
