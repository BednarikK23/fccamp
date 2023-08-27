package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
)

func main() {
	/*
		functions can be
		passed as an argument to other functions
		can be assigned to variable or even returned

		function that returns a function or accepts a function is called high-order Function

		- another way to think it is that function is just another type...

		example:
	*/
	fmt.Println(aggregate(2, 3, 4, add))
	fmt.Println(aggregate(2, 3, 4, mul))

	/*
		good uses:
		- HTTP API handlers
		- Pub/Sub handlers
		- Onclick callbacks
		Any time you need to run custom code at a time in the future, functions as values might make sense.
	*/

	/*
		CURRING
		- like special kind of high order function
		- takes a function and returns new function, its like
		- enhancing function of some ability
	*/
	// now the multiply functions transform to square
	// - it will be done two times because of selfMath
	squareFunc := selfMath(multiply)
	doubleFunc := selfMath(add)

	fmt.Println(squareFunc(5))
	// prints 25

	fmt.Println(doubleFunc(5))
	// prints 10
	loggerMain()

	/*
		DEFER
		- defer is concept specific in Go
		- allows a function to be executed automatically just before its enclosing function returns
		- the deferred call's arguments are evaluated immediately,
			but the function call is not executed until surrounding function returns

		- Deferred functions are typically used to close database connections, file handlers and the like.
		for example CopyFile() - look, actually really cool u ll not forget to close it...
	*/
	deferMain()

	/*
			CLOSURES
			A closure is a function that references variable from outside of its own function body
		The function may access and assign to the referenced variables

		In this example the concatter() function returns a function that has reference to an enclosed doc value
		Each successive call to harryPotterAggregator mutates that some doc variable
	*/
	closures()
	closureMain()

	/*
		ANONYMOUS FUNCTIONS - they have no name, wow
		We've been using them throughout this chapter, but we haven't really talked about them yet.

		Anonymous functions are useful when defining a function that will only be used once or to create a quick closure.
	*/
	nums := []int{1, 2, 3, 4, 5}

	// Here we define an anonymous function that doubles an int
	// and pass it to doMath
	allNumsDoubled := doMath(func(x int) int {
		return x + x
	}, nums)

	fmt.Println(allNumsDoubled)
	// prints:
	// [2 4 6 8 10]
	anonymousMain()
}

// doMath accepts a function that converts one int into another
// and a slice of ints. It returns a slice of ints that have been
// converted by the passed in function.
func doMath(f func(int) int, nums []int) []int {
	var results []int
	for _, n := range nums {
		results = append(results, f(n))
	}
	return results
}

func concatter() func(string) string {
	doc := ""
	return func(word string) string {
		doc += word + " "
		return doc
	}
}

func closures() {
	harryPotterAggregator := concatter()
	harryPotterAggregator("Mr.")
	harryPotterAggregator("and")
	harryPotterAggregator("Mrs.")
	harryPotterAggregator("Dursley")
	harryPotterAggregator("of")
	harryPotterAggregator("number")
	harryPotterAggregator("four,")
	harryPotterAggregator("Privet")

	fmt.Println(harryPotterAggregator("Drive"))
	// Mr. and Mrs. Dursley of number four, Privet Drive
}

// CopyFile copies a file from srcName to dstName on the local filesystem.
func CopyFile(dstName, srcName string) (written int64, err error) {

	// Open the source file
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	// Close the source file when the CopyFile function returns
	defer src.Close()

	// Create the destination file
	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	// Close the destination file when the CopyFile function returns
	defer dst.Close()

	return io.Copy(dst, src)
}

func multiply(x, y int) int {
	return x * y
}

func selfMath(mathFunc func(int, int) int) func(int) int {
	return func(x int) int {
		return mathFunc(x, x)
	}
}

func add(x, y int) int {
	return x + y
}

func mul(x, y int) int {
	return x * y
}

// aggregate applies the given math function to the first 3 inputs
func aggregate(a, b, c int, arithmetic func(int, int) int) int {
	return arithmetic(arithmetic(a, b), c)
}

// getLogger takes a function that formats two strings into
// a single string and returns a function that formats two strings but prints
// the result instead of returning it
func getLogger(formatter func(string, string) string) func(string, string) {
	return func(a, b string) {
		fmt.Println(formatter(a, b))
	}
}

func test(first string, errors []error, formatter func(string, string) string) {
	defer fmt.Println("====================================")
	logger := getLogger(formatter)
	fmt.Println("Logs:")
	for _, err := range errors {
		logger(first, err.Error())
	}
}

func colonDelimit(first, second string) string {
	return first + ": " + second
}
func commaDelimit(first, second string) string {
	return first + ", " + second
}

