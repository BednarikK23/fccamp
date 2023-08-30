package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	/*
		As we've mentioned, Go does not support classes.
		For a long time, that meant that Go code couldn't easily be reused in many circumstances.
		For example, imagine some code that splits a slice into 2 equal parts.
		The code that splits the slice doesn't really care about the values stored in the slice.
		Unfortunately in Go we would need to write it multiple times for each type, which is a very un-DRY thing to do.
		splitIntSlice(), splitStringSlice()
		In Go 1.20 however, support for generics was released, effectively solving this problem!

		TYPE PARAMETERS
		Put simply, generics allow us to use variables to refer to specific types.
		This is an amazing feature because it allows us to write abstract functions that drastically reduce code duplication.
		func splitAnySlice[T any](s []T) ([]T, []T)
		In the example above, T is the name of the type parameter for the splitAnySlice function,
		and we've said that it must match the "any" constraint, which means it can be anything.
		This makes sense because the body of the function doesn't care about the types of things stored in the slice.
	*/
	firstInts, secondInts := splitAnySlice([]int{0, 1, 2, 3})
	fmt.Println(firstInts, secondInts)

	/*
		way we can get 0 value of any type to return:
		if l == 0 {
			var zeroVal T
			return zeroVal
		}
	*/
	fstMain()

	// WHY:
	// GENERICS REDUCE REPETITIVE CODE
	// GENERICS ARE USED MORE OFTEN IN LIBRARIES AND PACKAGES

	/*
		CONSTRAINS
		Constraints are just interfaces that allow us to write generics
		that only operate within the constraint of a given interface type.
		Sometimes you need the logic in your generic function to know something about the types it operates on
		if we don't need to know anything about type like in exercse with slices,
		we can use builtin constraint "any". - any is the same as the empty interface

		CREATING A CUSTOM CONSTRAINT
		example of a concat function - takes a slice of values and concatenates the values into a string
		his should work with any type that can represent itself as a string, even if it's not a string under the hood
	*/
	constMain()

	/*
		INTERFACE TYPE LIST
		When generics were released, a new way of writing interfaces was also released at the same time!
		We can now simply list a bunch of types to get a new interface/constraint.:
		See Order interface
	*/
	/*
		PARAMETRIC CONSTRAINTS
		Your interface definitions, which can later be used as constraints, can accept type parameters as well.:
		See store interface
	*/
	billerMain()
}

type store[P product] interface {
	Sell(P)
}

// Ordered is a type constraint that matches any ordered type.
// An ordered type is one that supports the <, <=, >, and >= operators.
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
	~float32 | ~float64 |
	~string
}

type product interface {
	Price() float64
	Name() string
}

func splitIntSlice(s []int) ([]int, []int) {
	mid := len(s) / 2
	return s[:mid], s[mid:]
}

func splitStringSlice(s []string) ([]string, []string) {
	mid := len(s) / 2
	return s[:mid], s[mid:]
}

func splitAnySlice[T any](s []T) ([]T, []T) {
	mid := len(s) / 2
	return s[:mid], s[mid:]
}

type stringer interface {
	String() string
}

func concat[T stringer](vals []T) string {
	result := ""
	for _, val := range vals {
		// this is where the .String() method
		// is used. That's why we need a more specific
		// constraint instead of the "any" constraint
		result += val.String()
	}
	return result
}

// ------------------------------------------------------------------

func getLast[T any](a []T) T {
	l := len(a)
	if l == 0 {
		var zeroVal T
		return zeroVal
	}
	return a[len(a)-1]
}

type email struct {
	message        string
	senderEmail    string
	recipientEmail string
}

type payment struct {
	amount         int
	senderEmail    string
	recipientEmail string
}

func fstMain() {
	test([]email{}, "email")
	test([]email{
		{
			"Hi Margo",
			"janet@example.com",
			"margo@example.com",
		},
		{
			"Hey Margo I really wanna chat",
			"janet@example.com",
			"margo@example.com",
		},
		{
			"ANSWER ME",
			"janet@example.com",
			"margo@example.com",
		},
	}, "email")
	test([]payment{
		{
			5,
			"jane@example.com",
			"sally@example.com",
		},
		{
			25,
			"jane@example.com",
			"mark@example.com",
		},
		{
			1,
			"jane@example.com",
			"sally@example.com",
		},
		{
			16,
			"jane@example.com",
			"margo@example.com",
		},
	}, "payment")
}

func test[T any](s []T, desc string) {
	last := getLast(s)
	fmt.Printf("Getting last %v from slice of length: %v\n", desc, len(s))
	for i, v := range s {
		fmt.Printf("Item #%v: %v\n", i+1, v)
	}
	fmt.Printf("Last item in list: %v\n", last)
	fmt.Println(" --- ")
}

// -------------------------------------------------------------

func chargeForLineItem[T lineItem](newItem T, oldItems []T, balance float64) ([]T, float64, error) {
	if balance < newItem.GetCost() {
		return nil, 0.0, errors.New("insufficient funds")
	}
	oldItems = append(oldItems, newItem)
	return oldItems, balance - newItem.GetCost(), nil
}

type lineItem interface {
	GetCost() float64
	GetName() string
}

type subscription struct {
	userEmail string
	startDate time.Time
	interval  string
}

func (s subscription) GetName() string {
	return fmt.Sprintf("%s subscription", s.interval)
}

