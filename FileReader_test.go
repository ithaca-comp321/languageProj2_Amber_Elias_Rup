package main

import (
	"fmt"
	"io/ioutil"
	"testing"
	"time"
)

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatalf("%s != %s", a, b)
	}
}

func TestReadFile(t *testing.T) {
	start := time.Now()
	data, err := ioutil.ReadFile("test1.data")
	duration := time.Since(start)
	fmt.Println("Time to read file:", duration)
	if err != nil {
		t.Fatal("Could not open file")
	}

	if string(data) != "(945)422-9345\n(945)422-9345\nfejkfe\n(945)422-9345" {
		t.Fatal("Strings contents dont match")
	}

}

func TestCountFromFile(t *testing.T) {
	start := time.Now()
	data, err := ioutil.ReadFile("test1.data")
	data2, err2 := ioutil.ReadFile("test1.data")

	duration := time.Since(start)
	fmt.Println("Time to read file:", duration)

	//Checking phoneNumbers in file
	start = time.Now()
	var nums = phoneNumbersInFile(string(data), string(data2))
	duration = time.Since(start)
	fmt.Println("Time to process file with goroutines:", duration)
	if err != nil || err2 != nil {
		t.Fatal("Could not open file")
	}
	if nums != 6 {
		t.Fatal("Parallel count is wrong, got: ", nums)
	}

	//Sequential test
	start = time.Now()
	nums = sequentialPhoneNumbersInFile(string(data))
	nums2 := sequentialPhoneNumbersInFile(string(data2))
	total := nums + nums2



	duration = time.Since(start)
	fmt.Println("Time to process file sequentially:", duration)
	if total != 6 {
		t.Fatal("Sequential count is wrong, got: ", total)
	}

	// NEW FILE  - 10000numbers.data generated using https://www.randomlists.com/phone-numbers
	// Switched to using 600000numbers.data which is just the above one copied a lot of times, could make it even bigger
	// Feel free to throw some garbage into that file to test other stuff, this is mostly for timing
	start = time.Now()
	data, err = ioutil.ReadFile("600000numbers.data")
	data2, err2 = ioutil.ReadFile("600000numbers.data")
	duration = time.Since(start)
	fmt.Println("Time to read large file:", duration)

	//Checking phoneNumbers in file
	start = time.Now()
	nums = phoneNumbersInFile(string(data), string(data2))
	duration = time.Since(start)
	fmt.Println("Time to process large file with goroutines:", duration)
	if err != nil {
		t.Fatal("Could not open large file")
	}
	if nums != 1200000 {
		t.Fatal("Parallel count is wrong, got: ", nums)
	}

	//Sequential test
	start = time.Now()
	nums = sequentialPhoneNumbersInFile(string(data))
	nums2 = sequentialPhoneNumbersInFile(string(data2))
	total = nums + nums2
	duration = time.Since(start)
	fmt.Println("Time to process large file sequentially:", duration)
	if total != 1200000 {
		t.Fatal("Sequential count is wrong, got:", total)
	}

}
