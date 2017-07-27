package structure

import "fmt"

func Slince()  {
	sli := []string{"a", "b", "c", "d", "e"}
	fmt.Printf("sli1 : %v\n", sli)
	change_sli(sli, 1)
	fmt.Printf("sli2 : %v\n", sli)
}

func change_sli(sli []string, key int)  {
	sli[key] = "z"
}

func Map()  {
	ma := map[string]int{"abc":123, "bcd":234}

	for key, value := range ma{
		fmt.Printf("key : %s, value : %d \n", key, value)
	}

	change_map(ma, "ac")
	fmt.Printf("---------\n")
	for key, value := range ma{
		fmt.Printf("key : %s, value : %d \n", key, value)
	}
}
/*
	0 默认整数
	"" 默认字符串
	nil 默认其他一些结构
 */
func change_map(ma map[string]int, key string)  {
	if ma[key] != 0 {
		ma[key] = 0
	}
}