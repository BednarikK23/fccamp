package main

// see for more:
/// https://dave.cheney.net/2014/03/19/channel-axioms

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*
		Go was designed to be concurrent, which is a trait fairly unique to Go.
		It excels at performing many tasks simultaneously safely using a simple syntax.

		There isn't a popular programming language in existence where
		spawning concurrent execution is quite as elegant, at least in my opinion.

		Concurrency is as simple as using the go keyword when calling a function:
		go doSomething()
	*/
	mailMain()

	/*
		CHANNELS
		- are a typed, thread-safe queue.
		Channels allow different goroutines to communicate with each other.

		-Init
		Like maps and slices, channels must be created before use.
		They also use the same make keyword:
	*/
	// ch := make(chan int)

	// SEND DATA TO A CHANNEL:
	// ch <- 69
	// <- is channel operator, data flows in direction of the arrow.
	// <- IS BLOCKING device, will block until goroutine is ready to receive the data...

	// RECEIVING DATA FROM CHANNEL:
	// v := <-ch
	// This reads and removes a value from the channel and saves it into the variable v.
	// This operation will block until there is a value in the channel to be read

	filerMain()

	// Empty structs are often used as tokens in Go programs.
	// In this context, a token is a unary value.
	// In other words, we don't care what is passed through the channel.
	// We care when and if it is passed.
	// We can block and wait until something is sent on a channel using the following syntax
	// <-ch
	// This will block until it pops a single item off the channel, then continue, discarding the item.

	waitMain()

	// BUFFERED CHANNELS - channels can be optionally buffered
	// CREATING CHENNEL WITH BUFFER
	// ch := make(chan int, 10)
	// sending on buffered channel only blocks if the buffer is full
	// receiving block only if buffer is empty
	buffMain()

	// CLOSING CHANNELS IN GO
	// Channels can be explicitly closed by a sender:
	ch := make(chan int)
	// ...
	close(ch)
	_, ok := <-ch
	// CHECKING IF A CHANNEL IS CLOSED - used by reading side
	// Similar to the ok value when accessing data in a map, receivers can check
	// the ok value when receiving from a channel to test if a channel was closed.
	if ok { // ok is false if the channel is empty and closed.
		close(ch)
	}

	// DON'T SEND ON A CLOSED CHANNEL
	// - it will cause panic of this goroutine,
	// that's why only SENDING SIDE should CLOSE channel
	// Closing isn't necessary. if left channels open, they'll still be garbage collected if they're unused.
	// You should close channels to indicate explicitly to a receiver that nothing else is going to come across.
	closeMain()

	// RANGE
	// similar to slices and maps, channels can be ranged over:
	// for item := range ch { }
	//  -> receives values over the channel (blocking), will exit when channel is closed

	fibMain()

	// SELECT - similar to switch, but uniq for channels
	// - used to listen to multiple channels at the same time
	// the first channel with value ready to be received will fire and its body will execute
	// If multiple channels are ready at the same time one is chosen randomly.
	// in ok we store whether, or not channel been closed by sender yet
	/*
		select {
		case i, ok := <-chInts:
			fmt.Println(i)
		case s, ok := <-chStrings:
			fmt.Println(s)
		default:
			fmt.Println(shit)
		}
	*/
	logMain()

	// DEFAULT CASE:
	// executes immediately if no other channel has a value ready.
	// A default case stops the select statement from blocking.
	/*
		select {
		  case v := <-ch:
		    // use v
		  default:
		    // receiving from ch would block
		    // so do something else
		}
	*/

	/*
		TICKERS
		time.Tick() is a standard library function that returns a channel that sends a value on a given interval.
		time.After() sends a value once after the duration has passed.
		time.Sleep() blocks the current goroutine for the specified amount of time.
	*/

	/*
		READ-ONLY CHANNELS
		A channel can be marked as read-only by casting it from a chan to a <-chan type. For example:
	*/
	c := make(chan int)
	readCh(c)
	/*
		WRITE-ONLY CHANNELS
		The same goes for write-only channels, but the arrow's position moves.
	*/
	writeCh(c)

	saveBackupsMain()

	// TIPS, TRICKS & FUN FACTS:
	/*
		A SEND TO A NIL CHANNEL BLOCKS FOREVER
		A RECEIVE FROM A NIL CHANNEL BLOCKS FOREVER
		A SEND TO A CLOSED CHANNEL PANICS
		A RECEIVE FROM A CLOSED CHANNEL RETURNS THE ZERO VALUE IMMEDIATELY
	*/
}

func readCh(ch <-chan int) {
	// ch can only be read from
	// in this function
}
func writeCh(ch chan<- int) {
	// ch can only be written to
	// in this function
}

// ------------------------------------------

func sendEmail(message string) {
	// keyword go does not accept any return value from functions,
	// it just runs the function and go to the next step -> fmt.Printf(...), it doesnt wait
	go func() {
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("Email received: '%s'\n", message)
	}()
	fmt.Printf("Email sent: '%s'\n", message)
}

// Don't touch below this line

func testMail(message string) {
	sendEmail(message)
	time.Sleep(time.Millisecond * 500)
	fmt.Println("========================")
}

