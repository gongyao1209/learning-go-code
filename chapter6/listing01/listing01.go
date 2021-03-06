// This sample program demonstrates how to create goroutines and
// how the scheduler behaves.
package main

import (
	"fmt"
	"runtime"
	"sync"
)

// main is the entry point for all Go programs.
func main() {
	// Allocate 1 logical processor for the scheduler to use.
	// 分配一个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(2)

	// wg is used to wait for the program to finish.
	// wg 用来等待程序完成
	// Add a count of two, one for each goroutine.
	// 计数器 加2，表示要等待两个 goroutine
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")

	// Declare an anonymous function and create a goroutine.
	// 声明一个匿名函数，并创建一个goroutine
	go func() {
		// Schedule the call to Done to tell main we are done.
		// 在函数退出时 调用Done 来通知main：函数工作已经完成
		defer wg.Done()

		// Display the alphabet three times
		// 显示字母表3次
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
			fmt.Println()
		}
	}()

	// Declare an anonymous function and create a goroutine.
	// 声明一个匿名函数，并创建一个goroutine
	go func() {
		// Schedule the call to Done to tell main we are done.
		defer wg.Done()

		// Display the alphabet three times
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
			fmt.Println()
		}
	}()

	// Wait for the goroutines to finish.
	// 不管 goroutine的执行情况，main函数首先会执行。他们愿意说什么就说什么吧，我们必将平稳的达到我们的既定目标
	fmt.Println("Waiting To Finish")
	// 这里等待两个 goroutine完成工作
	wg.Wait()

	fmt.Println("\nTerminating Program")
}
