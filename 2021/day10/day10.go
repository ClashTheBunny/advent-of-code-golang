package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	data, err := os.ReadFile("day10.txt")
	if err != nil {
		log.Fatal(err)
	}

	codeLines := strings.Split(strings.TrimSpace(string(data)), "\n")
	score := 0
	fix_scores := []int64{}
	for _, line := range codeLines {
		opens := []rune{}
	LineLoop:
		for _, char := range line {
			x := char
			switch char {
			case '{', '(', '<', '[':
				opens = append(opens, char)
			case '}', ')', '>', ']':
				x, opens = opens[len(opens)-1], opens[:len(opens)-1]
			}
			switch char {
			case '>':
				if x != '<' {
					score += 25137
					opens = []rune{}
					break LineLoop
				}
			case '}':
				if x != '{' {
					score += 1197
					opens = []rune{}
					break LineLoop
				}
			case ']':
				if x != '[' {
					score += 57
					opens = []rune{}
					break LineLoop
				}
			case ')':
				if x != '(' {
					score += 3
					opens = []rune{}
					break LineLoop
				}
			}
		}
		if len(opens) > 0 {
			fix_score := int64(0)
			for i := range opens {
				switch opens[len(opens)-1-i] {
				case '<':
					fix_score = fix_score*5 + 4
				case '{':
					fix_score = fix_score*5 + 3
				case '[':
					fix_score = fix_score*5 + 2
				case '(':
					fix_score = fix_score*5 + 1
				}
			}
			fix_scores = append(fix_scores, fix_score)
		}
	}
	fmt.Println(score)
	sort.Slice(fix_scores, func(i, j int) bool { return fix_scores[i] < fix_scores[j] })
	fmt.Println(fix_scores[len(fix_scores)/2])
}
