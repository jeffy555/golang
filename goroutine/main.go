package main

import (
	"fmt"
	"sync"
	"time"
)

//Go Routine
var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main() {
	//SayHello function will not print the Say Hello. It will run inside the main function itself. Never actually have any time to print it. We can overcome using Sleep
	go SayHello()
	time.Sleep(100 * time.Millisecond)

	var msg = "Jeffy"
	go func(msg string) {

		fmt.Println(msg)
	}(msg)
	//msg is overridden by Goodbye so if we print out then Goodbye will be printed and not hello.
	// But if we pass a message parameter and pass it by value Hello will be passed here then we can decouple the msg in the main function from goroutine
	// Final output will be hello. But this is not a best approach. Best approach is WaitGroup.

	msg = "Goodbyeeee"
	time.Sleep(100 * time.Millisecond)
	wg.Add(1)
	str := "HelloWorld"
	go func(str string) {

		fmt.Println(str)
		wg.Done()
	}(str)
	str = "HelloGO"
	wg.Wait()
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go Bye()
		m.Lock()
		go increment()
	}
	wg.Wait()

}

//Race condition will happen during the waitgroup as well. To overcome this we are using Mutex
//Mutex bring infinite Readers but only one writer
//Only part of the problem is fixed. Because Counter function is not getting called properly.
//Now we are moving the mutex lock to the go routine function to do the proper execution and the Unlock will still be there inside the function.

func Bye() {
	//m.RLock()
	fmt.Println("Hello Counter", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	//m.Lock()
	counter++
	m.Unlock()
	wg.Done()

}

func SayHello() {

	fmt.Println("Say Hello")

}

//Best Practices
//Dont create goroutines in libraries
//Avoid Subtle memory leaks
//check for race conditions at compile time
//go run -race main.go --> to check race condition during compile time
