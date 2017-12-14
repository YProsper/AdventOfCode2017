package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func partOne(firewall map[int]int) int {
	penalty := 0
	for k, v := range firewall {
		if k%((v-1)*2) == 0 {
			penalty += k * v
		}
	}
	return penalty
}

func partTwo(firewall map[int]int) int {
	delay := 0

	success := false
	for !success {
		delay++
		failed := false
		for k, v := range firewall {
			if (k+delay)%((v-1)*2) == 0 {
				failed = true
				break
			}
		}
		success = !failed
	}

	return delay
}

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)

	firewall := make(map[int]int)

	for scanner.Scan() {
		values := strings.Split(scanner.Text(), ": ")
		l, _ := strconv.Atoi(values[0])
		d, _ := strconv.Atoi(values[1])

		firewall[l] = d
	}

	fmt.Println("Part 1 answer:", partOne(firewall))
	fmt.Println("Part 2 answer:", partTwo(firewall))
}
