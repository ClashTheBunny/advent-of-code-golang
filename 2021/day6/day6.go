package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	days := int64(256)

	data, err := os.ReadFile("day6.txt")
	if err != nil {
		log.Fatal(err)
	}
	initial_state := strings.Split(strings.TrimSpace(string(data)), ",")

	// initial_state := strings.Split("3,4,3,1,2", ",")
	state_map := make(map[int64]int64, 9)
	for state := range initial_state {
		number, _ := strconv.ParseInt(initial_state[state], 0, 64)
		state_map[number] += 1
	}

	for days > 0 {
		for fish := int64(0); fish < 9; fish++ {
			state_map[fish-1] = state_map[fish]
		}
		state_map[6] += state_map[-1]
		state_map[8] = state_map[-1]
		state_map[-1] = 0
		days--
		fmt.Println(state_map)
	}
	total := int64(0)
	for fish := range state_map {
		total += state_map[fish]
	}
	fmt.Println(state_map, total)

}
