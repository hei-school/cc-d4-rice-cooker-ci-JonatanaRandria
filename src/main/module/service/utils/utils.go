package main

import (
	"fmt"
	"strconv"
)

// showMenu displays a menu, validates user input, and calls the specified operation.
func showMenu(message string, valid []int, release int, operation func(int)) {
	fmt.Print(message)
	var input string
	fmt.Scanln(&input)

	if choice, err := strconv.Atoi(input); err == nil && contains(valid, choice) {
		operation(choice)

		if choice != release {
			showMenu(message, valid, release, operation)
		}
	} else {
		fmt.Printf("Invalid, your input must be one of %v\n", valid)
		showMenu(message, valid, release, operation)
	}
}

// contains checks if a slice contains a given element.
func contains(slice []int, elem int) bool {
	for _, v := range slice {
		if v == elem {
			return true
		}
	}
	return false
}
