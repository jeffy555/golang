package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	//Listen accepts a connection and port number and it returns a connection and error
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	//We are defering the close
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		//Accept give us conn
		//Conn is an interface that reads and writes data to the connection(conn)
		//Read(b []byte)(n int,err error)
		//Write(b []byte)(n int,err error)

		io.WriteString(conn, "\nHello from TCP Server\n")
		//Ask for writer
		//Writes the string to a writer, conn writes to the writer
		fmt.Fprintln(conn, "Hows your day?")
		fmt.Fprintf(conn, "%v", "Good")
		conn.Close()
	}
}
