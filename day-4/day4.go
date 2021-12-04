package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type BingoNum struct {
	num    int
	marked bool
}

func isWinner(board [5][5]BingoNum, x int, y int) bool {
	win := true
	for i := 0; i < 5; i++ {
		if !board[i][y].marked {
			win = false
			break
		}
	}
	if win {
		return true
	}
	win = true
	for i := 0; i < 5; i++ {
		if !board[x][i].marked {
			win = false
			break
		}
	}
	return win
}

func getScore(board [5][5]BingoNum, lastNum int) int {
	var sumUnmarked int
	for _, row := range board {
		for _, bNum := range row {
			if !bNum.marked {
				sumUnmarked += bNum.num
			}
		}
	}
	fmt.Printf("Sum unmarked: %d, Last Number: %d\n", sumUnmarked, lastNum)
	return sumUnmarked * lastNum
}

func play1(order []int, boards [][5][5]BingoNum) int {
	for _, num := range order {
		for i, board := range boards {
		search:
			for j, row := range board {
				for k, bNum := range row {
					if bNum.num == num {
						boards[i][j][k].marked = true
						if isWinner(boards[i], j, k) {
							return getScore(boards[i], num)
						}
						break search
					}
				}
			}
		}
	}
	return 0
}

func play2(order []int, boards [][5][5]BingoNum) int {
	winners := make(map[int]bool)
	for _, num := range order {
		for i, board := range boards {
			if !winners[i] {
			search:
				for j, row := range board {
					for k, bNum := range row {
						if bNum.num == num {
							boards[i][j][k].marked = true
							if isWinner(boards[i], j, k) {
								if len(boards)-1 == len(winners) {
									return getScore(boards[i], num)
								}
								winners[i] = true
							}
							break search
						}
					}
				}
			}
		}
	}
	return 0
}

func main() {
	byteFile, _ := ioutil.ReadFile("data.txt")
	splitFile := strings.Split(string(byteFile), "\r\n\r\n")

	order := []int{}
	order = append(order)
	for _, s := range strings.Split(splitFile[0], ",") {
		n, _ := strconv.Atoi(s)
		order = append(order, n)
	}

	boards := [][5][5]BingoNum{}
	for _, strBoard := range splitFile[1:] {
		newBoard := [5][5]BingoNum{}

		for i, strRow := range strings.Split(strBoard, "\r\n") {
			for j, strNum := range strings.Fields(strings.TrimSpace(strRow)) {
				num, _ := strconv.Atoi(strNum)
				newBoard[i][j] = BingoNum{
					num:    num,
					marked: false,
				}
			}
		}

		boards = append(boards, newBoard)
	}

	fmt.Printf("Part 1: %d\n", play1(order, boards))
	fmt.Printf("Part 2: %d\n", play2(order, boards))
}
