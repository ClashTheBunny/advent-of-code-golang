package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Point struct {
	i int
	j int
}

func main() {
	// data, err := os.ReadFile("day15.txt")
	data, err := os.ReadFile("day15.txt")
	// 533 too high
	// 529 too high
	if err != nil {
		log.Fatal(err)
	}

	squidmap_1 := strings.Split(strings.TrimSpace(string(data)), "\n")
	squiddict := make(map[Point][]int, 0)
	for j, squids := range squidmap_1 {
		for i, squidvalue := range squids {
			squiddict[Point{i, j}] = []int{int(squidvalue - '0'), 0}
			if i == len(squids)-1 {
				_, found := squiddict[Point{i + 1, j}]
				squiddict[Point{i + 1, j}] = []int{99, 99999}
				if found {
					log.Fatal("Trying to overwrite squid")
				}
			}
			if i == 0 {
				_, found := squiddict[Point{i - 1, j}]
				squiddict[Point{i - 1, j}] = []int{99, 99999}
				if found {
					log.Fatal("Trying to overwrite squid")
				}
			}
			fmt.Printf("%d", squiddict[Point{i, j}])
		}
		if j == len(squidmap_1)-1 {
			j++
			for i, _ := range squids {
				_, found := squiddict[Point{i, j}]
				squiddict[Point{i, j}] = []int{99, 99999}
				if found {
					log.Fatal("Trying to overwrite squid")
				}
				if i == len(squids)-1 {
					_, found = squiddict[Point{i + 1, j}]
					squiddict[Point{i + 1, j}] = []int{99, 99999}
					if found {
						log.Fatal("Trying to overwrite squid")
					}
				}
				if i == 0 {
					_, found := squiddict[Point{i - 1, j}]
					squiddict[Point{i - 1, j}] = []int{99, 99999}
					if found {
						log.Fatal("Trying to overwrite squid")
					}
				}
			}
		}
		if j == 0 {
			j--
			for i, _ := range squids {
				_, found := squiddict[Point{i, j}]
				squiddict[Point{i, j}] = []int{99, 99999}
				if found {
					log.Fatal("Trying to overwrite squid")
				}
				if i == len(squids)-1 {
					_, found = squiddict[Point{i + 1, j}]
					squiddict[Point{i + 1, j}] = []int{99, 99999}
					if found {
						log.Fatal("Trying to overwrite squid")
					}
				}
				if i == 0 {
					_, found := squiddict[Point{i - 1, j}]
					squiddict[Point{i - 1, j}] = []int{99, 99999}
					if found {
						log.Fatal("Trying to overwrite squid")
					}
				}
			}
		}
	}

	start := Point{0, 0}
	end := Point{len(squidmap_1) - 1, len(squidmap_1[0]) - 1}

	for j := len(squidmap_1) - 1; j >= 0; j-- {
		for i := len(squidmap_1[0]) - 1; i >= 0; i-- {
			point := Point{i, j}
			// if point == end {
			// 	fmt.Printf("%d", squiddict[point])
			// 	continue
			// }
			if point == start {
				fmt.Printf("%d", squiddict[point])
				break
			}
			cost := min(squiddict[Point{i + 1, j}], squiddict[Point{i, j + 1}])
			squiddict[point][1] = cost[0] + squiddict[point][0]
			fmt.Printf("%d", squiddict[point])
		}
		fmt.Printf("\n")
	}
	totalCost := 0
outerLoop:
	for i := 0; i < len(squidmap_1); {
		for j := 0; j < len(squidmap_1[0]); {
			point := Point{i, j}
			if point != start {
				totalCost += squiddict[point][0]
			}
			if point == end {
				break outerLoop
			}
			pointx := Point{i + 1, j}
			pointy := Point{i, j + 1}
			fmt.Printf("%d: %d", point, squiddict[point])
			squidx, foundx := squiddict[pointx]
			squidy, foundy := squiddict[pointy]
			if !foundx && !foundy {
				break outerLoop
			}
			if !foundx {
				j++
			}
			if !foundy {
				i++
			}
			if squidx[1] > squidy[1] {
				j++
			} else {
				i++
			}
		}
	}
	fmt.Println("exampleAnswer:", 40, totalCost)
}

func min(x, y []int) []int {
	if x[1] < y[1] {
		return x
	} else {
		return y
	}
}
