package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

func partOne(filename string) int {
	inputFile, _ := os.Open(filename)
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanLines)

	nbValids := 0
	var re = regexp.MustCompile(`[a-z]*`)
	for scanner.Scan() {
		if !hasDuplicate(re.FindAllString(scanner.Text(), -1)) {
			nbValids++
		}
	}

	return nbValids
}

func hasDuplicate(phrase []string) bool {
	set := make(map[string]bool)
	for _, v := range phrase {
		_, ok := set[v]
		if !ok {
			set[v] = true
		} else {
			return true
		}
	}
	return false
}

func partTwo(filename string) int {
	inputFile, _ := os.Open(filename)
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanLines)

	nbValids := 0
	var re = regexp.MustCompile(`[a-z]*`)
	for scanner.Scan() {
		if !hasPalyndrome(re.FindAllString(scanner.Text(), -1)) {
			nbValids++
		}
	}

	return nbValids
}

func hasPalyndrome(phrase []string) bool {
	set := make(map[string]bool)
	for _, v := range phrase {
		word := sortLetters(v)
		_, ok := set[word]
		if !ok {
			set[word] = true
		} else {
			return true
		}
	}
	return false
}

func sortLetters(s string) string {
	letters := strings.Split(s, "")
	sort.Strings(letters)
	return strings.Join(letters, "")
}

func main() {
	filename := os.Args[1]

	// --- Part One ---
	fmt.Println("Part 1 answer : ", partOne(filename))

	// --- Part Two ---
	fmt.Println("Part 2 answer : ", partTwo(filename))
}
