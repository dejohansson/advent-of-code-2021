package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var m map[int]int

func spawn(days int) int {
	d, ok := m[days]
	if ok {
		return d
	}

	var res int
	if days > 9 {
		res = 1 + spawn(days-9) + spawn(days-7)
	} else if days > 7 {
		res = 1 + spawn(days-7)
	} else {
		res = 1
	}

	m[days] = res
	return res
}

func main() {
	m = make(map[int]int)
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	input := strings.Split(scanner.Text(), ",")

	fish := []int{}
	for _, s := range input {
		num, _ := strconv.Atoi(s)
		fish = append(fish, num)
	}

	part1 := 0
	part2 := 0
	for _, f := range fish {
		part1 += spawn(80-f) + 1
		part2 += spawn(256-f) + 1
	}

	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
