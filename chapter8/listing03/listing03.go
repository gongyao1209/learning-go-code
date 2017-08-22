// This sample program demonstrates how to use the base log package.
package main

import (
	"log"
)

func init() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	// Println writes to the standard logger.
	// 写在标准日志记录器里面
	log.Println("message")

	// Fatalln is Println() followed by a call to os.Exit(1).
	// 在调用println之后紧接着调用os.Exit(1)
	log.Fatalln("fatal message")

	// Panicln is Println() followed by a call to panic().
	// 调用Println() 之后 紧接着调用panic()
	log.Panicln("panic message")
}
