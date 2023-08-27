package main

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"sort"
)

func main() {
	/*
		Pythons dictionaries are essentially maps in go

		O(1) lookup

		Maps are like slice passed by reference into function.
		-> when a map is passed int a function we can make changes into original in this function

		TYPES:
		- every type can be used as value
		- as a key has to by COMPARABLE TYPE, also cannot be: slice, map or function!

	*/
	// DECLARATION:
	// 1)
	ages := make(map[string]int) // empty map
	ages["Kiki"] = 21
	ages["Nina"] = 21
	ages["Neli"] = 23

	var ages2 = map[string]int{
		"John": 37,
		"Mary": 44,
	}
	var timeZone = map[string]int{
		"UTC": 0 * 60 * 60,
		"EST": -5 * 60 * 60,
		"CST": -6 * 60 * 60,
		"MST": -7 * 60 * 60,
		"PST": -8 * 60 * 60,
	}
	// LENGTH same as slices...
	fmt.Println(len(ages2), len(timeZone))

	// INSERT elem
	ages["Kuba"] = 5
	// GET elem
	kubasAge := ages["Kuba"]
	fmt.Println(kubasAge)
	getMain()

	// DELETE elem
	delete(ages, "Neli")
	deleteMain()

	// CHECK EXISTENCE
	// if key is in ages then ok is true else false
	// if key is not in ages then elem is zero value for elements type else mapped elem
	elem, ok := ages["Kiki"]
	fmt.Println(elem, ok)
	elem, ok = ages["Neli"]
	fmt.Println(elem, ok)

	// cool construct, if key doesn't exist, then initialize:
	names := map[string]int{}
	if _, ok := names["elon"]; !ok {
		// if the key doesn't exist yet,
		// initialize its value to 0
		names["elon"] = 0
	}
	countMain()

	// NESTED
	hits := make(map[string]map[string]int)
	// this is totally possible, but when adding data as for any given outer key
	// you must check if the inner map exists, and create it if needed:
	add(hits, "Czech", "moravia")

	// simplier way using struct Key With mapping path to country:
	c := make(map[Key]int)
	c[Key{"/", "vn"}]++
	nestedMain()

}

type Key struct {
	Path, Country string
}

func add(m map[string]map[string]int, path, country string) {
	// because go is panicking if it exists, so you have to check for each map,
	// and it resolves into extra code, but it is totally doable...:
	mm, ok := m[path]
	if !ok {
		mm = make(map[string]int)
		m[path] = mm
	}
	// because we made sure that inner map exist and if name doesn't yet exist,
	// it will return 0 tht we can increase and store 1, if it does exist we just increase value...
	mm[country]++
}

// ------------------------------------------------------------------

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	userMap := make(map[string]user)
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("invalid size")
	}

	for i := 0; i < len(names); i++ {
		userMap[names[i]] = user{
			name:        names[i],
			phoneNumber: phoneNumbers[i],
		}
	}
	return userMap, nil
}

type user struct {
	name        string
	phoneNumber int
}

func test(names []string, phoneNumbers []int) {
	fmt.Println("Creating map...")
	defer fmt.Println("====================================")
	users, err := getUserMap(names, phoneNumbers)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, name := range names {
		fmt.Printf("key: %v, value:\n", name)
		fmt.Println(" - name:", users[name].name)
		fmt.Println(" - number:", users[name].phoneNumber)
	}
}

func getMain() {
	/*
		We can speed up our contact-info lookups by using a map! Looking up a value in a map by its key is much faster than searching through a slice.

		Complete the getUserMap function. It takes a slice of names and a slice of phone numbers, and returns a map of name -> user structs and potentially an error. A user struct just contains a user's name and phone number.

		If the length of names and phoneNumbers is not equal, return an error with the string "invalid sizes".

		The first name in the names slice matches the first phone number, and so on.
	*/
	test(
		[]string{"John", "Bob", "Jill"},
		[]int{14355550987, 98765550987, 18265554567},
	)
	test(
		[]string{"John", "Bob"},
		[]int{14355550987, 98765550987, 18265554567},
	)
	test(
		[]string{"George", "Sally", "Rich", "Sue"},
		[]int{20955559812, 38385550982, 48265554567, 16045559873},
	)
}

// ------------------------------------------------------------------

func deleteIfNecessary(users map[string]user2, name string) (deleted bool, err error) {
	user, ok := users[name]
	if !ok {
		return false, errors.New("not found")
	}
	if user.scheduledForDeletion {
		delete(users, name)
		return true, nil
	}
	return false, nil
}

type user2 struct {
	name                 string
	number               int
	scheduledForDeletion bool
}

