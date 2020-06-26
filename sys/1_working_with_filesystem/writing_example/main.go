package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Please specify a path and some content")
		return
	}
	if err := ioutil.WriteFile(os.Args[1], []byte(os.Args[2]), 0644); err != nil {
		fmt.Println("Error: ", err)
	}
}

// This example writes all the content at once in a single operation.
// This requires that we allocate all the content in memory using a byte slice.
// If the content is too large, memory usage can become a problem for the OS, which could kill the process of our application.
