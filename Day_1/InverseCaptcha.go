package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func inverseCaptcha(input string, match int) int {
	result := 0

	for i, v := range input {
		if input[i] == input[(i+match)%len(input)] {
			result += int(v - '0')
		}
	}

	return result
}

func main() {
	inputBytes, _ := ioutil.ReadFile(os.Args[1])
	input := string(inputBytes)

	// --- Part One ---
	fmt.Println("Part 1 answer : ", inverseCaptcha(input, 1))

	// --- Part Two ---
	fmt.Println("Part 2 answer : ", inverseCaptcha(input, len(input)/2))
}
