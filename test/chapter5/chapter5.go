package chapter5

import "fmt"

func Test(name string)  {
	fmt.Printf("name: %s\n", name)
}

func (u *User)SetMobile1(new_mobile string) {
	u.mobile = new_mobile
}