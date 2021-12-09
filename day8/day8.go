package main

import (
	"fmt"
	"log"
	"math"
	"math/bits"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("day8.txt")
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
		zero := uint8(0b0)
		one := uint8(0b0)
		two := uint8(0b0)
		three := uint8(0b0)
		four := uint8(0b0)
		five := uint8(0b0)
		six := uint8(0b0)
		seven := uint8(0b0)
		eight := uint8(0b0)
		nine := uint8(0b0)
		// for digitIndex, digit := range split_ins {
		i := 0
		for i < 10 {
			for _, digit := range split_ins {
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
				switch bits.OnesCount8(bitDigit) {
				case 2: // It's a one
					one = bitDigit
				case 3: // It's a seven
					seven = bitDigit
				case 4: // It's a four
					four = bitDigit
				case 7: // It's an 8
					eight = bitDigit
				case 6: // It's an 8
					if (one > 0) && (bitDigit&one != one) {
						six = bitDigit
					} else if (four > 0 && one > 0) && (bitDigit&one == one) && (bitDigit&four == four) {
						nine = bitDigit
					} else if (four > 0 && one > 0) && (bitDigit&one == one) && (bitDigit&four != four) {
						zero = bitDigit
					} else {
					}
				case 5: // It's an 8
					if (seven > 0) && (bitDigit&seven == seven) {
						three = bitDigit
					} else if (six > 0 && three > 0) && bits.OnesCount8(bitDigit&six) == 5 {
						five = bitDigit
					} else if (six > 0 && three > 0) && bits.OnesCount8(bitDigit&six) == 4 {
						two = bitDigit
					} else {
						// fmt.Printf("not found (%d): %08b\n", i, bitDigit)
						// fmt.Printf("%08b %08b %08b %08b %08b %08b %08b %08b %08b %08b\n", zero, one, two, three, four, five, six, seven, eight, nine)
					}
				default:
					fmt.Println("woops")
				}
			}
			i++
		}
		fmt.Printf("%08b %08b %08b %08b %08b %08b %08b %08b %08b %08b\n", zero, one, two, three, four, five, six, seven, eight, nine)
		split_outs := strings.Split(outs[in_out], " ")
		total := 0
		for i, digit := range split_outs {
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
			multiplier := int(math.Pow(10, float64(3-i)))
			switch bitDigit {
			case zero:
			case one:
				total += 1 * multiplier
				unique += 1
			case two:
				total += 2 * multiplier
			case three:
				total += 3 * multiplier
			case four:
				total += 4 * multiplier
				unique += 1
			case five:
				total += 5 * multiplier
			case six:
				total += 6 * multiplier
			case seven:
				total += 7 * multiplier
				unique += 1
			case eight:
				total += 8 * multiplier
				unique += 1
			case nine:
				total += 9 * multiplier
			default:
				fmt.Println("not fount(", i, "): ", bitDigit, one, one == bitDigit)
			}
		}
		fmt.Println(this_in_out, total)
		grandTotal += total
	}

	fmt.Println(unique, grandTotal)
}
