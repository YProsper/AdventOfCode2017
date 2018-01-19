package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func partOneAndTwo(filename string) (int, int) {
	area := readInputFile(filename)
	states := make(map[string]int)
	nbCycles := 0

	for {
		nbCycles++
		max := 0
		for i, v := range area {
			if v > area[max] || (v == area[max] && i < max) {
				max = i
			}
		}

		a := area[max]
		area[max] = 0

		i := (max + 1) % len(area)
		for a > 0 {
			area[i]++
			a--
			i = (i + 1) % len(area)
		}

		s := ""
		for _, v := range area {
			s += strconv.Itoa(v)
		}

		cycles, ok := states[s]
		if ok {
			return len(states) + 1, nbCycles - cycles
		}
		states[s] = nbCycles
	}
}

func readInputFile(filename string) []int {
	input, _ := os.Open(filename)
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanWords)

	result := make([]int, 0, 16)
	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		result = append(result, value)
	}

	return result
}

func main() {
	filename := os.Args[1]
	a1, a2 := partOneAndTwo(filename)
	fmt.Printf("Part 1 anwser: %d\n", a1)
	fmt.Printf("Part 2 anwser: %d", a2)
}
