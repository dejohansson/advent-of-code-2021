package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	depth1 := int64(0)
	position1 := int64(0)
	depth2 := int64(0)
	position2 := int64(0)
	aim := int64(0)

	for scanner.Scan() {
		row := strings.Fields(scanner.Text())
		command := row[0]
		num, _ := strconv.ParseInt(row[1], 10, 0)

		if command == "forward" {
			position1 += num
			position2 += num
			depth2 += aim * num
		} else if command == "down" {
			depth1 += num
			aim += num
		} else if command == "up" {
			depth1 -= num
			aim -= num
		}
	}

	fmt.Printf("Part 1: %d\n", depth1*position1)
	fmt.Printf("Part 2: %d\n", depth2*position2)
}
