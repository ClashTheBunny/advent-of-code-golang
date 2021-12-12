package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/kr/pretty"
)

type Point struct {
	i int
	j int
}

func increaseSquid(squiddict map[Point][]int, point Point) int {
	energy := 0
	var tmp []int
	var found bool
	for di := -1; di <= 1; di++ {
		for dj := -1; dj <= 1; dj++ {
			tmp, found = squiddict[Point{point.i + di, point.j + dj}]
			if found && tmp[0] > 9 && tmp[1] == 1 {
				energy += 1
			}
		}
	}
	return energy
}

func main() {
	days := 100
	// to high: 1662
	// to high: 1624

	data, err := os.ReadFile("day11_test.txt")
	if err != nil {
		log.Fatal(err)
	}

	squidmap_1 := strings.Split(strings.TrimSpace(string(data)), "\n")
	squiddict := make(map[Point][]int, 0)
	for i, squids := range squidmap_1 {
		for j, squid := range squids {
			squiddict[Point{i, j}] = []int{int(squid - '0'), 0}
		}
	}

	fmt.Println(squiddict)
	flashes := 0
	for day := 1; day <= days; day++ {
		for _, squid := range squiddict {
			squid[0]++
			if squid[0] > 9 {
				squid[1] = 1
			}
		}
		changed := true
		for changed {
			changed = false
			for k, _ := range squiddict {
				if squiddict[k][1] == 0 {
					increase := increaseSquid(squiddict, k)
					if increase > 0 {
						squiddict[k][0] += increase
						if squiddict[k][0] > 9 {
							squiddict[k][1] = 1
						}
						changed = true
					}
				}
			}
			for k, _ := range squiddict {
				if squiddict[k][1] == 1 {
					squiddict[k][1] = 2
					flashes += 1
				}
			}
		}
		fmt.Println(day)
		fmt.Printf("%# v \n", pretty.Formatter(squiddict))
		fmt.Println(flashes)
		for k, _ := range squiddict {
			if squiddict[k][0] > 9 {
				squiddict[k][0] = 0
			}
			if squiddict[k][1] == 2 {
				squiddict[k][1] = 0
			}
		}
	}
	if flashes == 1656 {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
	fmt.Println(flashes == 1656)
	// var basins []int64
	// for i, squids := range squidmap {
	// 	next_row := max_row
	// 	if i+1 < len(squidmap) {
	// 		next_row = squidmap[i+1]
	// 	}
	// 	last_row := max_row
	// 	if i-1 >= 0 {
	// 		last_row = squidmap[i-1]
	// 	}
	// 	for j, squid := range squids {
	// 		next_col := int64(10)
	// 		if j+1 < len(squids) {
	// 			next_col = squids[j+1]
	// 		}
	// 		last_col := int64(10)
	// 		if j-1 >= 0 {
	// 			last_col = squids[j-1]
	// 		}
	// 		if next_row[j] > squid && last_row[j] > squid && next_col > squid && last_col > squid {
	// 			risk += 1 + squidmap[i][j]
	// 			sizeOfBasin := int64(0)
	// 			sizeOfBasin, done = findBasin(squidmap, Point{i, j}, done)
	// 			basins = append(basins, sizeOfBasin)
	// 		}
	// 	}
	// }
	// sort.Slice(basins, func(i, j int) bool { return basins[i] < basins[j] })
	// part2 := int64(1)
	// for _, basin := range basins[len(basins)-3:] {
	// 	part2 *= basin
	// }
	// fmt.Println(risk, part2)
}
