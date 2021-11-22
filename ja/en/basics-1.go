package main

import (
	"fmt"
)

func main() {

	// score
	totalScore := 0

	// 
	questions := []string{
		"",
		"",
		"",
	}

	// 
	answers := []string{
		"",
		"",
		"",
	}

	// 
	if len(questions) == len(answers) {
		for i := 0; i < len(questions); i++ {
			ask(questions[i], answers[i], &totalScore)
		}
	} else {
		fmt.Printf("The length of given arrays is not the same.")
	}

	fmt.Printf("Result: %d\n", totalScore)

}

func ask(q string, a string, totalScore *int) {

	// user input
	var input string

	// 
	fmt.Printf("Translate this sentense:\n")
	fmt.Printf("%s\n", q)
	fmt.Scan(&input)

	if input == a {
		fmt.Printf("Great!")
		*totalScore += 10
	} else {
		fmt.Printf("Miss")
	}

	fmt.Printf("\n\n")

}
