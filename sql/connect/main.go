package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {
	//Establish a connection
	//func Open(driverName, dataSourceName string) (*DB, error)
	//Returns DB and error. Driver=mysql and username:password@tcp(hostname)/dbname?charset=utf8

	db, err = sql.Open("mysql", "admin:jefferson@tcp(mysql.cymhiqgwvgdn.ap-south-1.rds.amazonaws.com)/app?charset=utf8")
	check(err)
	//Closing the connection at the end
	defer db.Close()
	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	err := http.ListenAndServe(":8080", nil)
	check(err)
}

func index(w http.ResponseWriter, req *http.Request) {
	_, err := io.WriteString(w, "Successfully Completed")
	check(err)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
