package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Strategy struct {
  Their int
  Our   int
}

const (
  loss = 0
  tie = 3
  win = 6
)

func returnValid(x, y int) int {
  if x < 1 || x >3 {
    return y
  }
  return x
}

func main() {
	data, err := os.ReadFile("2022/day2/day2.txt")
	if err != nil {
		log.Fatal(err)
	}
	data_s := strings.TrimSpace(string(data))
  // fmt.Printf("%v", data_s)
  // fmt.Printf("\n")
	strategies := strings.Split(data_s, "\n")

  score := 0
  strategyList := make([]*Strategy, 0)
	// Part 1
	for _, strategy := range strategies {
    strategySteps := strings.Split(strategy," ")
    thisStrategy := &Strategy{
      Their: int(strategySteps[0][0]) - 'A' + 1,
      Our: int(strategySteps[1][0]) - 'X' + 1,
    }
    strategyList = append(strategyList, thisStrategy)
    switch result := thisStrategy.Our - thisStrategy.Their; result {
    case -1, 2:
      score+=0
    case 1, -2:
      score+=6
    case 0:
      score+=3
    }
    score += thisStrategy.Our
    // fmt.Printf("%v", strategyList[len(strategyList)-1])
	}
  fmt.Printf("%v\n", score)


  // Part 2
  score = 0
  for _, desiredResult := range strategyList {
    switch desiredResult.Our {
    case 1:
      //loss
      score += returnValid(desiredResult.Their - 1, desiredResult.Their + 2)
    case 2:
      //tie
      score += 3 + desiredResult.Their
    case 3:
      //win
      score += 6 + returnValid(desiredResult.Their - 2, desiredResult.Their + 1)
    }
  }
  fmt.Printf("%v\n", score)
}
