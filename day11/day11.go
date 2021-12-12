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

func increaseSquid(squiddict map[Point][]int, point Point) {
	var tmp []int
	var found bool
	for di := -1; di <= 1; di++ {
		for dj := -1; dj <= 1; dj++ {
			tmp, found = squiddict[Point{point.i + di, point.j + dj}]
			if found && tmp[0] <= 9 && tmp[0] > 0 {
				tmp[0] += 1
			}
		}
	}
}

func main() {
	days := 100
	// to high: 1662
	// to high: 1624

	data, err := os.ReadFile("day11.txt")
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
		}
		flashed := true
		for flashed {
			flashed = false
			for k, _ := range squiddict {
				if squiddict[k][0] > 9 {
					increaseSquid(squiddict, k)
					squiddict[k][0] = 0
					flashed = true
					flashes++
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
}