func deleteTest(users map[string]user2, name string) {
	fmt.Printf("Attempting to delete %s...\n", name)
	defer fmt.Println("====================================")
	deleted, err := deleteIfNecessary(users, name)
	if err != nil {
		fmt.Println(err)
		return
	}
	if deleted {
		fmt.Println("Deleted:", name)
		return
	}
	fmt.Println("Did not delete:", name)
}

func deleteMain() {
	/*
		It's important to keep up with privacy regulations and to respect our user's data. We need a function that will delete user records.

		Complete the deleteIfNecessary function.

		If the user doesn't exist in the map, return the error not found.
		If they exist but aren't scheduled for deletion, return deleted as false with no errors.
		If they exist and are scheduled for deletion, return deleted as true with no errors and delete their record from the map.
	*/
	users := map[string]user2{
		"john": {
			name:                 "john",
			number:               18965554631,
			scheduledForDeletion: true,
		},
		"elon": {
			name:                 "elon",
			number:               19875556452,
			scheduledForDeletion: true,
		},
		"breanna": {
			name:                 "breanna",
			number:               98575554231,
			scheduledForDeletion: false,
		},
		"kade": {
			name:                 "kade",
			number:               10765557221,
			scheduledForDeletion: false,
		},
	}
	deleteTest(users, "john")
	deleteTest(users, "musk")
	deleteTest(users, "santa")
	deleteTest(users, "kade")

	keys := []string{}
	for name := range users {
		keys = append(keys, name)
	}
	sort.Strings(keys)

	fmt.Println("Final map keys:")
	for _, name := range keys {
		fmt.Println(" - ", name)
	}
}

// ----------------------------------------------------------------------

func getCounts(userIDs []string) map[string]int {
	userCount := make(map[string]int)
	for _, id := range userIDs {
		if _, ok := userCount[id]; !ok {
			// if the key doesn't exist yet,
			// initialize its value to 0
			userCount[id] = 0
		}
		userCount[id]++
	}
	return userCount
}

func countTest(userIDs []string, ids []string) {
	fmt.Printf("Generating counts for %v user IDs...\n", len(userIDs))

	counts := getCounts(userIDs)
	fmt.Println("Counts from select IDs:")
	for _, k := range ids {
		v := counts[k]
		fmt.Printf(" - %s: %d\n", k, v)
	}
	fmt.Println("=====================================")
}

func countMain() {
	/*
		We have a slice of user ids, and each instance of an id in the slice indicates that a message was sent to that user. We need to count up how many times each user's id appears in the slice to track how many messages they received.

		Implement the getCounts function. It should return a map of string -> int so that each int is a count of how many times each string was found in the slice.
	*/
	userIDs := []string{}
	for i := 0; i < 10000; i++ {
		h := md5.New()
		io.WriteString(h, fmt.Sprint(i))
		key := fmt.Sprintf("%x", h.Sum(nil))
		userIDs = append(userIDs, key[:2])
	}

	countTest(userIDs, []string{"00", "ff", "dd"})
	countTest(userIDs, []string{"aa", "12", "32"})
	countTest(userIDs, []string{"bb", "33"})
}

// ----------------------------------------------------------------------

func addName(m map[rune]map[string]int, char rune, name string) {
	if _, ok := m[char]; !ok {
		m[char] = make(map[string]int, 4)
	}
	// because we made sure that inner map exist and if name doesn't yet exist,
	// it will return 0 tht we can increase and store 1, if it does exist we just increase value...
	m[char][name]++
}

func getNameCounts(names []string) map[rune]map[string]int {
	m := make(map[rune]map[string]int)
	for _, name := range names {
		if name == "" {
			continue
		}

		addName(m, rune(name[0]), name)
	}
	return m
}

// don't edit below this line

func nestedTest(names []string, initial rune, name string) {
	fmt.Printf("Generating counts for %v names...\n", len(names))

	nameCounts := getNameCounts(names)
	count := nameCounts[initial][name]
	fmt.Printf("Count for [%c][%s]: %d\n", initial, name, count)
	fmt.Println("=====================================")
}

func nestedMain() {
	/*
		Because Textio is a glorified customer database, we have a lot of internal logic for sorting and dealing with customer names.

		Complete the getNameCounts function. It takes a slice of strings (names) and returns a nested map where the first key is
		all the unique first characters of the names, the second key is all the names themselves, and the value is the count of each name.
	*/
	nestedTest(getNames(50), 'M', "Matthew")
	nestedTest(getNames(100), 'G', "George")
	nestedTest(getNames(150), 'D', "Drew")
	nestedTest(getNames(200), 'P', "Philip")
	nestedTest(getNames(250), 'B', "Bryant")
	nestedTest(getNames(300), 'M', "Matthew")
}

