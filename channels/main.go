package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {

	ch := make(chan int)
	wg.Add(2)
	go func() {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}()
	go func() {
		i := 42
		ch <- i
		//Go do not care about the below line because channel send and receive already completed.
		i = 43
		wg.Done()
	}()
	wg.Wait()

	//Send only and receive only channel
	//Always pass the value when handling send and receive channel to pass the value between the channels
	ch1 := make(chan int, 50)
	wg.Add(2)
	go func(ch1 <-chan int) {
		for {

			if i2, ok := <-ch1; ok {

				fmt.Println(i2)
			} else {
				break
			}
		}
		wg.Done()

	}(ch1)
	go func(ch1 chan<- int) {
		//What if I send many inputs to one receiver. We can handle using the buffered channels
		//Channel closing: Before starting the for range loop, Sending channel will send the second data so the for loop will be deadlocked
		//In order to avoid this we need to close the channel, so that whenever a loop executes the channel will be closed and the loop can print the value
		//You can also use defer function to close the channel at the last step
		//We can also use ok syntax to check if the channel sends data or not and then we can handle it using a for and if loop

		ch1 <- 51
		ch1 <- 53
		ch1 <- 54
		ch1 <- 56
		defer close(ch1)
		wg.Done()
	}(ch1)
	wg.Wait()
}
