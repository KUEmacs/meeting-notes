package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	data, err := ioutil.ReadFile("doesn't exist")
	// Checking if error exists
	if err != nil {
		fmt.Println("Error occured: ", err)
		os.Exit(1)
	}
}
