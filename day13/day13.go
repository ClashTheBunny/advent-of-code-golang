package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func fold_along(dots map[[2]int64]bool, direction []int64) map[[2]int64]bool {
	newdots := make(map[[2]int64]bool)
	for dot := range dots {
		if dot[direction[0]] > direction[1] {
			dot[direction[0]] = direction[1] - (dot[direction[0]] - direction[1])
		}
		newdots[dot] = true
	}
	return newdots
}

func main() {

	data, err := os.ReadFile("day13.txt")
	if err != nil {
		log.Fatal(err)
	}

	folds_and_dots := strings.Split(strings.TrimSpace(string(data)), "\n\n")

	dots := make(map[[2]int64]bool, 0)
	for _, dot := range strings.Split(folds_and_dots[0], "\n") {
		tmp := strings.Split(dot, ",")
		x_cord, err := strconv.ParseInt(tmp[0], 0, 64)
		if err != nil {
			log.Fatal(err)
		}
		y_cord, err := strconv.ParseInt(tmp[1], 0, 64)
		if err != nil {
			log.Fatal(err)
		}
		dots[[2]int64{x_cord, y_cord}] = true
	}

	folds := make([][]int64, 0)
	for _, fold := range strings.Split(folds_and_dots[1], "\n") {
		tmp := strings.Split(fold, "=")
		direction := int64(strings.Split(tmp[0], " ")[2][0]) - 'x'
		fold_location, err := strconv.ParseInt(tmp[1], 0, 64)
		if err != nil {
			log.Fatal(err)
		}
		folds = append(folds, []int64{direction, fold_location})
	}
	for i, fold := range folds {
		dots = fold_along(dots, fold)
		if i == 0 {
			fmt.Println(len(dots))
		}
	}
	x := int64(0)
	y := int64(0)
	for _, fold := range folds {
		if fold[0] == 0 {
			x = fold[1]
		}
		if fold[0] == 1 {
			y = fold[1]
		}
	}
	for j := int64(0); j <= y; j++ {
		for i := int64(0); i <= x; i++ {
			if dots[[2]int64{i, j}] {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
}
