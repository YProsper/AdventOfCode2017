package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func partOne(filename string) int {
	input, _ := os.Open(filename)
	defer input.Close()

	scanner := bufio.NewScanner(input)

	table := make([]int, 0, 50)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		table = append(table, i)
	}

	ip, nbJmps := 0, 0
	for ip < len(table) {
		jmp := table[ip]
		table[ip]++
		ip += jmp
		nbJmps++
	}

	return nbJmps
}

func partTwo(filename string) int {
	input, _ := os.Open(filename)
	defer input.Close()

	scanner := bufio.NewScanner(input)

	table := make([]int, 0, 50)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		table = append(table, i)
	}

	ip, nbJmps := 0, 0
	for ip < len(table) {
		jmp := table[ip]
		if jmp >= 3 {
			table[ip]--
		} else {
			table[ip]++
		}
		ip += jmp
		nbJmps++
	}

	return nbJmps
}

func main() {
	filename := os.Args[1]
	fmt.Println("Part 1 answer: ", partOne(filename))
	fmt.Println("Part 2 answer: ", partTwo(filename))
}
