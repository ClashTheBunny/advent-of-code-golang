package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("day3.txt")
	if err != nil {
		log.Fatal(err)
	}
	data_s := strings.TrimSpace(string(data))
	commands := strings.Split(data_s, "\n")

	epsilon, gamma := "", ""

	for bit := range commands[0] {
		total := 0
		for line := range commands {
			if commands[line][bit] == 49 {
				total++
			}
		}
		if total > len(commands)/2 {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}
	gammaval, err := strconv.ParseUint(gamma, 2, 16)
	epsilonval, err := strconv.ParseUint(epsilon, 2, 16)
	fmt.Println("Gamma: ", gamma, gammaval, " Epsilon: ", epsilon, epsilonval, "Power:", gammaval*epsilonval)

	o2 := make([]string, len(commands))
	co2 := make([]string, len(commands))

	copy(o2, commands)
	copy(co2, commands)

	for bit := range o2[0] {
		totalOneBits := 0
		var goodBit byte
		for line := range o2 {
			if o2[line][bit] == 49 {
				totalOneBits++
			}
		}
		if totalOneBits < len(o2)-totalOneBits {
			goodBit = byte(48)
		} else {
			goodBit = byte(49)
		}
		n := 0
		for line := range o2 {
			if o2[line][bit] == goodBit {
				o2[n] = o2[line]
				n++
			}
		}
		o2 = o2[:n]
		if n == 1 {
			break
		}
	}

	for bit := range co2[0] {
		totalOneBits := 0
		var goodBit byte
		for line := range co2 {
			if co2[line][bit] == 49 {
				totalOneBits++
			}
		}
		if totalOneBits >= len(co2)-totalOneBits {
			goodBit = byte(48)
		} else {
			goodBit = byte(49)
		}
		n := 0
		for line := range co2 {
			if co2[line][bit] == goodBit {
				co2[n] = co2[line]
				n++
			}
		}
		co2 = co2[:n]
		if n == 1 {
			break
		}
	}
	o2_val, err := strconv.ParseUint(o2[0], 2, 16)
	co2_val, err := strconv.ParseUint(co2[0], 2, 16)
	fmt.Println("Total raiting:", co2_val*o2_val)

	// other account
	// too low 5388270
	// too high 5570677
	// too high 5547955

	// main account
	// too high 491022
	// too high 484218
	// too low  350730
}
