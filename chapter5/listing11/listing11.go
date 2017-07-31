// Sample program to show how to declare methods and how the Go
// compiler supports them.
package main

import (
	"fmt"
)

// user defines a user in the program.
type user struct {
	name  string
	email string
}

type admin struct {
	people user
	age int
}

// notify implements a method with a value receiver.
func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n",
		u.name,
		u.email)
}

// changeEmail implements a method with a pointer receiver.
func (u *user) changeEmail(email string) {
	u.email = email
}

// main is the entry point for the application.
func main() {
	// Values of type user can be used to call methods
	// declared with a value receiver.
	bill := user{"Bill", "bill@email.com"}
	bill.notify()

	// Pointers of type user can also be used to call methods
	// declared with a value receiver.
	lisa := &user{"Lisa", "lisa@email.com"}
	lisa.notify()

	// Values of type user can be used to call methods
	// declared with a pointer receiver.
	bill.changeEmail("bill@newdomain.com")
	bill.notify()

	// Pointers of type user can be used to call methods
	// declared with a pointer receiver.
	lisa.changeEmail("lisa@newdomain.com")
	lisa.notify()

	//已经初始化了
	gyao := user{
		email:"gongyao1992@qq.com",
		name:"gongyao",
	}
	//gyao1 := user{"gongyao1", "gongyao1992@qq.com"}

	gyao.notify()
	gyao.changeName("wanghui")
	gyao.notify()
	//gyao1.notify()

	gyao2 := admin{
		people:gyao,
		age:25,
	}

	gyao2.people.notify()
	gyao2.people.changeName("yaoke")
	gyao2.people.notify()

	gyao3 := admin{
		people:user{
			name:"haonan",
			email:"heshun@qq.com",
		},
		age:21,
	}
	gyao3.notify2()
}

func (u *user)changeName(new_name string) {
	u.name = new_name
}

func (ad admin) notify2() {
	fmt.Printf("name:%s, email:%s, age:%d\n", ad.people.name, ad.people.email, ad.age)
}
