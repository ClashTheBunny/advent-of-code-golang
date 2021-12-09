package main

import (
	"fmt"
	"log"
	"math"
	"math/bits"
	"os"
	"strings"
)

func min(theArray []uint8) uint8 {
	min := uint8(128)
	for _, item := range theArray {
		if item < min {
			min = item
		}
	}
	return min
}

func digitToBitDigit(digit string) uint8 {
	bitDigit := uint8(0b0)
	for _, char := range digit {
		switch char {
		case 'a':
			bitDigit = bitDigit | 0b1000000
		case 'b':
			bitDigit = bitDigit | 0b100000
		case 'c':
			bitDigit = bitDigit | 0b10000
		case 'd':
			bitDigit = bitDigit | 0b1000
		case 'e':
			bitDigit = bitDigit | 0b100
		case 'f':
			bitDigit = bitDigit | 0b10
		case 'g':
			bitDigit = bitDigit | 0b1
		}
	}
	return bitDigit
}

func main() {
	data, err := os.ReadFile("day8_test.txt")
	if err != nil {
		log.Fatal(err)
	}

	in_outs := strings.Split(strings.TrimSpace(string(data)), "\n")

	ins := make([]string, len(in_outs))
	outs := make([]string, len(in_outs))

	unique := 0
	grandTotal := 0
	for in_out := range in_outs {
		this_in_out := strings.Split(in_outs[in_out], "|")
		ins[in_out] = strings.TrimSpace(this_in_out[0])
		outs[in_out] = strings.TrimSpace(this_in_out[1])

		split_ins := strings.Split(ins[in_out], " ")

		digits := make([]uint8, 10)
		for min(digits) == 0 {
			for _, digit := range split_ins {
				bitDigit := digitToBitDigit(digit)
				switch bits.OnesCount8(bitDigit) {
				case 2: // It's a one
					digits[1] = bitDigit
				case 3: // It's a seven
					digits[7] = bitDigit
				case 4: // It's a four
					digits[4] = bitDigit
				case 7: // It's an 8
					digits[8] = bitDigit
				case 6:
					if (digits[1] > 0) && (bitDigit&digits[1] != digits[1]) {
						digits[6] = bitDigit
					} else if (digits[4] > 0 && digits[1] > 0) && (bitDigit&digits[1] == digits[1]) && (bitDigit&digits[4] == digits[4]) {
						digits[9] = bitDigit
					} else if (digits[4] > 0 && digits[1] > 0) && (bitDigit&digits[1] == digits[1]) && (bitDigit&digits[4] != digits[4]) {
						digits[0] = bitDigit
					} else {
					}
				case 5:
					if (digits[7] > 0) && (bitDigit&digits[7] == digits[7]) {
						digits[3] = bitDigit
					} else if (digits[6] > 0 && digits[3] > 0) && bits.OnesCount8(bitDigit&digits[6]) == 5 {
						digits[5] = bitDigit
					} else if (digits[6] > 0 && digits[3] > 0) && bits.OnesCount8(bitDigit&digits[6]) == 4 {
						digits[2] = bitDigit
					} else {
						// fmt.Printf("not found (%d): %08b\n", i, bitDigit)
						// fmt.Printf("%08b\n", digits)
					}
				default:
					log.Fatal("woops")
				}
			}
		}
		digitMap := make(map[uint8]int)
		for i, digit := range digits {
			digitMap[digit] = i
		}
		// fmt.Printf("%08b\n", digits)
		split_outs := strings.Split(outs[in_out], " ")
		total := 0
		for i, digit := range split_outs {
			bitDigit := digitToBitDigit(digit)
			multiplier := int(math.Pow(10, float64(3-i)))
			total += multiplier * digitMap[bitDigit]
			switch digitMap[bitDigit] {
			case 1, 4, 7, 8:
				unique += 1
			}
		}
		// fmt.Println(this_in_out, total)
		grandTotal += total
	}

	fmt.Println(unique, grandTotal)
}
