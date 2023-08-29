package main

import (
	"fmt"
	tinytime "github.com/wagslane/go-tinytime"
	"time"
)

func main() {
	tt := tinytime.New(1585750374)

	tt = tt.Add(time.Hour * 48)
	fmt.Println(tt)
}
