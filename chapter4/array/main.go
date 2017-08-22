package main

import (
	"fmt"
)

func main() {
	var arr1 [5]int
	arr2 := [5]int{1, 2, 3, 4, 5}
	//arr3 := [...]int{1, 2, 3, 4, 5, 6, 7}
	//arr4 := [5]int{7, 8, 9}

	arr1[0] = 2
	arr2[3] = 40

	//所有元素都是指针的数组
	arr5 := [5]*int{0: new(int), 1:new(int)}
	*arr5[0] = 10
	*arr5[1] = 20

	//&arr1
	fmt.Printf("arr1 position %s\n", arr2)
	fmt.Printf("arr2 position %s\n", *arr5[0])

	fmt.Print("HelloWorld ", len(arr1))
}