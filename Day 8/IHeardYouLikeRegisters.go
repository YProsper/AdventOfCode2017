package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func partOneAndTwo(instructions [][]string) (int, int) {
	registers := make(map[string]int)
	allTimeMax := math.MinInt32

	for _, instruction := range instructions {
		reg1 := instruction[0]
		reg2 := instruction[4]

		_, ok := registers[reg1]
		if !ok {
			registers[reg1] = 0
		}

		regValue, ok := registers[reg2]
		if !ok {
			registers[reg2] = 0
		}

		condIsTrue := false
		switch condValue, _ := strconv.Atoi(instruction[6]); instruction[5] {
		case ">":
			condIsTrue = regValue > condValue
		case "<":
			condIsTrue = regValue < condValue
		case ">=":
			condIsTrue = regValue >= condValue
		case "<=":
			condIsTrue = regValue <= condValue
		case "==":
			condIsTrue = regValue == condValue
		case "!=":
			condIsTrue = regValue != condValue
		}

		if condIsTrue {
			switch v, _ := strconv.Atoi(instruction[2]); instruction[1] {
			case "inc":
				registers[reg1] += v
			case "dec":
				registers[reg1] -= v
			}

			if registers[reg1] > allTimeMax {
				allTimeMax = registers[reg1]
			}
		}
	}

	max := math.MinInt32
	for _, v := range registers {
		if v > max {
			max = v
		}
	}
	return max, allTimeMax
}

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)

	instructions := make([][]string, 0, 200)
	for scanner.Scan() {
		instructions = append(instructions, strings.Fields(scanner.Text()))
	}

	// --- Part One and Two---
	max, allTimeMax := partOneAndTwo(instructions)

	fmt.Println("Part 1 answer: ", max)

	fmt.Println("Part 2 answer: ", allTimeMax)
}
