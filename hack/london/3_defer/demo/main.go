package main

import (
	"fmt"
	"log"
)

func one() {
	fmt.Println("one")
}

func two() {
	fmt.Println("two")
}

func three() {
	fmt.Println("three")
}

func main() {
	defer one()
	defer two()
	defer three()
	for _, item := range []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
		"ten",
	} {
		defer func(item string) {
			log.Println(item)
		}(item)
	}
}
