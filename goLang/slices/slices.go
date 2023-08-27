package main

import "fmt"

func main() {
	/*
	   ARRAYS
	    - are always fixed
	*/

	// var myInts [5]int  ...
	messages := [3]string{
		"Hello",
		"I want coffee",
		"Hehe",
	}

	primes := [6]int{2, 3, 5, 7, 11, 13}
	stringIt(primes[:]) // could not use primes only - its array, need to use slice...

	/*
		SLICES
		- we can slice array, same as in python
		- in go most common, built on top of arrays but working mainly with slices in go...
		- slices hold an underlying array, if you assign one slice to another both refer to same array
		- function taking slice argument, changes it makes to the element of the slice WILL BE VISIBLE TO THE CALLER
		-  => like pointers, its better to pass slice then array + count...
		- [included:not-included]
		- [:] => all, basically copy...
	*/
	slice := primes[1:4]
	stringIt(slice)
	/*
		SLICE CREATION
		- how to make slice without creating array explicitly (it will create array under the hood for us):
		make takes:
			- types of elems
			- current length
			- capcity to what slice can grow..., this argument is optional, when not set capacity = length
				- if cap exceeds the slice is reallocated...
		- functions:
			len - returns current length of slice/array
			cap - returns allocated capacity of the array
			cap, len - when slice IS NIL RETURNS 0
	*/
	costs := make([]float64, 5, 10) // type, length, capacity

	for i := 0; i < len(messages); i++ {
		cost := float64(len(messages[i])) * 0.01
		costs[i] = cost * 0.01
	}

	total := summ(1, 2, 3)
	fmt.Println(total)

	// if you have real slice and want to use spread function we can use this
	// inverse spread operator to pass this slice into function
	names := []string{"bob", "sue"}
	printStrings(names...)

	/*
		APPEND - built-in function to dynamically add elements to a slice
		func append(slice []Type, elems ...Type) []Type
		if the underlying capacity is not large enough, append() will create new underlying array and point the slice to it

		NOT TO DO!!!
		- slice2 = append(slice1, something)
		- always save append into same slice like so:
		- slice1 = append(slice1, something)
		- if append reallocate slice then slice1 and slice2 will have different address
		- else slice2 and slice1 have same address, and they can override each other and its just headache
		- ALWAYS assign the result of append function to the same slice!
	*/
	names = append(names, "Kiki")
	names = append(names, "Neli", "Nina")
	anotherNames := []string{"Kubo", "Drla"}
	names = append(names, anotherNames...)

	printStrings(names...)
	/*
		SLICE OF SLICES

	*/
	printMatrix(createMatrix(3, 3))
	printMatrix(createMatrix(10, 10))

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

func createMatrix(rows, cols int) [][]int {
	matrix := make([][]int, 0)
	for i := 0; i < rows; i++ {
		row := make([]int, 0)
		for j := 0; j < cols; j++ {
			row = append(row, i*j)
		}
		matrix = append(matrix, row)
	}
	return matrix
}

func printMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Println()
}

func summ(nums ...int) int {
	/*
		SPREAD OPERATOR
		variadic function definition
		in function nums ...int works like normal slice, the difference is on callers side
		where he can give the functions as many arguments as he wishes...
	*/
	res := 0
	for i := 0; i < len(nums); i++ {
		res += nums[i]
	}
	return res
}

func printStrings(strings ...string) {
	for i := 0; i < len(strings); i++ {
		fmt.Println(strings[i])
	}
}

func stringIt(nums []int) {
	for i := 0; i < len(nums); i++ {
		println(nums[i])
	}
}
