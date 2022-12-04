package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/docopt/docopt-go"
)

func printFreqCount(template_pairs map[[2]string]int64, template []string) {
	freq_count := make(map[string]int64, 0)
	for pair, v := range template_pairs {
		freq_count[pair[0]] += v
		freq_count[pair[1]] += v
	}
	freq_count[template[0]] += 1
	freq_count[template[len(template)-1]] += 1
	const MaxInt = int64(^uint(0) >> 1)
	const MinInt = -MaxInt - 1
	max_freq := MinInt
	min_freq := MaxInt
	max := "A"
	min := "B"
	_ = min
	_ = max
	for letter, freq := range freq_count {
		switch true {
		case freq > max_freq:
			max = letter
			max_freq = freq
		case freq < min_freq:
			min = letter
			min_freq = freq
		}
	}
	fmt.Println(freq_count, max_freq/2, min_freq/2, max_freq/2-min_freq/2)
}

func main() {
	usage := `2021 Day 14 1.0

Usage:
  day14 <x>...
  day14 -h | --help
  day14 --version

Options:
  -h --help     Show this screen.
  --version     Show version.
	<x>...        number of loops`

	arguments, _ := docopt.ParseArgs(usage, os.Args[1:], "1.0")

	days_to_print := make([]int64, 0)
	for _, day := range arguments["<x>"].([]string) {
		day_int, _ := strconv.ParseInt(day, 0, 64)
		days_to_print = append(days_to_print, day_int)
	}

	data, err := os.ReadFile("2021/day14/day14.txt")
	// data, err := os.ReadFile("2021/day14/day14_ex1.txt")
	if err != nil {
		log.Fatal(err)
	}

	template_and_replacements := strings.Split(strings.TrimSpace(string(data)), "\n\n")

	rules := make(map[[2]string]string, 0)
	for _, replacement_rule := range strings.Split(template_and_replacements[1], "\n") {
		tmp := strings.Split(replacement_rule, " -> ")
		rules[[2]string{string(tmp[0][0]), string(tmp[0][1])}] = tmp[1]
	}

	template := strings.Split(template_and_replacements[0], "")
	template_pairs := make(map[[2]string]int64, 0)
	for i := range template {
		if i < len(template)-1 {
			template_pairs[[2]string{template[i], template[i+1]}] += 1
		}
	}
	fmt.Println(rules)
	fmt.Println(template_pairs)

	steps := int64(0)

	for len(days_to_print) > 0 {
		newPoly := make(map[[2]string]int64, 0)
		for k, v := range template_pairs {
			newPoly[k] = v
		}
		for pair := range template_pairs {
			newPoly[[2]string{pair[0], rules[pair]}] += template_pairs[pair]
			newPoly[[2]string{rules[pair], pair[1]}] += template_pairs[pair]
			newPoly[[2]string{pair[0], pair[1]}] -= template_pairs[pair]
		}
		template_pairs = newPoly
		if steps+1 == days_to_print[0] {
			days_to_print = days_to_print[1:]
			printFreqCount(template_pairs, template)
		}
		steps++
	}

}
