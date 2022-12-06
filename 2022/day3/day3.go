package main

import (
	"fmt"
	"log"
	"os"
	// "strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("2022/day3/day3.txt")
	if err != nil {
		log.Fatal(err)
	}
	backpacks := strings.Split(strings.TrimSpace(string(data)), "\n")

	_ = backpacks

  totalPriority := 0
	for _, backpack := range backpacks {
		backpackLength := len(backpack)
		backpackSet := make(map[string]int)
		for j, char := range backpack {
			side := j / (backpackLength/2)
			if side == 0 {
				backpackSet[string(char)] = int(char)
			} else {
        if _, ok := backpackSet[string(char)]; ok {
          if int(char) > 96 {
            totalPriority+=int(char)-96
          } else {
            totalPriority+=int(char)-38
          }
          break
				}
			}
		}
	}

	fmt.Print(totalPriority,"\n")

  backpacksSet := make(map[int]map[string]int)

  totalPriority = 0
  // backpacksLength := len(backpacks)
	for i, backpack := range backpacks {
    group := i / 3
    if i % 3 == 0 {
      backpacksSet[group] = make(map[string]int)
    }
    // fmt.Print(group)
		for _, char := range backpack {
      if backpacksSet[group][string(char)] == i%3 {
        backpacksSet[group][string(char)]++
      }
		}
	}

  for _, groupMap := range backpacksSet {
    for letter, letterFreq := range groupMap {
      if letterFreq == 3 {
        if int(letter[0]) > 96 {
          totalPriority+=int(letter[0])-96
        } else {
          totalPriority+=int(letter[0])-38
        }
        break
      }
    }
  }

	fmt.Print(totalPriority,"\n")
}
