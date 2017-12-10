package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func partOne(input []byte) int {
	lengths := make([]int, 0, len(input))
	for _, v := range strings.Split(string(input), ",") {
		i, _ := strconv.Atoi(v)
		if i < 256 {
			lengths = append(lengths, i)
		}
	}

	list := make([]int, 256)
	for i := range list {
		list[i] = i
	}

	currentPosition := 0
	skipSize := 0
	knotHashRound(list, lengths, &currentPosition, &skipSize)

	return list[0] * list[1]
}

func partTwo(input []byte) string {
	tmp := make([]byte, len(input))
	copy(tmp, input)
	tmp = append(tmp, []byte{17, 31, 73, 47, 23}...)

	lengths := make([]int, 0, len(tmp))
	for _, v := range tmp {
		lengths = append(lengths, int(v))
	}

	list := make([]int, 256)
	for i := range list {
		list[i] = i
	}

	currentPosition := 0
	skipSize := 0
	for i := 0; i < 64; i++ {
		knotHashRound(list, lengths, &currentPosition, &skipSize)
	}

	denseHash := make([]int, 0, 16)

	for i := 0; i < 16; i++ {
		tmp := list[i*16]
		for j := i*16 + 1; j < i*16+16; j++ {
			tmp ^= list[j]
		}
		denseHash = append(denseHash, tmp)
	}

	result := ""
	for _, v := range denseHash {
		t := strconv.FormatInt(int64(v), 16)
		if len(t) < 2 {
			result += "0"
		}
		result += t
	}

	return result
}

func knotHashRound(input []int, lengths []int, currentPosition, skipSize *int) {
	for _, v := range lengths {

		for s := 0; s < v/2; s++ {
			i := (*currentPosition + s) % len(input)
			j := (*currentPosition + (v - 1) - s) % len(input)
			input[i], input[j] = input[j], input[i]
		}

		*currentPosition = (*currentPosition + v + *skipSize) % len(input)
		*skipSize++
	}
}

func main() {

	bytes, _ := ioutil.ReadFile("input.txt")

	fmt.Println("Part 1 answer :", partOne(bytes))
	fmt.Println("Part 2 answer :", partTwo(bytes))

}