func loggerMain() {
	/*
		The Mailio API needs a very robust error-logging system so we can see when things are going
		awry in the back-end system. We need a function that can create a custom "logger" (a function that prints to the console) given a specific formatter.

		Complete the getLogger function. It should return a new function that prints the formatted inputs
		using the given formatter function. The inputs should be passed into the formatter function in the order they are given to the logger function.

		btw if u see tests u can see that they create two different adders
		and they are used for computation of different things
	*/
	dbErrors := []error{
		errors.New("out of memory"),
		errors.New("cpu is pegged"),
		errors.New("networking issue"),
		errors.New("invalid syntax"),
	}
	test("Error on database server", dbErrors, colonDelimit)

	mailErrors := []error{
		errors.New("email too large"),
		errors.New("non alphanumeric symbols found"),
	}
	test("Error on mail server", mailErrors, commaDelimit)
}

// ----------------------------------------------------------------------

const (
	logDeleted  = "user deleted"
	logNotFound = "user not found"
	logAdmin    = "admin deleted"
)

func logAndDelete(users map[string]user, name string) (log string) {
	// we could use delete() on 3 lines - before every return, or do this!! fantastic!
	defer delete(users, name)

	user, ok := users[name]
	if !ok {
		return logNotFound
	}
	if user.admin {
		return logAdmin
	}
	return logDeleted
}

type user struct {
	name   string
	number int
	admin  bool
}

func deferTest(users map[string]user, name string) {
	fmt.Printf("Attempting to delete %s...\n", name)
	defer fmt.Println("====================================")
	log := logAndDelete(users, name)
	fmt.Println("Log:", log)
}

func deferMain() {
	users := map[string]user{
		"john": {
			name:   "john",
			number: 18965554631,
			admin:  true,
		},
		"elon": {
			name:   "elon",
			number: 19875556452,
			admin:  true,
		},
		"breanna": {
			name:   "breanna",
			number: 98575554231,
			admin:  false,
		},
		"kade": {
			name:   "kade",
			number: 10765557221,
			admin:  false,
		},
	}

	fmt.Println("Initial users:")
	usersSorted := []string{}
	for name := range users {
		usersSorted = append(usersSorted, name)
	}
	sort.Strings(usersSorted)
	for _, name := range usersSorted {
		fmt.Println(" -", name)
	}
	fmt.Println("====================================")

	deferTest(users, "john")
	deferTest(users, "santa")
	deferTest(users, "kade")

	fmt.Println("Final users:")
	usersSorted = []string{}
	for name := range users {
		usersSorted = append(usersSorted, name)
	}
	sort.Strings(usersSorted)
	for _, name := range usersSorted {
		fmt.Println(" -", name)
	}
	fmt.Println("====================================")
}

// -----------------------------------------------------------------------------------

func adder() func(int) int {
	counter := 0
	return func(x int) int {
		counter += x
		return counter
	}

}

type emailBill struct {
	costInPennies int
}

func closureTest(bills []emailBill) {
	defer fmt.Println("====================================")
	countAdder, costAdder := adder(), adder()
	for _, bill := range bills {
		fmt.Printf("You've sent %d emails and it has cost you %d cents\n", countAdder(1), costAdder(bill.costInPennies))
	}
}

func closureMain() {
	/*
		Keeping track of how many emails we send is mission-critical at Mailio. Complete the adder() function.

		It should return a function that adds its input (an int) to an enclosed sum value, then return the new sum.
		In other words, it keeps a running total of the sum variable within a closure.
	*/
	closureTest([]emailBill{
		{45},
		{32},
		{43},
		{12},
		{34},
		{54},
	})

	closureTest([]emailBill{
		{12},
		{12},
		{976},
		{12},
		{543},
	})

	closureTest([]emailBill{
		{743},
		{13},
		{8},
	})
}

// ---------------------------------------------------------------------------

func printReports(messages []string) {
	for _, message := range messages {
		printCostReport(func(msg string) int {
			return len(msg) * 2
		}, message)
	}
}

func anonymousTest(messages []string) {
	defer fmt.Println("====================================")
	printReports(messages)
}

func anonymousMain() {
	/*
		ASSIGMENT:
		Complete the printReports function.

		Call printCostReport once for each message.
		Pass in an anonymous function as the costCalculator that returns an int equal to twice the length of the input message.
	*/
	anonymousTest([]string{
		"Here's Johnny!",
		"Go ahead, make my day",
		"You had me at hello",
		"There's no place like home",
	})

	anonymousTest([]string{
		"Hello, my name is Inigo Montoya. You killed my father. Prepare to die.",
		"May the Force be with you.",
		"Show me the money!",
		"Go ahead, make my day.",
	})
}

func printCostReport(costCalculator func(string) int, message string) {
	cost := costCalculator(message)
	fmt.Printf(`Message: "%s" Cost: %v cents`, message, cost)
	fmt.Println()
}
