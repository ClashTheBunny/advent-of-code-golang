package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sort"
)

func main() {
	data, err := os.ReadFile("2022/day1/day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	data_s := strings.TrimSpace(string(data))
	backpacks := strings.Split(data_s, "\n\n")

	// Part 1
  backpackTotals := make([]int64,0)
  maxBackpack := int64(0)
	for _, backpack := range backpacks {
    snacks := strings.Split(backpack,"\n")
    backpackTotal := int64(0)
    for _, calorieSnack := range snacks {
			current, err := strconv.Atoi(calorieSnack)
			if err != nil {
				log.Fatal(err)
			}
      backpackTotal+=int64(current)
		}
    log.Print(backpackTotal)
    backpackTotals = append(backpackTotals, backpackTotal)
    if backpackTotal > maxBackpack {
      maxBackpack = backpackTotal
    }
	}
  sort.Slice(backpackTotals, func(i, j int) bool {
    return backpackTotals[i] > backpackTotals[j]
  })
	fmt.Println(backpackTotals, maxBackpack)
  // Part 2
  fmt.Print(backpackTotals[0:3])
  topThree := 0
  for _, caloriesSnacks := range backpackTotals[0:3] {
    topThree += int(caloriesSnacks)
  }
  fmt.Print(topThree)
}
