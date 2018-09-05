// Sample program to show how to declare methods and how the Go
// compiler supports them.
package main

import (
	"fmt"
	"time"
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

func (u user) change() {
	u.name = "gongyao"
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

	ab := lisa.changeName("gongyao")
	fmt.Println(ab)

	// Values of type user can be used to call methods
	// declared with a pointer receiver.
	//bill.changeEmail("bill@newdomain.com")
	//bill.notify()

	// Pointers of type user can be used to call methods
	// declared with a pointer receiver.
	//lisa.changeEmail("lisa@newdomain.com")
	//lisa.notify()
	fmt.Printf("%s", time.Now())
}

func (u *user)changeName(new_name string) (a string){
	u.name = new_name
	a = new_name

	return
}

func (ad admin) notify2() {
	fmt.Printf("name:%s, email:%s, age:%d\n", ad.people.name, ad.people.email, ad.age)
}
