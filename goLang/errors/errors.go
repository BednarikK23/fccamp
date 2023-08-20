package main

import (
	"errors"
	"fmt"
)

// Go programs express errors with error values.
// Error is any type that implements the simple built-in error interface:
/*
type error interface {
	Error() string
}
*/

// when something can go wrong in a func, that func should return an error as its last return value.
// Any code that calls this func should check if error is nil.

// as you can see advantage is that we can save some data about error...
type userError struct {
	id   int
	name string
	code int
}

// ...and then just implement one stupid function..:
func (e userError) Error() string {
	return fmt.Sprintf("%s (%d) has problem with ther account, error: %d", e.name, e.id, e.code)
}

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (float64, error) {
	cost, err := sendSMS(msgToCustomer)
	if err != nil {
		return 0.0, err
	}
	costForSpouse, err := sendSMS(msgToSpouse)
	if err != nil {
		return 0.0, err
	}
	return cost + costForSpouse, nil
}

func sendSMS(message string) (float64, error) {
	const maxLen = 25
	const costPerChar = .0002
	if len(message) > maxLen {
		// Hope u didnt forget what sprintf does
		errMsg := fmt.Sprintf("cannot send text over %d characters long", maxLen)
		// if you dont wanna implement your interface:
		return 0.0, fmt.Errorf("%s", errMsg)
	}
	return costPerChar * float64(len(message)), nil
}

func divide(dividend, divisor int) (int, error) { // error is special interface
	if divisor == 0 {
		// errors package makes it easy to deal with errors...
		return 0, errors.New("Cannot devide by zero")
	}
	return dividend / divisor, nil // -> nil means no error
}

func main() {
	costs, err := sendSMSToCouple("lorem 1", "lorem2")
	if err != nil {
		return
	}
	fmt.Println("It cost us: ", costs)
}
