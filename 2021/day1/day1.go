package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("2021/day1/day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	data_s := strings.TrimSpace(string(data))
	sonar_readings := strings.Split(data_s, "\n")

	// Part 1
	increase := 0
	for reading := range sonar_readings {
		if reading > 0 {
			current, err := strconv.Atoi(sonar_readings[reading])
			if err != nil {
				log.Fatal(err)
			}
			last, err := strconv.Atoi(sonar_readings[reading-1])
			if err != nil {
				log.Fatal(err)
			}
			if current > last {
				increase++
			}
		}
	}
	fmt.Println(increase)

	// Part 2
	// Part 1
	increase = 0
	for reading := range sonar_readings {
		if reading > 2 {
			current, err := strconv.Atoi(sonar_readings[reading])
			if err != nil {
				log.Fatal(err)
			}
			last, err := strconv.Atoi(sonar_readings[reading-3])
			if err != nil {
				log.Fatal(err)
			}
			if current > last {
				increase++
			}
		}
	}
	fmt.Println(increase)
}
