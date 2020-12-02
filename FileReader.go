package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"
	"sync"
)

func phoneNumbersInFile(path string) int {
	file := strings.NewReader(path)

	//just googled how to regex and make sure phone number works
	var telephone = regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)

	// make buffered channels then wait
	jobs := make(chan string)
	results := make(chan int)
	wg := new(sync.WaitGroup)

	// start up some workers and run
	for w := 1; w <= 5; w++ {
		wg.Add(1)
		go matchPhoneNumbers(jobs, results, wg, telephone)
	}

	// Go over a file line by line and queue up a ton of work
	go func() {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			jobs <- scanner.Text()
		}
		close(jobs)
	}()

	// Collect all results BUT MAKE SURE WE CLOSE CHANNEL WHEN PROCESSED
	go func() {
		wg.Wait()
		close(results)
	}()

	// Add up the results from the results channel.
	counts := 0
	for v := range results {
		counts += v
	}

	return counts
}

func matchPhoneNumbers(jobs <-chan string, results chan<- int, wg *sync.WaitGroup, telephone *regexp.Regexp) {
	// Decrease counter for wait-group when go routine finishes
	defer wg.Done()
	for j := range jobs {
		if telephone.MatchString(j) {
			results <- 1
		}
	}
}

func main() {
	// Just passing numbers in for now but need it to process from file
	const input = "(555) 123-3456\n(555) 123-3456\n(555) 123-3456"
	numberOfTelephoneNumbers := phoneNumbersInFile(input)
	fmt.Println(numberOfTelephoneNumbers)
}
