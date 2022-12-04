package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type Point struct {
	i int
	j int
}

func findBasin(heightmap [][]int64, point Point, done map[Point]bool) (int64, map[Point]bool) {
	total := int64(0)
	if done[point] || point.i < 0 || point.j < 0 || point.i >= len(heightmap) || point.j >= len(heightmap[0]) || heightmap[point.i][point.j] == 9 {
		done[point] = true
		return total, done
	}
	done[point] = true
	var over int64
	var under int64
	var left int64
	var right int64

	if !done[Point{point.i, point.j + 1}] {
		over, done = findBasin(heightmap, Point{point.i, point.j + 1}, done)
	}
	if !done[Point{point.i, point.j - 1}] {
		under, done = findBasin(heightmap, Point{point.i, point.j - 1}, done)
	}
	if !done[Point{point.i - 1, point.j}] {
		left, done = findBasin(heightmap, Point{point.i - 1, point.j}, done)
	}
	if !done[Point{point.i + 1, point.j}] {
		right, done = findBasin(heightmap, Point{point.i + 1, point.j}, done)
	}

	total += 1 + over + under + right + left
	return total, done
}

func main() {
	data, err := os.ReadFile("day9.txt")
	if err != nil {
		log.Fatal(err)
	}

	heightmap_1 := strings.Split(strings.TrimSpace(string(data)), "\n")
	heightmap := make([][]int64, len(heightmap_1))
	max_row := make([]int64, len(heightmap_1[0]))
	for i, heights := range heightmap_1 {
		heightmap[i] = make([]int64, len(heights))
		for j, digit := range heights {
			max_row[j] = int64(10)
			heightmap[i][j] = int64(digit - '0')
		}
	}
	risk := int64(0)
	done := make(map[Point]bool)
	var basins []int64
	for i, heights := range heightmap {
		next_row := max_row
		if i+1 < len(heightmap) {
			next_row = heightmap[i+1]
		}
		last_row := max_row
		if i-1 >= 0 {
			last_row = heightmap[i-1]
		}
		for j, digit := range heights {
			next_col := int64(10)
			if j+1 < len(heights) {
				next_col = heights[j+1]
			}
			last_col := int64(10)
			if j-1 >= 0 {
				last_col = heights[j-1]
			}
			if next_row[j] > digit && last_row[j] > digit && next_col > digit && last_col > digit {
				risk += 1 + heightmap[i][j]
				sizeOfBasin := int64(0)
				sizeOfBasin, done = findBasin(heightmap, Point{i, j}, done)
				basins = append(basins, sizeOfBasin)
			}
		}
	}
	sort.Slice(basins, func(i, j int) bool { return basins[i] < basins[j] })
	part2 := int64(1)
	for _, basin := range basins[len(basins)-3:] {
		part2 *= basin
	}
	fmt.Println(risk, part2)
}
