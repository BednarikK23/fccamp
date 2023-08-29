package main

// https://pkg.go.dev/strings#ReplaceAll
import (
	"fmt"
	"strings"
)

func main() {
	/*
		Pointer allow us to manipulate data in memory directly without making duplicate data
		thanks to it, we can change value inside function
	*/

	// INIT:
	// * defines pointer, zero value for pointer is nil
	var p *string
	myString := "hello"
	p = &myString
	fmt.Println(p)

	// * also dereference pointer
	fmt.Println(*p, p)
	// NIL
	/*
		if pointer points to nothing then dereferencing it will cause runtime error - check
	*/
	remMain()

	/*
		POINTER RECEIVERS - A receiver type on a method can be a pointer
		- very frequent on methods because methods will be doing changes into the object...

	*/
	c := car{
		color: "white",
	}
	c.setColor("blue")
	fmt.Println(c.color)

	recMain()
}

type car struct {
	color string
}

func (c *car) setColor(color string) {
	c.color = color
}

func removeProfanity(message *string) {
	if message == nil {
		return
	}
	mess := *message
	mess = strings.ReplaceAll(mess, "dang", "****")
	mess = strings.ReplaceAll(mess, "shoot", "*****")
	mess = strings.ReplaceAll(mess, "hack", "****")
	*message = mess
}

func remTest(messages []string) {
	for _, message := range messages {
		removeProfanity(&message)
		fmt.Println(message)
	}
}

func remMain() {

	/*
		emoveProfanity function.

		It should use the strings.ReplaceAll function to replace all instances of the following words in the input message with asterisks.

		"dang" -> "****"
		"shoot" -> "*****"
		"heck" -> "****"
		It should mutate the value in the pointer and return nothing.

		Do not alter the function signature.
	*/
	messages1 := []string{
		"well shoot, this is awful",
		"dang robots",
		"dang them to heck",
	}

	messages2 := []string{
		"well shoot",
		"Allan is going straight to heck",
		"dang... that's a tough break",
	}

	remTest(messages1)
	remTest(messages2)
}

// ------------------------------------------------------------------------------

func (e *email) setMessage(newMessage string) {
	e.message = newMessage
}

type email struct {
	message     string
	fromAddress string
	toAddress   string
}

func test(e *email, newMessage string) {
	fmt.Println("-- before --")
	e.print()
	fmt.Println("-- end before --")
	e.setMessage(newMessage)
	fmt.Println("-- after --")
	e.print()
	fmt.Println("-- end after --")
	fmt.Println("==========================")
}

func (e email) print() {
	fmt.Println("message:", e.message)
	fmt.Println("fromAddress:", e.fromAddress)
	fmt.Println("toAddress:", e.toAddress)
}

func recMain() {
	/*
		Fix the bug in the code so that setMessage sets the message field of the given email structure,
		and the new value persists outside the scope of the setMessage method.
	*/
	test(&email{
		message:     "this is my first draft",
		fromAddress: "sandra@mailio-test.com",
		toAddress:   "bullock@mailio-test.com",
	}, "this is my second draft")

	test(&email{
		message:     "this is my third draft",
		fromAddress: "sandra@mailio-test.com",
		toAddress:   "bullock@mailio-test.com",
	}, "this is my fourth draft")

}
