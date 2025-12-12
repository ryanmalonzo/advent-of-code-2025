package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput() *os.File {
	const FilePath = "../input.txt"

	file, err := os.Open(FilePath)
	if err != nil {
		panic(err)
	}
	return file
}

func parseInput(file *os.File) []string {
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, strings.TrimSpace(scanner.Text()))
	}

	return lines
}

func solve() {
	input := parseInput(readInput())
	currentValue := 50
	const MaxValue = 100
	numberOfZeros := 0

	for _, line := range input {
		direction := line[0]
		distance, _ := strconv.Atoi(line[1:])

		for range distance {
			if direction == 'R' {
				currentValue++
			} else {
				currentValue--
			}

			if currentValue < 0 {
				currentValue = MaxValue - 1
			} else if currentValue == MaxValue {
				currentValue = 0
			}

			if currentValue == 0 {
				numberOfZeros++
			}
		}
	}

	fmt.Println(numberOfZeros)
}

func main() {
	solve()
}
