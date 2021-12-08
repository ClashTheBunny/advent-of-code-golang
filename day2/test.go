package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("day2.txt")
	if err != nil {
		log.Fatal(err)
	}
	data_s := strings.TrimSpace(string(data))
	commands := strings.Split(data_s, "\n")

	// Part 1
	depth, distance := 0, 0
	for command := range commands {
		dir_mag := strings.Split(commands[command], " ")
		direction, magnitude := dir_mag[0], dir_mag[1]
		magnitude_i, err := strconv.Atoi(magnitude)
		if err != nil {
			log.Fatal(err)
		}
		switch direction {
		case "down":
			depth += magnitude_i
		case "up":
			depth -= magnitude_i
		case "forward":
			distance += magnitude_i
		default:
			log.Fatal(direction)
		}
	}
	fmt.Println(depth, distance, depth*distance)

	// Part 2
	depth, distance, aim := 0, 0, 0
	for command := range commands {
		dir_mag := strings.Split(commands[command], " ")
		direction, magnitude := dir_mag[0], dir_mag[1]
		magnitude_i, err := strconv.Atoi(magnitude)
		if err != nil {
			log.Fatal(err)
		}
		switch direction {
		case "down":
			aim += magnitude_i
		case "up":
			aim -= magnitude_i
		case "forward":
			distance += magnitude_i
			depth += magnitude_i * aim
		default:
			log.Fatal(direction)
		}
	}
	fmt.Println(depth, distance, depth*distance)
}
