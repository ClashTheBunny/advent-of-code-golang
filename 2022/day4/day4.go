package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("2022/day4/day4.txt")
	if err != nil {
		log.Fatal(err)
	}
	data_s := strings.TrimSpace(string(data))

  total := 0
  anyOverlap := 0
  re := regexp.MustCompile(`(.*)-(.*),(.*)-(.*)`)
  for _, elfPair := range strings.Split(data_s, "\n") {
    matches := re.FindAllStringSubmatch(elfPair, -1)[0]
    e1start, _ := strconv.Atoi(matches[1])
    e1finish, _ := strconv.Atoi(matches[2])
    e2start, _ := strconv.Atoi(matches[3])
    e2finish, _ := strconv.Atoi(matches[4])
    if (e1start <= e2start && e1finish >= e2finish) || ( e2start <= e1start && e2finish >= e1finish ) {
        total++
    }
    if ! ( e2finish < e1start || e1finish < e2start ) {
      anyOverlap++
      fmt.Printf("%v\n", elfPair)
    }

  }

  // Part 1
  // 425 too low
  // 697 too high
  // 609 too high

  // Part 2
  // 936 too high
  // 694 too low
  fmt.Print(total, anyOverlap)
}
