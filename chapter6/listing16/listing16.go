// This sample program demonstrates how to use a mutex
// to define critical sections of code that need synchronous
// access.
// 此示例展示如何使用互斥锁 定义
// 一段 需要 同步访问的代码临界区资源 的同步访问
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	// counter is a variable incremented by all goroutines.
	counter int

	// wg is used to wait for the program to finish.
	wg sync.WaitGroup

	// mutex is used to define a critical section of code.
	// mutex 用来定义一段代码临界区
	mutex sync.Mutex
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
	fmt.Printf("Final Counter: %d\n", counter)
}

// incCounter increments the package level Counter variable
// using the Mutex to synchronize and provide safe access.
// incCounter 使用互斥锁来同步并保证安全访问，
// 增加包里counter变量的值
func incCounter(id int) {
	// Schedule the call to Done to tell main we are done.
	defer wg.Done()

	for count := 0; count < 10; count++ {
		// Only allow one goroutine through this
		// critical section at a time.
		// 同一时刻只允许一个goroutine进入这个
		// 临界区
		mutex.Lock()
		{
			// Capture the value of counter.
			value := counter

			// Yield the thread and be placed back in queue.
			// 当前goroutine从线程退出，并放回到队列。在锁🔐面前不值一提
			fmt.Printf("goroutine %d counter is %d\n", id, counter)
			runtime.Gosched()

			// Increment our local value of counter.
			value++

			// Store the value back into counter.
			counter = value
		}
		mutex.Unlock()
		// Release the lock and allow any
		// waiting goroutine through.
		// 释放锁，允许其他正在等待的goroutine
		// 进入临界区
	}
}
