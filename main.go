package main

import "fmt"

// Iterative length calculation O(n)
func IterativeStrLen(inputString string) int {
	inputStrLen := 0
	for i := 0; i < len(inputString); i++ {
		inputStrLen += 1
	}
	return inputStrLen
}

// Recursive length calculation O(n)
func RecursiveStrLen(inputString string) int {
	if inputString == "" {
		return 0
	}
	return 1 + RecursiveStrLen(inputString[1:])
}

func main() {
	inputStr := "LekanProgramming"

	fmt.Println(IterativeStrLen(inputStr))
	fmt.Println(RecursiveStrLen(inputStr))
}
