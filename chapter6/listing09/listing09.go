// This sample program demonstrates how to create race
// conditions in our programs. We don't want to do this.
// 这个示例程序展示如何在程序里造成竞争状态。多个goroutine访问相同的资源
// 实际上不希望出现这种情况
package main

import (
	"fmt"
	"sync"
	"runtime"
)

var (
	// counter is a variable incremented by all goroutines.
	// counter 是所有goroutine都要增加其值的变量
	counter int

	// wg is used to wait for the program to finish.
	// wg用来等待程序结束
	wg sync.WaitGroup
)

// main is the entry point for all Go programs.
func main() {
	// Add a count of two, one for each goroutine.
	wg.Add(2)

	// Create two goroutines.
	go incCounter(1)
	go incCounter(2)

	// Wait for the goroutines to finish.
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

// incCounter increments the package level counter variable.
func incCounter(id int) {
	// Schedule the call to Done to tell main we are done.
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// Capture the value of Counter.
		value := counter

		// Yield the thread and be placed back in queue.
		// 当前 goroutine 从线程退出，并且放回到队列
		runtime.Gosched()

		// Increment our local value of Counter.
		value++

		// Store the value back into Counter.
		counter = value
	}
}
