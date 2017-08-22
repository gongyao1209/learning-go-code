package main

import (
	"fmt"
	sync1 "sync"
	"database/sql"
)

var wg sync1.WaitGroup

func printer(ch chan int) {
	for i := range ch {
		fmt.Printf("Received %d \n", i)
	}
	wg.Done()
}

func printer1(ch chan int) {
	for i := range ch{
		fmt.Printf("gongyao %d\n", i)
	}
	wg.Done()
}

func init() {
	sql.Open("postgres", "mydb");
}
// main is the entry point for the program.
/*
	Go和C很一样，你只用你的东西，它的东西你都不要动，也不要想看看其中有什么。会发生未知的错误
 */
func main() {
	c := make(chan int)
	wg.Add(2)

	go printer(c)
	go printer1(c)

	// Send 10 integers on the channel.
	for i := 1; i <= 10; i++ {
		c <- i
	}

	close(c)
	wg.Wait()
}