func mailMain() {
	testMail("Hello there Stacy!")
	testMail("Hi there John!")
	testMail("Hey there Jane!")
}

// ------------------------------------------------------------------

func filterOldEmails(emails []email) {
	isOldChan := make(chan bool)

	go sendIsOld(isOldChan, emails)

	isOld := <-isOldChan
	fmt.Println("email 1 is old:", isOld)
	isOld = <-isOldChan
	fmt.Println("email 2 is old:", isOld)
	isOld = <-isOldChan
	fmt.Println("email 3 is old:", isOld)
}

func sendIsOld(isOldChan chan<- bool, emails []email) {
	for _, e := range emails {
		if e.date.Before(time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)) {
			isOldChan <- true
			continue
		}
		isOldChan <- false
	}
}

type email struct {
	body string
	date time.Time
}

func filterTest(emails []email) {
	filterOldEmails(emails)
	fmt.Println("==========================================")
}

func filerMain() {
	filterTest([]email{
		{
			body: "Are you going to make it?",
			date: time.Date(2019, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "I need a break",
			date: time.Date(2021, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "What were you thinking?",
			date: time.Date(2022, 0, 0, 0, 0, 0, 0, time.UTC),
		},
	})
	filterTest([]email{
		{
			body: "Yo are you okay?",
			date: time.Date(2018, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "Have you heard of that website Boot.dev?",
			date: time.Date(2017, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "It's awesome honestly.",
			date: time.Date(2016, 0, 0, 0, 0, 0, 0, time.UTC),
		},
	})
	filterTest([]email{
		{
			body: "Today is the day!",
			date: time.Date(2019, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "What do you want for lunch?",
			date: time.Date(2021, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "Why are you the way that you are?",
			date: time.Date(2022, 0, 0, 0, 0, 0, 0, time.UTC),
		},
	})
	filterTest([]email{
		{
			body: "Did we do it?",
			date: time.Date(2019, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "Letsa Go!",
			date: time.Date(2021, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "Okay...?",
			date: time.Date(2022, 0, 0, 0, 0, 0, 0, time.UTC),
		},
	})
}

// ---------------------------------------------------------------------------

func waitForDbs(numDBs int, dbChan chan struct{}) {
	for i := 0; i < numDBs; i++ {
		<-dbChan
	}
}

func waitTest(numDBs int) {
	dbChan := getDatabasesChannel(numDBs)
	fmt.Printf("Waiting for %v databases...\n", numDBs)
	waitForDbs(numDBs, dbChan)
	time.Sleep(time.Millisecond * 10) // ensure the last print statement happens
	fmt.Println("All databases are online!")
	fmt.Println("=====================================")
}

func waitMain() {
	/*
		Our Mailio server isn't able to boot up until it receives the signal
		that its databases are all online, and it learns about them
		being online by waiting for tokens (empty structs) on a channel.

		Complete the waitForDbs function. It should block until it receives numDBs tokens on the dbChan channel.
		Each time it reads a token the getDatabasesChannel goroutine will print a message to the console for you.
	*/
	waitTest(3)
	waitTest(4)
	waitTest(5)
}

func getDatabasesChannel(numDBs int) chan struct{} {
	ch := make(chan struct{})
	go func() {
		for i := 0; i < numDBs; i++ {
			ch <- struct{}{}
			fmt.Printf("Database %v is online\n", i+1)
		}
	}()
	return ch
}

//------------------------------------------------------

func addEmailsToQueue(emails []string) chan string {
	emailsToSend := make(chan string, len(emails))
	for _, email := range emails {
		emailsToSend <- email
	}
	return emailsToSend
}

func sendEmails(batchSize int, ch chan string) {
	for i := 0; i < batchSize; i++ {
		email := <-ch
		fmt.Println("Sending email:", email)
	}
}

func buffTest(emails ...string) {
	fmt.Printf("Adding %v emails to queue...\n", len(emails))
	ch := addEmailsToQueue(emails)
	fmt.Println("Sending emails...")
	sendEmails(len(emails), ch)
	fmt.Println("==========================================")
}

func buffMain() {

	/*
		We want to be able to send emails in batches. A writing goroutine will write an entire batch
		of email messages to a buffered channel, and later, once the channel is full,
		a reading goroutine will read all of the messages from the channel and send them out to our clients.

		Complete the addEmailsToQueue function.
		It should create a buffered channel with a buffer large enough to store all of the emails it's given.
		It should then write the emails to the channel in order, and finally return the channel.

	*/
	buffTest("Hello John, tell Kathy I said hi", "Whazzup bruther")
	buffTest("I find that hard to believe.", "When? I don't know if I can", "What time are you thinking?")
	buffTest("She says hi!", "Yeah its tomorrow. So we're good.", "Cool see you then!", "Bye!")
}

// --------------------------------------------------------------------

func countReports(numSentCh chan int) int {
	count := 0
	for {
		_, ok := <-numSentCh
		if !ok {
			break
		}
		count++
	}
	return count
}

func closeTest(numBatches int) {
	numSentCh := make(chan int)
	go sendReports(numBatches, numSentCh)

	fmt.Println("Start counting...")
	numReports := countReports(numSentCh)
	fmt.Printf("%v reports sent!\n", numReports)
	fmt.Println("========================")
}

func closeMain() {
	/*
		Use an infinite for loop to read from the channel:
		If the channel is closed, break out of the loop
		Otherwise, keep a running total of the number of reports sent
		Return the total number of reports sent
	*/
	closeTest(3)
	closeTest(4)
	closeTest(5)
	closeTest(6)
}

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
		fmt.Printf("Sent batch of %v reports\n", numReports)
		time.Sleep(time.Millisecond * 100)
	}
	close(ch)
}

// -------------------------------------------------------------------------
func concurrrentFib(n int) {
	nums := make(chan int)
	go fibonacci(n, nums)
	for num := range nums {
		fmt.Println(num)
	}
}

func fibTest(n int) {
	fmt.Printf("Printing %v numbers...\n", n)
	concurrrentFib(n)
	fmt.Println("==============================")
}

func fibMain() {
	/*
		Complete the concurrrentFib function. It should:

		Create a new channel of ints
		Call fibonacci in a goroutine, passing it the channel and the number of Fibonacci numbers to generate, n
		Use a range loop to read from the channel and print out the numbers one by one, each on a new line
	*/
	fibTest(10)
	fibTest(5)
	fibTest(20)
	fibTest(13)
}

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
		time.Sleep(time.Millisecond * 10)
	}
	close(ch)

}

// ------------------------------------------------------------------------------------

func logMessages(chEmails, chSms chan string) {
	for {
		select {
		case s, ok := <-chSms:
			logSms(s)
			if !ok {
				return
			}
		case s, ok := <-chEmails:
			logEmail(s)
			if !ok {
				return
			}
		}
	}
}

func logSms(sms string) {
	fmt.Println("SMS:", sms)
}

func logEmail(email string) {
	fmt.Println("Email:", email)
}

func logTest(sms []string, emails []string) {
	fmt.Println("Starting...")

	chSms, chEmails := sendToLogger(sms, emails)

	logMessages(chEmails, chSms)
	fmt.Println("===============================")
}

func logMain() {
	/*
		Complete the logMessages function.

		Use an infinite for loop and a select statement to log the emails and sms messages
		as they come in order across the two channels.
		Add a condition to return from the function when one of the two channels closes, whichever is first.

		Use the logSms and logEmail functions to log the messages.
	*/

	rand.Seed(0)
	logTest(
		[]string{
			"hi friend",
			"What's going on?",
			"Welcome to the business",
			"I'll pay you to be my friend",
		},
		[]string{
			"Will you make your appointment?",
			"Let's be friends",
			"What are you doing?",
			"I can't believe you've done this.",
		},
	)
	logTest(
		[]string{
			"this song slaps hard",
			"yooo hoooo",
			"i'm a big fan",
		},
		[]string{
			"What do you think of this song?",
			"I hate this band",
			"Can you believe this song?",
		},
	)
}

func sendToLogger(sms, emails []string) (chSms, chEmails chan string) {
	chSms = make(chan string)
	chEmails = make(chan string)
	go func() {
		for i := 0; i < len(sms) && i < len(emails); i++ {
			done := make(chan struct{})
			s := sms[i]
			e := emails[i]
			t1 := time.Millisecond * time.Duration(rand.Intn(1000))
			t2 := time.Millisecond * time.Duration(rand.Intn(1000))
			go func() {
				time.Sleep(t1)
				chSms <- s
				done <- struct{}{}
			}()
			go func() {
				time.Sleep(t2)
				chEmails <- e
				done <- struct{}{}
			}()
			<-done
			<-done
			time.Sleep(10 * time.Millisecond)
		}
		close(chSms)
		close(chEmails)
	}()
	return chSms, chEmails
}

// -----------------------------------------------------------------

func saveBackups(snapshotTicker, saveAfter <-chan time.Time) {
	for {
		select {
		case <-snapshotTicker:
			takeSnapshot()
		case <-saveAfter:
			saveSnapshot()
			return
		default:
			waitForData()
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func takeSnapshot() {
	fmt.Println("Taking a backup snapshot...")
}

func saveSnapshot() {
	fmt.Println("All backups saved!")
}

func waitForData() {
	fmt.Println("Nothing to do, waiting...")
}

func test() {
	snapshotTicker := time.Tick(800 * time.Millisecond)
	saveAfter := time.After(2800 * time.Millisecond)
	saveBackups(snapshotTicker, saveAfter)
	fmt.Println("===========================")
}

func saveBackupsMain() {
	test()
	/*
		Complete the saveBackups function.

		It should read values from the snapshotTicker and saveAfter channels simultaneously.

		- If a value is received from snapshotTicker, call takeSnapshot()
		- If a value is received from saveAfter, call saveSnapshot() and return from the function: you're done.
		- If neither channel has a value ready, call waitForData() and then time.Sleep() for 500 milliseconds.
			After all, we want to show in the logs that the snapshot service is running
	*/
}
