package main

import (
	"fmt"
	"io/ioutil"
)

func partOne(input string) int {
	currentScore, score := 0, 0
	garbage := false
	for i := 0; i < len(input); i++ {
		switch c := input[i]; c {
		case '{':
			if !garbage {
				currentScore++
			}
		case '}':
			if !garbage {
				score += currentScore
				currentScore--
			}
		case '<':
			if !garbage {
				garbage = true
			}
		case '>':
			garbage = false
		case '!':
			i++
		}
	}

	return score
}

func partTwo(input string) int {
	nbChars := 0
	garbage := false
	for i := 0; i < len(input); i++ {
		switch c := input[i]; c {
		case '<':
			if garbage {
				nbChars++
			} else {
				garbage = true
			}
		case '>':
			garbage = false
		case '!':
			i++
		default:
			if garbage {
				nbChars++
			}
		}
	}

	return nbChars
}

func main() {
	bytes, _ := ioutil.ReadFile("input.txt")

	input := string(bytes)

	fmt.Println("Part 1 answer:", partOne(input))
	fmt.Println("Part 2 answer:", partTwo(input))
}
