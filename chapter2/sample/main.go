package main

import (
	"log"
	"os"

	_"../sample/matchers"
	"../sample/search"
)

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

// main is the entry point for the program.
func main() {
	// Perform the search for the specified term.
	search.Run("president")
}

/**
	1、每个代码文件都属于一个包，包名应该与代码文件所在的文件夹同名；
	2、Go可以先声明再初始化，也可以初始化的时候声明
	3、使用指针可以在函数件或者 goroutine 间共享数据
	4、goroutine 来并发和同步
	5、Go是面向方法的，和C一样可以建立自己的struct
 */
