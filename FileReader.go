package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
	"sync"
	"time"
)

func phoneNumbersInFile(filePath string) int {
	file := strings.NewReader(filePath)

	//just googled how to regex and make sure phone number works
	var telephone = regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)

	// make buffered channels then wait
	jobs := make(chan string)
	results := make(chan int)

	//wait group is used to wait for all goroutines to finish
	wg := new(sync.WaitGroup)

	// start up some workers and run
	for worker := 1; worker <= 1000; worker++ {
		//add worker to group
		wg.Add(1)
		//call match phonenumbers to make sure the telephone numbers are valid
		go matchPhoneNumbers(jobs, results, wg, telephone)
	}

	// read the file and point jobs to all the text we read in
	go func() {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			jobs <- scanner.Text()
		}
		//close all jobs
		close(jobs)
	}() //syntax

	// Collect all results BUT MAKE SURE WE CLOSE CHANNEL WHEN PROCESSED
	go func() {
		//suspends execution of all
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

//helper func
func matchPhoneNumbers(jobs <-chan string, results chan<- int, wg *sync.WaitGroup, telephone *regexp.Regexp) {
	// Decrease counter for wg when go routine finishes
	defer wg.Done()
	for j := range jobs {
		if telephone.MatchString(j) {
			results <- 1
		}
	}
}

func sequentialPhoneNumbersInFile(filePath string) int {
	file := strings.NewReader(filePath)

	//regex is the most unreadable thing in existence, wow
	var telephone = regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
	result := 0
	scanner := bufio.NewScanner(file)
	//Go through all the lines in the file
	for scanner.Scan() {
		if telephone.MatchString(scanner.Text()) {
			result++
		}
	}
	return result
}

func main() {
	// read file and process it
	data, err := ioutil.ReadFile("600000numbers.data")

	if err != nil {
		fmt.Println("DID NOT WORK TRY AGAIN")
	}
	var start = time.Now()
	numberOfTelephoneNumbers := phoneNumbersInFile(string(data))
	duration := time.Since(start)
	fmt.Println("Done parallel:", numberOfTelephoneNumbers)
	fmt.Println("Time to read file:", duration)

	var start2 = time.Now()
	numberOfTelephoneNumbers = sequentialPhoneNumbersInFile(string(data))
	duration2 := time.Since(start2)
	fmt.Println("Done sequential:", numberOfTelephoneNumbers)
	fmt.Println("Time to read file:", duration2)

}
