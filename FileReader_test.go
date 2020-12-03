package main

import (
	"io/ioutil"
	"testing"
)

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatalf("%s != %s", a, b)
	}
}

func TestReadFile(t *testing.T) {
	data, err := ioutil.ReadFile("test1.data")
	if err != nil {
		t.Fatal("Could not open file")
	}

	if string(data) != "(945)422-9345\n(945)422-9345\nfejkfe\n(945)422-9345" {
		t.Fatal("Strings contents dont match")
	}

}
func TestCountFromFile(t *testing.T){
	data, err := ioutil.ReadFile("test1.data")
	//Checking phoneNumbers in file
	var nums = phoneNumbersInFile(string(data))
	if err != nil {
		t.Fatal("Could not open file")
	}
	if nums != 3 {
		t.Fatal("Count is wrong")
	}
	
}
