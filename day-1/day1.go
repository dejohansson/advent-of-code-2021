package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	counter := 0

	var depths []int64
	for scanner.Scan() {
		x, _ := strconv.ParseInt(scanner.Text(), 0, 64)
		depths = append(depths, x)
	}

	for i, x := range depths {
		if i > 2 {
			prev := depths[i-3] + depths[i-2] + depths[i-1]
			current := depths[i-2] + depths[i-1] + x
			if current > prev {
				counter += 1
			}
		}
	}

	fmt.Printf("%d\n", counter)
}
