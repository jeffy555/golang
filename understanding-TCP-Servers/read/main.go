package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	//Listen to the connection
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	//Defer the close later
	defer li.Close()

	for {
		//Accepts the connection
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		//Handling Go routine
		go handle(conn)
	}
}

//Handling it in my own connection
//Passing the input
//
func handle(conn net.Conn) {
	//New scanner has a pointer to a *scanner, it has method Text()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		//Scan() advances the scanner to the next token, which will then be available through bytes or text method
		line := scanner.Text()
		fmt.Println(line)
	}
	defer conn.Close()
	fmt.Println("Code comes here")
}
