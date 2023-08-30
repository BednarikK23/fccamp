package main

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

func main() {
	/*
			No need to hardly explain i had 3 courses included parallel programing and mutexes
			they are builtin primitives in standard library within the sync.Mutex
			allow lock and access to the data

			MAPS Are not thread safe

		The principle problem that mutexes help us avoid is the concurrent read/write problem.
		This problem arises when one thread is writing to a variable while another thread is reading from that same variable at the same time.
	*/
	mutMain()

	/*
		RW MUTEX - standard ibrary also exposes a sync.RWMutex
		in addition to Lock() and Unlock() RW mutex also has: RLock(), RUnlock()
		The sync.RWMutex can help with performance if we have a read-intensive process.
		Many goroutines can safely read from the map at the same time (multiple Rlock() calls can happen simultaneously)
		However, only one goroutine can hold a Lock() and all RLock()'s will also be excluded.
	*/
	rwMain()
}

type safeCounter struct {
	counts map[string]int
	mu     *sync.Mutex
}

func (sc safeCounter) inc(key string) {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.slowIncrement(key)
}

func (sc safeCounter) val(key string) int {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	return sc.counts[key]
}

func (sc safeCounter) slowIncrement(key string) {
	tempCounter := sc.counts[key]
	time.Sleep(time.Microsecond)
	tempCounter++
	sc.counts[key] = tempCounter
}

type emailTest struct {
	email string
	count int
}

func mutTest(sc safeCounter, emailTests []emailTest) {
	emails := make(map[string]struct{})

	var wg sync.WaitGroup
	for _, emailT := range emailTests {
		emails[emailT.email] = struct{}{}
		for i := 0; i < emailT.count; i++ {
			wg.Add(1)
			go func(emailT emailTest) {
				sc.inc(emailT.email)
				wg.Done()
			}(emailT)
		}
	}
	wg.Wait()

	emailsSorted := make([]string, 0, len(emails))
	for email := range emails {
		emailsSorted = append(emailsSorted, email)
	}
	sort.Strings(emailsSorted)

	for _, email := range emailsSorted {
		fmt.Printf("Email: %s has %d emails\n", email, sc.val(email))
	}
	fmt.Println("=====================================")
}

func mutMain() {
	/*
		We send emails across many different goroutines at Mailio.
		To keep track of how many we've sent to a given email address, we use an in-memory map.

		Our safeCounter struct is unsafe! Update the inc() and val() methods so that
		they utilize the safeCounter's mutex to ensure that the map is not accessed by multiple
		goroutines at the same time.
	*/

	sc := safeCounter{
		counts: make(map[string]int),
		mu:     &sync.Mutex{},
	}
	mutTest(sc, []emailTest{
		{
			email: "john@example.com",
			count: 23,
		},
		{
			email: "john@example.com",
			count: 29,
		},
		{
			email: "jill@example.com",
			count: 31,
		},
		{
			email: "jill@example.com",
			count: 67,
		},
	})
	mutTest(sc, []emailTest{
		{
			email: "kaden@example.com",
			count: 23,
		},
		{
			email: "george@example.com",
			count: 126,
		},
		{
			email: "kaden@example.com",
			count: 31,
		},
		{
			email: "george@example.com",
			count: 453,
		},
	})
}

// ---------------------------------------------------------------------

type safeC struct {
	counts map[string]int
	mu     *sync.RWMutex
}

func (sc safeC) inc1(key string) {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.slowIncrement1(key)
}

func (sc safeC) val1(key string) int {
	sc.mu.RLock()
	defer sc.mu.RUnlock()
	return sc.counts[key]
}

func (sc safeC) slowIncrement1(key string) {
	tempCounter := sc.counts[key]
	time.Sleep(time.Microsecond)
	tempCounter++
	sc.counts[key] = tempCounter
}

type emailT struct {
	email string
	count int
}

func rwTest(sc safeC, emailTests []emailT) {
	emails := make(map[string]struct{})

	var wg sync.WaitGroup
	for _, emailTe := range emailTests {
		emails[emailTe.email] = struct{}{}
		for i := 0; i < emailTe.count; i++ {
			wg.Add(1)
			go func(emailTe emailT) {
				sc.inc1(emailTe.email)
				wg.Done()
			}(emailTe)
		}
	}
	wg.Wait()

	emailsSorted := make([]string, 0, len(emails))
	for email := range emails {
		emailsSorted = append(emailsSorted, email)
	}
	sort.Strings(emailsSorted)

	sc.mu.RLock()
	defer sc.mu.RUnlock()
	for _, email := range emailsSorted {
		fmt.Printf("Email: %s has %d emails\n", email, sc.val1(email))
	}
	fmt.Println("=====================================")
}

func rwMain() {
	/*
		Let's update our same code from the last assignment, but this time
		we can speed it up by allowing readers to read from the map concurrently.

		Update the val method to only lock the mutex for reading.
		Notice that if you run the code with a write lock it will block forever.
	*/

	sc := safeC{
		counts: make(map[string]int),
		mu:     &sync.RWMutex{},
	}
	rwTest(sc, []emailT{
		{
			email: "john@example.com",
			count: 23,
		},
		{
			email: "john@example.com",
			count: 29,
		},
		{
			email: "jill@example.com",
			count: 31,
		},
		{
			email: "jill@example.com",
			count: 67,
		},
	})
	rwTest(sc, []emailT{
		{
			email: "kaden@example.com",
			count: 23,
		},
		{
			email: "george@example.com",
			count: 126,
		},
		{
			email: "kaden@example.com",
			count: 31,
		},
		{
			email: "george@example.com",
			count: 453,
		},
	})
}
