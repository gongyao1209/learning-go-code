// This sample program demonstrates how to use the work package
// to use a pool of goroutines to get work done.

// 使用goroutines池进行工作
package main

import (
	"../../work"
	"fmt"
	"log"
	"runtime"
	"time"
)

// names provides a set of names to display.
var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
}

// namePrinter provides special support for printing names.
type namePrinter struct {
	name string
}

// Task implements the Worker interface.
func (m *namePrinter) Task() {
	log.Println(m.name)
	time.Sleep(time.Second)
}

// main is the entry point for all Go programs.
func main() {
	// 创建两个进程的工作池
	// Create a work pool with 2 goroutines.
	p := work.New(2)

	fmt.Println(runtime.NumGoroutine())

	p.Shutdown()

	fmt.Println(runtime.NumGoroutine())

	return
	//return
	//var wg sync.WaitGroup
	//wg.Add(100 * len(names))

	for i := 0; i < 100; i++ {
		// Iterate over the slice of names.
		for _, name := range names {
			// Create a namePrinter and provide the
			// specific name.
			np := namePrinter{
				name: name,
			}

			//go func() {
				// Submit the task to be worked on. When RunTask
				// returns we know it is being handled.
				p.Run(&np)
				//wg.Done()
			//}()

			fmt.Println(runtime.NumGoroutine())
		}
	}

	//wg.Wait()

	// Shutdown the work pool and wait for all existing work
	// to be completed.
	p.Shutdown()
}