func (s subscription) GetCost() float64 {
	if s.interval == "monthly" {
		return 25.00
	}
	if s.interval == "yearly" {
		return 250.00
	}
	return 0.0
}

type oneTimeUsagePlan struct {
	userEmail        string
	numEmailsAllowed int
}

func (otup oneTimeUsagePlan) GetName() string {
	return fmt.Sprintf("one time usage plan with %v emails", otup.numEmailsAllowed)
}

func (otup oneTimeUsagePlan) GetCost() float64 {
	const costPerEmail = 0.03
	return float64(otup.numEmailsAllowed) * costPerEmail
}

func constMain() {
	/*
		Complete the chargeForLineItem function. First, it should check if the user has a balance with enough funds to be able to pay for the cost of the newItem.
		If they don't then return an "insufficient funds" error.

		If they do have enough funds:
		- Add the line item to the user's history by appending the newItem to the slice of oldItems. This new slice is your first return value.
		- Calculate the user's new balance by subtracting the cost of the new item from their balance. This is your second return value.
	*/
	constTest(subscription{
		userEmail: "john@example.com",
		startDate: time.Now().UTC(),
		interval:  "yearly",
	},
		[]subscription{},
		1000.00,
	)
	constTest(subscription{
		userEmail: "jane@example.com",
		startDate: time.Now().UTC(),
		interval:  "monthly",
	},
		[]subscription{
			{
				userEmail: "jane@example.com",
				startDate: time.Now().UTC().Add(-time.Hour * 24 * 7),
				interval:  "monthly",
			},
			{
				userEmail: "jane@example.com",
				startDate: time.Now().UTC().Add(-time.Hour * 24 * 7 * 52 * 2),
				interval:  "yearly",
			},
		},
		686.20,
	)
	constTest(oneTimeUsagePlan{
		userEmail:        "dillon@example.com",
		numEmailsAllowed: 5000,
	},
		[]oneTimeUsagePlan{},
		756.20,
	)
	constTest(oneTimeUsagePlan{
		userEmail:        "dalton@example.com",
		numEmailsAllowed: 100000,
	},
		[]oneTimeUsagePlan{
			{
				userEmail:        "dalton@example.com",
				numEmailsAllowed: 34200,
			},
		},
		32.20,
	)
}

func constTest[T lineItem](newItem T, oldItems []T, balance float64) {
	fmt.Println(" --- ")
	fmt.Printf("Charging customer for a '%s', current balance is %v...\n", newItem.GetName(), balance)
	newItems, newBalance, err := chargeForLineItem(newItem, oldItems, balance)
	if err != nil {
		fmt.Printf("Got error: %v\n", err)
		return
	}
	fmt.Printf("New balance is: %v. Total number of line items is now %v\n", newBalance, len(newItems))
}

// -----------------------------------------------------------------------------

type biller[C customer] interface {
	Charge(C) bill
	Name() string
}

type userBiller struct {
	Plan string
}

func (ub userBiller) Charge(u user) bill {
	amount := 50.0
	if ub.Plan == "pro" {
		amount = 100.0
	}
	return bill{
		Customer: u,
		Amount:   amount,
	}
}

func (sb userBiller) Name() string {
	return fmt.Sprintf("%s user biller", sb.Plan)
}

type orgBiller struct {
	Plan string
}

func (ob orgBiller) Name() string {
	return fmt.Sprintf("%s org biller", ob.Plan)
}

func (ob orgBiller) Charge(o org) bill {
	amount := 2000.0
	if ob.Plan == "pro" {
		amount = 3000.0
	}
	return bill{
		Customer: o,
		Amount:   amount,
	}
}

type customer interface {
	GetBillingEmail() string
}

type bill struct {
	Customer customer
	Amount   float64
}

type user struct {
	UserEmail string
}

func (u user) GetBillingEmail() string {
	return u.UserEmail
}

type org struct {
	Admin user
	Name  string
}

func (o org) GetBillingEmail() string {
	return o.Admin.GetBillingEmail()
}

func billerMain() {
	/*
		There are two kinds of billers:

		- userBiller (cheaper)
		- orgBiller (more expensive)
		A customer is either a user or an org. A user will be billed with a userBiller and an org with an orgBiller.

		Create the new biller interface. It should have 2 methods:

		- Charge
		- Name
	*/
	testBiller[user](
		userBiller{Plan: "basic"},
		user{UserEmail: "joe@example.com"},
	)
	testBiller[user](
		userBiller{Plan: "basic"},
		user{UserEmail: "samuel.boggs@example.com"},
	)
	testBiller[user](
		userBiller{Plan: "pro"},
		user{UserEmail: "jade.row@example.com"},
	)
	testBiller[org](
		orgBiller{Plan: "basic"},
		org{Admin: user{UserEmail: "challis.rane@example.com"}},
	)
	testBiller[org](
		orgBiller{Plan: "pro"},
		org{Admin: user{UserEmail: "challis.rane@example.com"}},
	)
}

func testBiller[C customer](b biller[C], c C) {
	fmt.Printf("Using '%s' to create a bill for '%s'\n", b.Name(), c.GetBillingEmail())
	bill := b.Charge(c)
	fmt.Printf("Bill created for %v dollars\n", bill.Amount)
	fmt.Println(" --- ")
}
