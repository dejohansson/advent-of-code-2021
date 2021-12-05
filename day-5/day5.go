package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const FloorSize = 1000

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	floor1 := [FloorSize][FloorSize]int{}
	floor2 := [FloorSize][FloorSize]int{}

	for scanner.Scan() {
		row := strings.Split(strings.Replace(scanner.Text(), " -> ", ",", 1), ",")
		x1, _ := strconv.Atoi(row[0])
		y1, _ := strconv.Atoi(row[1])
		x2, _ := strconv.Atoi(row[2])
		y2, _ := strconv.Atoi(row[3])

		// fmt.Printf("%d,%d -> %d,%d\n", x1, y1, x2, y2)

		if x1 == x2 {
			if y1 > y2 {
				for y := y2; y <= y1; y++ {
					floor1[x1][y] += 1
					floor2[x1][y] += 1
				}
			} else {
				for y := y1; y <= y2; y++ {
					floor1[x1][y] += 1
					floor2[x1][y] += 1
				}
			}
		} else if y1 == y2 {
			if x1 > x2 {
				for x := x2; x <= x1; x++ {
					floor1[x][y1] += 1
					floor2[x][y1] += 1
				}
			} else {
				for x := x1; x <= x2; x++ {
					floor1[x][y1] += 1
					floor2[x][y1] += 1
				}
			}
		} else {
			x, y := x1, y1
			for {
				floor2[x][y] += 1

				if x == x2 {
					break
				}

				if x2 > x1 {
					x += 1
				} else {
					x -= 1
				}
				if y2 > y1 {
					y += 1
				} else {
					y -= 1
				}
			}
		}
	}

	intersections1 := 0
	for _, r := range floor1 {
		for _, p := range r {
			// fmt.Printf("%d ", p)
			if p > 1 {
				intersections1 += 1
			}
		}
		// fmt.Println("")
	}

	fmt.Printf("Part 1: %d\n\n", intersections1)

	intersections2 := 0
	for _, r := range floor2 {
		for _, p := range r {
			// fmt.Printf("%d ", p)
			if p > 1 {
				intersections2 += 1
			}
		}
		// fmt.Println("")
	}

	fmt.Printf("Part 2: %d\n", intersections2)
}
