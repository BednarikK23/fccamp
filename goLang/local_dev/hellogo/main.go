package main

import (
	"fmt"
	// dont forget to modify your .mod file
	// "{REMOTE}/{USERNAME}/mystrings"
	"github.com/BednarikK23/fccamp/tree/master/goLang/local_dev/mystrings"
)

func main() {
	fmt.Println("hello world")
	fmt.Println(
		mystrings.Reverse("hello world"))
}
