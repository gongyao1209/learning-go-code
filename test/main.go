package main

import (
	//chap5 "../test/chapter5"
	"fmt"
	//"runtime"
	"os"
	"log"
)

func main()  {
	//name2 := string("gongyao")
	//fmt.Printf("%s", name2)
	//chap5.Test(name2)
	//
	//var yaoke chap5.User
	//yaoke.SetName("yaoke")
	//yaoke.SetMobile("186")
	//yaoke.SetMobile1("176")
	//yaoke.Show()
	//
	//fmt.Printf("Cpu number is :%d\n", runtime.NumCPU())
	//
	//fmt.Print(chap5.New(12))

	file_name := "/Users/gongyao/workspace/go_code1/code/chapter3/wordcount/gowords.txt"
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Read %d bytes: %q\n", count, data[:count])

	os.Getpagesize()

	fmt.Printf("%d", os.Getpagesize())

	//src := io.Reader("aaa")
	//var s scanner.Scanner
	//
	//s.Init(src)
	//tok := s.Scan()
	//for tok != scanner.EOF {
	//	tok = s.Scan()
	//}
}