func getNames(length int) []string {
	names := []string{
		"Grant", "Eduardo", "Peter", "Matthew", "Matthew", "Matthew", "Peter", "Peter", "Henry", "Parker", "Parker", "Parker", "Collin", "Hayden", "George", "Bradley", "Mitchell", "Devon", "Ricardo", "Shawn", "Taylor", "Nicolas", "Gregory", "Francisco", "Liam", "Kaleb", "Preston", "Erik", "Alexis", "Owen", "Omar", "Diego", "Dustin", "Corey", "Fernando", "Clayton", "Carter", "Ivan", "Jaden", "Javier", "Alec", "Johnathan", "Scott", "Manuel", "Cristian", "Alan", "Raymond", "Brett", "Max", "Drew", "Andres", "Gage", "Mario", "Dawson", "Dillon", "Cesar", "Wesley", "Levi", "Jakob", "Chandler", "Martin", "Malik", "Edgar", "Sergio", "Trenton", "Josiah", "Nolan", "Marco", "Drew", "Peyton", "Harrison", "Drew", "Hector", "Micah", "Roberto", "Drew", "Brady", "Erick", "Conner", "Jonah", "Casey", "Jayden", "Edwin", "Emmanuel", "Andre", "Phillip", "Brayden", "Landon", "Giovanni", "Bailey", "Ronald", "Braden", "Damian", "Donovan", "Ruben", "Frank", "Gerardo", "Pedro", "Andy", "Chance", "Abraham", "Calvin", "Trey", "Cade", "Donald", "Derrick", "Payton", "Darius", "Enrique", "Keith", "Raul", "Jaylen", "Troy", "Jonathon", "Cory", "Marc", "Eli", "Skyler", "Rafael", "Trent", "Griffin", "Colby", "Johnny", "Chad", "Armando", "Kobe", "Caden", "Marcos", "Cooper", "Elias", "Brenden", "Israel", "Avery", "Zane", "Zane", "Zane", "Zane", "Dante", "Josue", "Zackary", "Allen", "Philip", "Mathew", "Dennis", "Leonardo", "Ashton", "Philip", "Philip", "Philip", "Julio", "Miles", "Damien", "Ty", "Gustavo", "Drake", "Jaime", "Simon", "Jerry", "Curtis", "Kameron", "Lance", "Brock", "Bryson", "Alberto", "Dominick", "Jimmy", "Kaden", "Douglas", "Gary", "Brennan", "Zachery", "Randy", "Louis", "Larry", "Nickolas", "Albert", "Tony", "Fabian", "Keegan", "Saul", "Danny", "Tucker", "Myles", "Damon", "Arturo", "Corbin", "Deandre", "Ricky", "Kristopher", "Lane", "Pablo", "Darren", "Jarrett", "Zion", "Alfredo", "Micheal", "Angelo", "Carl", "Oliver", "Kyler", "Tommy", "Walter", "Dallas", "Jace", "Quinn", "Theodore", "Grayson", "Lorenzo", "Joe", "Arthur", "Bryant", "Roman", "Brent", "Russell", "Ramon", "Lawrence", "Moises", "Aiden", "Quentin", "Jay", "Tyrese", "Tristen", "Emanuel", "Salvador", "Terry", "Morgan", "Jeffery", "Esteban", "Tyson", "Braxton", "Branden", "Marvin", "Brody", "Craig", "Ismael", "Rodney", "Isiah", "Marshall", "Maurice", "Ernesto", "Emilio", "Brendon", "Kody", "Eddie", "Malachi", "Abel", "Keaton", "Jon", "Shaun", "Skylar", "Ezekiel", "Nikolas", "Santiago", "Kendall", "Axel", "Camden", "Trevon", "Bobby", "Conor", "Jamal", "Lukas", "Malcolm", "Zackery", "Jayson", "Javon", "Roger", "Reginald", "Zachariah", "Desmond", "Felix", "Johnathon", "Dean", "Quinton", "Ali", "Davis", "Gerald", "Rodrigo", "Demetrius", "Billy", "Rene", "Reece", "Kelvin", "Leo", "Justice", "Chris", "Guillermo", "Matthew", "Matthew", "Matthew", "Kevon", "Steve", "Frederick", "Clay", "Weston", "Dorian", "Hugo", "Roy", "Orlando", "Terrance", "Kai", "Khalil", "Khalil", "Khalil", "Graham", "Noel", "Willie", "Nathanael", "Terrell", "Tyrone",
	}
	if length > len(names) {
		length = len(names)
	}
	return names[:length]
}
