package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func base2Tobase10(base2 []uint8) int {
	base10 := 0
	for i := range base2 {
		base10 += int(base2[len(base2)-1-i]) * int(math.Pow(2, float64(i)))
	}
	return base10
}

func getOGR(input [][]uint8, p int) int {
	// fmt.Printf("Len: %d, P: %d\n", len(input), p)
	if len(input) == 1 {
		return base2Tobase10(input[0])
	}

	zeroRows := [][]uint8{}
	oneRows := [][]uint8{}

	for _, r := range input {
		if r[p] == 0 {
			zeroRows = append(zeroRows, r)
		} else {
			oneRows = append(oneRows, r)
		}
	}

	if len(zeroRows) > len(oneRows) {
		return getOGR(zeroRows, p+1)
	} else {
		return getOGR(oneRows, p+1)
	}
}

func getCSR(input [][]uint8, p int) int {
	// fmt.Printf("Len: %d, P: %d\n", len(input), p)
	if len(input) == 1 {
		return base2Tobase10(input[0])
	}

	zeroRows := [][]uint8{}
	oneRows := [][]uint8{}

	for _, r := range input {
		if r[p] == 0 {
			zeroRows = append(zeroRows, r)
		} else {
			oneRows = append(oneRows, r)
		}
	}

	if len(zeroRows) <= len(oneRows) {
		return getCSR(zeroRows, p+1)
	} else {
		return getCSR(oneRows, p+1)
	}
}

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rowCounter := 1
	bitCounter := []int{}
	input := [][]uint8{{}}

	scanner.Scan()
	for _, strBit := range scanner.Text() {
		b, _ := strconv.Atoi(string(strBit))
		bitCounter = append(bitCounter, b)
		input[0] = append(input[0], uint8(b))
	}

	for scanner.Scan() {
		row := []uint8{}
		for i, strBit := range scanner.Text() {
			b, _ := strconv.Atoi(string(strBit))
			bitCounter[i] += b
			row = append(row, uint8(b))
		}
		rowCounter++
		input = append(input, row)
	}

	gammaRate := 0
	epsilonRate := 0
	for i := range bitCounter {
		if bitCounter[len(bitCounter)-1-i] > rowCounter/2 {
			gammaRate += int(math.Pow(2, float64(i)))
		} else {
			epsilonRate += int(math.Pow(2, float64(i)))
		}
	}

	fmt.Printf("Gamma: %d, Epsilon: %d\n", gammaRate, epsilonRate)
	fmt.Printf("Part 1: %d\n", gammaRate*epsilonRate)

	ogr := getOGR(input, 0)
	csr := getCSR(input, 0)
	fmt.Printf("OGR: %d, CSR: %d\n", ogr, csr)
	fmt.Printf("Part 2: %d\n", ogr*csr)
}
