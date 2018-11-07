// Sample program to show how to use an interface in Go.
// go中接口的使用
package main

import (
	"fmt"
)

// notifier is an interface that defined notification
// type behavior.
type notifier interface {
	notify()
}

// user defines a user in the program.
type user struct {
	name  string
	email string
}

// notify implements a method with a pointer receiver.
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

// main is the entry point for the application.
func main() {
	// Create a value of type User and send a notification.
	//u := user{"Bill", "bill@email.com"}
	var u notifier = &user{"Bill", "bill@email.com"}

	//u.notify()
	sendNotification(u)

	// ./listing36.go:32: cannot use u (type user) as type
	//                     notifier in argument to sendNotification:
	//   user does not implement notifier
	//                          (notify method has pointer receiver)

	var c Person = &Chinese{}
	var e Person = &English{}

	JustSpeak(c)
	JustSpeak(e)
}

// sendNotification accepts values that implement the notifier
// interface and sends notifications.
func sendNotification(n notifier) {
	n.notify()
}

type Person interface {
	Speak()
}

type Chinese struct{}

func (p *Chinese)Speak()  {
	fmt.Println("你好，世界")
}

type English struct{}

func (p *English)Speak()  {
	fmt.Println("hello, world")
}

func JustSpeak(p Person)  {
	p.Speak()
}