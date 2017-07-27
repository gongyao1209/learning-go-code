package main

import (
	a "../structure"
)

func main() {

	a.Slince()
	a.Map()
	////利用make创建切片。只指定长度，那么切片的容量和长度相等
	////slice := make([]string, 5)
	//
	////3-表示长度，5-容量。容量>=长度。长度--看到的长度，容量--还有很多看不到的。冰山一角，一角是长度，冰山是容量
	////slice1 := make([]int, 3, 5)
	//
	////使用字面量来建立切片，[]
	////slice2 := []string{"red", "blue", "green", "yellow"}
	//
	////slice3 := []int{10, 20, 30}
	//
	////设置长度和容量
	////slice4 := []string{99: ""}
	//
	//
	///*
	//切片和数组差别：数组大小确定，切片大小不确定
	// */
	////array := [5]string{10, 20, 30}
	//slice := []string{"a", "b", "c", "d", "e"}
	//slice[0] = "f"
	//
	////创建一个新的切片 就是 把底层数组切出一部分
	//slice1 := slice[1:3]
	//
	//slice1[0] = "g"
	//
	//slice1 = append(slice1, "z")
	//
	//slice1 = append(slice1, "aa")
	//slice1 = append(slice1, "bb")
	//
	//
	//slice2 := []string{"a", "b", "c", "d"}
	//slice3 := slice2[1:3:3]
	//slice3 = append(slice3, "e")
	//
	//slice3[2] = "z"
	//
	////fmt.Printf("slice2: %v\n", slice2)
	////fmt.Printf("slice3: %v\n", slice3)
	//
	//slice4 := append(slice2, slice3...)
	////fmt.Printf("slice4: %v\n", slice4)
	//
	//for index, value := range slice4{
	//	//所有迭代的 value 的地址都是相同的，依次变化
	//	fmt.Printf("value:%s, addr1:%x, addr2:%x\n", value, &value, &slice4[index])
	//	//fmt.Printf("slice4 %d is %s\n", index, value)
	//}
	//
	//
	////dict := make(map[string]int)
	//dict := map[string]string{"Red":"#005612", "Yellow":"#89786"}
	//fmt.Printf("%s\n", dict["Red"])
	//dict["Blue"] = "##6543"
	//fmt.Printf(dict["Blue"])
	//value, exists := dict["Blue"]
	//if exists {
	//	fmt.Printf(value)
	//}
	//
	//var color map[string]string
	//if color != nil {
	//	fmt.Printf("color not nil")
	//} else {
	//	fmt.Printf("color nil")
	//}
	////color["Red"] = "#005612"
	////fmt.Printf(color["Red"])
}
