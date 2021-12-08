package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Geiser struct {
	start Point `json:"start"`
	end   Point `json:"end"`
}

type Point struct {
	x int `json:"x"`
	y int `json:"y"`
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	data, err := os.ReadFile("day5.txt")
	if err != nil {
		log.Fatal(err)
	}
	data_s := strings.TrimSpace(string(data))
	geisers := strings.Split(data_s, "\n")

	geiser_map := make(map[Point]int)
	var geiser_slice []Geiser

	for geiser := range geisers {
		var thisGeiser Geiser
		_, err := fmt.Sscanf(geisers[geiser], "%d,%d -> %d,%d", &thisGeiser.start.x, &thisGeiser.start.y, &thisGeiser.end.x, &thisGeiser.end.y)
		if err != nil {
			log.Fatal(err)
		}
		geiser_slice = append(geiser_slice, thisGeiser)
		y := thisGeiser.start.y
		x := thisGeiser.start.x
		x_direction := -1
		if thisGeiser.end.x-thisGeiser.start.x > 0 {
			x_direction = 1
		} else if thisGeiser.end.x-thisGeiser.start.x == 0 {
			x_direction = 0
		}
		y_direction := -1
		if thisGeiser.end.y-thisGeiser.start.y > 0 {
			y_direction = 1
		} else if thisGeiser.end.y-thisGeiser.start.y == 0 {
			y_direction = 0
		}
		for x != thisGeiser.end.x || y != thisGeiser.end.y {
			geiser_map[Point{x, y}]++
			y += y_direction
			x += x_direction
		}
		geiser_map[Point{x, y}]++
	}
	n := 0
	for _, geiser_count := range geiser_map {
		if geiser_count > 1 {
			n++
		}
	}
	fmt.Println(n)
	// too low: 18445
	// too low: 18571
}
