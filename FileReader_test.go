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
	duration := time.Since(start)
	fmt.Println("Time to read file:", duration)
	//Checking phoneNumbers in file
	var nums = phoneNumbersInFile(string(data))
	if err != nil {
		t.Fatal("Could not open file")
	}
	if nums != 3 {
		t.Fatal("Count is wrong")
	}

}
