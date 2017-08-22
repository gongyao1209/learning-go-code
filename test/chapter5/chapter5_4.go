package chapter5

import "fmt"

type User struct {
	name string
	mobile string
}

func (u *User)SetName(new_name string) {
	u.name = new_name
}

func (u *User)SetMobile(new_mobile string) {
	u.mobile = new_mobile
}

func (u User)Show() {
	fmt.Printf("user name :%s, mobile:%s\n", u.name, u.mobile)
}

//func GetName(name string) {
//	return fmt.Sprintf("aaa%s", name)
//}

type AlertCounter int

// New creates and returns values of the unexported
// type alertCounter.
func New(value int) (a AlertCounter) {
	a = AlertCounter(value)
	return
}