package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type FuncInt64Int64 func(int64) int64

func memorized(fn FuncInt64Int64) FuncInt64Int64 {
	cache := make(map[int64]int64)

	return func(input int64) int64 {
		if val, found := cache[input]; found {
			return val
		}

		tmp := fn(input)
		cache[input] = tmp
		return tmp
	}
}

func Abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func crab_cost_part1(x_distance_traveled int64) int64 {
	return x_distance_traveled
}

func main() {

	var crabCostPart2mem FuncInt64Int64

	crabCostPart2mem = memorized(func(n int64) int64 {
		if n <= 1 {
			return n
		}
		return n + crabCostPart2mem(n-1)
	})

	// data, err := os.ReadFile("day7_example.txt")
	data, err := os.ReadFile("day7.txt")
	if err != nil {
		log.Fatal(err)
	}
	crab_list := strings.Split(strings.TrimSpace(string(data)), ",")
	crab_ints := []int64{}

	mean := int64(0)
	for crab := range crab_list {
		crab_int, _ := strconv.ParseInt(crab_list[crab], 0, 64)
		crab_ints = append(crab_ints, crab_int)
		mean += crab_int
	}
	sort.Slice(crab_ints, func(i, j int) bool { return crab_ints[i] < crab_ints[j] })
	crab_number := int64(len(crab_ints) / 2)
	mean /= crab_number
	median := int64(0)
	if crab_number%2 == 0 {
		median = crab_ints[crab_number]
	} else {
		median = (crab_ints[crab_number-1] + crab_ints[crab_number]) / 2
	}

	total := int64(0)
	for crab := range crab_list {
		total += crab_cost_part1(Abs(crab_ints[crab] - median))
	}
	fmt.Println("Part 1:", mean, median, total, crab_ints[0], crab_ints[len(crab_ints)-1])

	min := int64(100000000000000000)
	min_loc := int64(0)
	for x := crab_ints[len(crab_ints)-1]; x >= crab_ints[0]; x-- {
		total = int64(0)
		for crab := range crab_list {
			total += crabCostPart2mem(Abs(crab_ints[crab] - x))
		}
		if total < min {
			min = total
			min_loc = x
		}
		fmt.Println(x, total, min, min_loc, total-min)
	}
	fmt.Println("Part 2:", mean, median, total, min, min_loc)
	// Too low: 464
	// Too low: 951
}
