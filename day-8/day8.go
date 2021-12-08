package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func contains(a string, b string) bool {
	for _, x := range b {
		match := false
		for _, y := range a {
			if x == y {
				match = true
			}
		}
		if !match {
			return false
		}
	}
	return true
}

func equal(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}
	for _, x := range b {
		match := false
		for _, y := range a {
			if x == y {
				match = true
			}
		}
		if !match {
			return false
		}
	}
	return true
}

func part2(inputFile string) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "|")
		input := strings.Fields(strings.TrimSpace(line[0]))
		output := strings.Fields(strings.TrimSpace(line[1]))

		parsedInput := map[int]string{}

		// 1, 4, 7, 8
		for i := len(input) - 1; i >= 0; i-- {
			s := input[i]
			if len(s) == 2 {
				parsedInput[1] = s
			} else if len(s) == 4 {
				parsedInput[4] = s
			} else if len(s) == 3 {
				parsedInput[7] = s
			} else if len(s) == 7 {
				parsedInput[8] = s
			} else {
				continue
			}
			input = append(input[:i], input[i+1:]...)
		}

		// 3, 9
		for i := len(input) - 1; i >= 0; i-- {
			s := input[i]
			if len(s) == 5 {
				if contains(s, parsedInput[1]) {
					parsedInput[3] = s
					input = append(input[:i], input[i+1:]...)
				}
			} else if len(s) == 6 {
				if contains(s, parsedInput[4]) {
					parsedInput[9] = s
					input = append(input[:i], input[i+1:]...)
				}
			}
		}

		// 0, 5
		for i := len(input) - 1; i >= 0; i-- {
			s := input[i]
			if len(s) == 5 {
				if contains(parsedInput[9], s) {
					parsedInput[5] = s
					input = append(input[:i], input[i+1:]...)
				}
			} else if len(s) == 6 {
				if contains(s, parsedInput[1]) {
					parsedInput[0] = s
					input = append(input[:i], input[i+1:]...)
				}
			}
		}

		// 2, 6
		for i := len(input) - 1; i >= 0; i-- {
			s := input[i]
			if len(s) == 5 {
				parsedInput[2] = s
			} else if len(s) == 6 {
				parsedInput[6] = s
			}
		}

		outVal := 0
		for i, seq := range output {
			for num, pSeq := range parsedInput {
				if equal(seq, pSeq) {
					outVal += int(math.Pow(10, float64(3-i))) * num
					break
				}
			}
		}
		sum += outVal
	}

	fmt.Printf("Part 2: %d\n", sum)
}

func part1(inputFile string) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	uniqueOutputs := 0

	for scanner.Scan() {
		output := strings.Fields(strings.TrimSpace(strings.Split(scanner.Text(), "|")[1]))
		for _, o := range output {
			if len(o) == 2 || len(o) == 4 || len(o) == 3 || len(o) == 7 {
				uniqueOutputs++
			}
		}
	}

	fmt.Printf("Part 1: %d\n", uniqueOutputs)
}

func main() {
	fileName := "data.txt"
	part1(fileName)
	part2(fileName)
}
