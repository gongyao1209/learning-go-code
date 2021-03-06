// This sample program demonstrates how to use the atomic
// package to provide safe access to numeric types.
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	// counter is a variable incremented by all goroutines.
	counter int64

	// wg is used to wait for the program to finish.
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

	// Display the final value.
	fmt.Println("Final Counter:", counter)
}

// incCounter increments the package level counter variable.
func incCounter(id int) {
	// Schedule the call to Done to tell main we are done.
	// 在函数退出时，调用Done来通知main goroutine 已经完成任务
	defer wg.Done()

	for count := 0; count < 200000; count++ {
		// Safely Add One To Counter.
		// 安全的对counter加1
		//counter += 1
		atomic.AddInt64(&counter, 1)

		// Yield the thread and be placed back in queue.
		runtime.Gosched()
	}
}
