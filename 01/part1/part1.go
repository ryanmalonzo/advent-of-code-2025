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

func parseInput(file *os.File) []int {
	defer file.Close()

	var values []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		direction := rune(line[0])
		value, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		// If direction is L, store the opposite of value
		if direction == 'L' {
			value = value * -1
		}

		values = append(values, value)
	}

	return values
}

func solve() {
	input := parseInput(readInput())
	currentValue := 50
	const MaxValue = 100
	numberOfZeros := 0

	for _, value := range input {
		currentValue = (currentValue + value) % MaxValue
		if currentValue == 0 {
			numberOfZeros += 1
		}
	}

	fmt.Println(numberOfZeros)
}

func main() {
	solve()
}
