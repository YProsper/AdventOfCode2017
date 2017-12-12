package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func partOne(input map[int][]int) int {
	programs := make(map[int]bool)

	var queue []int
	queue = append(queue, 0)

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		_, ok := programs[v]
		if !ok {
			programs[v] = true
			queue = append(queue, input[v]...)
		}
	}

	return len(programs)
}

func partTwo(input map[int][]int) int {
	nbGroups := 0
	for k := range input {
		nbGroups++
		programs := make(map[int]bool)
		var queue []int
		queue = append(queue, k)

		for len(queue) > 0 {
			v := queue[0]
			queue = queue[1:]
			_, ok := programs[v]
			if !ok {
				programs[v] = true
				queue = append(queue, input[v]...)
				delete(input, v)
			}
		}
	}

	return nbGroups
}

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)

	m := make(map[int][]int)
	for scanner.Scan() {
		f := strings.Split(scanner.Text(), " <-> ")
		k, _ := strconv.Atoi(f[0])
		values := strings.Split(f[1], ",")
		for _, v := range values {
			t, _ := strconv.Atoi(strings.TrimSpace(v))
			m[k] = append(m[k], t)
		}

	}

	fmt.Println("Part 1 answer: ", partOne(m))

	fmt.Println("Part 2 answer: ", partTwo(m))

}
