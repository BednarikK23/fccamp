package main

import "fmt"

func bulkSend(numMessages int) float64 {
	totalCost := 0.0
	for i := 0; i < numMessages; i++ {
		totalCost += 1.0 + (0.1 * float64(i))
	}
	return totalCost
}

func main() {
	/*

			for INITIAL; CONDITION; AFTER {}

			all of the sections Initial, condition and after ar optional,
			sou you can omit them wiithout a problem like so:

			for INITIAL ; ; AFTER {} - this will run for ever

			- most languages have while loop, because while loop is just for loop without init and after,
			 	it is no need to have while loop in go so while loop in go are for loops like so:
			for x < y {}

		as usually, go provides continue and brake


	*/

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// 2
	totalCost, thresh := 0.0, 0.0
	for i := 0; ; i++ {
		totalCost += 1.0 + 0.1*float64(i)
		if totalCost > thresh {
			break
		}
	}

	// while loop:
	getMaxMessages(1.1, 100)

	/*
		RANGE
		GO provides syntactic sugar to iterate easily over elements of slice:
		for INDEX, ELEMENT := range SLICE {}
		- same as enumerate in python
	*/
	fruits := []string{"apple", "banana", "grape"}
	for i, fruit := range fruits {
		fmt.Println(i, fruit)
	}

}

func getMaxMessages(costMultiplier float64, maxCost int) int {
	actualCostPennies := 1.0
	maxMessagesToSend := 0
	for actualCostPennies <= float64(maxCost) {
		maxMessagesToSend++
		actualCostPennies *= costMultiplier
	}
	return maxMessagesToSend
}
