package main

import (
	chap5 "../test/chapter5"
	"fmt"
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

	var counter chap5.AlertCounter
	counter = chap5.New(11)
	fmt.Printf("Counter: %d\n", counter)
	//str := "gongyao"
	//= chap5.GetName(str)
}

