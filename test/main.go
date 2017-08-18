package main

import (
	chap5 "../test/chapter5"
	"fmt"
	"runtime"
)

func main()  {
	name2 := string("gongyao")
	fmt.Printf("%s", name2)
	chap5.Test(name2)

	var yaoke chap5.User
	yaoke.SetName("yaoke")
	yaoke.SetMobile("186")
	yaoke.SetMobile1("176")
	yaoke.Show()

	fmt.Printf("Cpu number is :%d\n", runtime.NumCPU())
}

