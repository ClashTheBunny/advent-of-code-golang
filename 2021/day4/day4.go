package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func boardColumn(board [][]int64, columnIndex int) (column []int64) {
	column = make([]int64, 0)
	for _, row := range board {
		column = append(column, row[columnIndex])
	}
	return
}

func max(myslice []int64) (m int64) {
	for i, e := range myslice {
		if e > m || i == 0 {
			m = e
		}
	}
	return
}

func main() {
	// data, err := os.ReadFile("day4.txt") # clashthebunny, not the right account
	// data, err := os.ReadFile("day4_sample.txt")
	data, err := os.ReadFile("day4_randall.txt")
	if err != nil {
		log.Fatal(err)
	}
	data_s := strings.TrimSpace(string(data))

	bingo := strings.Split(data_s, "\n\n")

	numbersCalled := strings.Split(bingo[0], ",")
	numbersCalledInt := make([]int64, len(numbersCalled))
	for number := range numbersCalled {
		numbersCalledInt[number], _ = strconv.ParseInt(numbersCalled[number], 0, 64)
	}

	bingoBoards := make([][][]int64, len(bingo[1:]))

	for board := range bingo[1:] {
		lines := strings.Split(bingo[board+1], "\n")
		bingoBoards[board] = make([][]int64, len(lines))
		for line := range lines {
			splitLine := strings.Fields(lines[line])
			splitLineInt := make([]int64, len(splitLine))
			for col := range splitLine {
				splitLineInt[col], _ = strconv.ParseInt(splitLine[col], 0, 64)
			}
			bingoBoards[board][line] = splitLineInt
		}
	}
	fmt.Println(numbersCalledInt)
	fmt.Println(bingoBoards)
	board := 0
	currentNumber := int64(0)

	finished := make(map[int]bool, 0)
VerifyLoop:
	for number := range numbersCalledInt {
		currentNumber = numbersCalledInt[number]

		for board := range bingoBoards {
			for row := range bingoBoards[board] {
				for col := range bingoBoards[board][row] {
					if bingoBoards[board][row][col] == currentNumber {
						bingoBoards[board][row][col] = -currentNumber
						if currentNumber == 0 {
							bingoBoards[board][row][col] = -1000
						}
					}
				}
			}
		}
		for board = range bingoBoards {
			for row := range bingoBoards[board] {
				if max(bingoBoards[board][row]) < 0 {
					// fmt.Println("bingo! board:", board, "row", row)
					finished[board] = true
					if len(finished) == 1 {
						fmt.Println("bingo! board ", board, ":", bingoBoards[board])
						sum := int64(0)
						for row := range bingoBoards[board] {
							for col := range bingoBoards[board][row] {
								if bingoBoards[board][row][col] > 0 {
									sum += bingoBoards[board][row][col]
								}

							}
						}
						fmt.Println(sum, currentNumber, sum*currentNumber)
					}
					if len(finished) == len(bingoBoards) {
						break VerifyLoop
					}
				}
			}
			for col := range bingoBoards[board][0] {
				currcol := boardColumn(bingoBoards[board], col)

				sort.Slice(currcol, func(i, j int) bool { return currcol[i] < currcol[j] })
				if currcol[len(bingoBoards[board][0])-1] < 0 {
					// fmt.Println("bingo! board:", board, "column", col)
					finished[board] = true
					if len(finished) == 1 {
						fmt.Println("bingo! board ", board, ":", bingoBoards[board])
						sum := int64(0)
						for row := range bingoBoards[board] {
							for col := range bingoBoards[board][row] {
								if bingoBoards[board][row][col] > 0 {
									sum += bingoBoards[board][row][col]
								}

							}
						}
						fmt.Println(sum, currentNumber, sum*currentNumber)
					}
					if len(finished) == len(bingoBoards) {
						break VerifyLoop
					}
				}
			}
		}
	}
	fmt.Println("bingo! board ", board, ":", bingoBoards[board])
	sum := int64(0)
	for row := range bingoBoards[board] {
		for col := range bingoBoards[board][row] {
			if bingoBoards[board][row][col] > 0 {
				sum += bingoBoards[board][row][col]
			}

		}
	}
	fmt.Println(sum, currentNumber, sum*currentNumber)

	// Part 1
	//5440 - too low
	//63424 - too low (right for someone else)

	// Part 2
	// 1080 too low
}
