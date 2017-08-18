// This sample program demonstrates how the goroutine scheduler
// will time slice goroutines on a single thread.
package main

import (
	"fmt"
	"runtime"
	"sync"
)

// wg is used to wait for the program to finish.
// wg用来等待程序完成
var wg sync.WaitGroup

// main is the entry point for all Go programs.
// main 是所有go程序入口
func main() {
	// Allocate 1 logical processors for the scheduler to use.
	// 分配一个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(1)

	// Add a count of two, one for each goroutine.
	// 计数器加2，表示等待两个goroutune执行
	wg.Add(2)

	// Create two goroutines.
	fmt.Println("Create Goroutines")
	go printPrime("A")
	go printPrime("B")

	// Wait for the goroutines to finish.
	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("Terminating Program")
}

// printPrime displays prime numbers for the first 5000 numbers.
func printPrime(prefix string) {
	// Schedule the call to Done to tell main we are done.
	defer wg.Done()

next1:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next1
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("Completed", prefix)
}
