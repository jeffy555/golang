package main

import (
	"fmt"
	"log"
)

func main() {
	//Defer Program
	//Defer a Println or function, this will instruct the Go execution to run at the end of the program. It will execute only after the func main().
	//If all the Println or function is in a deferred state then the last in the first out model will be executed.
	// In the below example technically it should have printed out end. But it will print out start.
	//The argument will pick during the time of the defer is called  and not when the function is executed

	//a := "Welcome"
	//defer fmt.Println(a)
	//a = "end"

	//Panic will happen after the deferred statement
	// Panic should happen only if it is really required. If the log file is not there in the server, please don't write the panic statement
	//Panic can be used in the HTTP web server when the port is not enabled or not listening on --> there is a reason to panic

	//fmt.Println("Start")
	//defer fmt.Println("Deferred")
	//panic("Bad thing happened")
	//fmt.Println("End")
	fmt.Println("-----------")

	//Recover Function
	//Used to recover from Panic
	fmt.Println("Recover Example")
	panicker()
	fmt.Println("Done")
}

func panicker() {
	fmt.Println("About to panic")
	defer func() {

		err := recover()
		if err != nil {
			log.Println(err)
			panic(err)
		}
	}()
	panic("Soemthing bad happened")
	fmt.Println("finished")
}
