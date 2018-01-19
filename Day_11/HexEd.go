package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func partOne(input string) int {
	x, y := 0, 0
	for _, v := range strings.Split(input, ",") {
		switch v {
		case "n":
			y--
		case "s":
			y++
		case "ne":
			x++
			y--
		case "nw":
			x--
		case "se":
			x++
		case "sw":
			x--
			y++
		}
	}

	return (abs(x) + abs(x+y) + abs(y)) / 2
}

func partTwo(input string) int {
	x, y, furthest := 0, 0, 0
	for _, v := range strings.Split(input, ",") {
		switch v {
		case "n":
			y--
		case "s":
			y++
		case "ne":
			x++
			y--
		case "nw":
			x--
		case "se":
			x++
		case "sw":
			x--
			y++
		}

		d := (abs(x) + abs(x+y) + abs(y)) / 2
		if d > furthest {
			furthest = d
		}
	}

	return furthest
}

func main() {
	bytes, _ := ioutil.ReadFile("input.txt")

	fmt.Println("Part 1 answer: ", partOne(string(bytes)))
	fmt.Println("Part 2 answer: ", partTwo(string(bytes)))
}
