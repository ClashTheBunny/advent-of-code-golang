package main

import (
	"fmt"
	"log"
	"os"
  "regexp"
	"strings"
	"strconv"
)

type crateState struct {
  columns [][]rune
}

type moveList struct {
  moves []move
}

type move struct {
  quantity int
  source int
  dest int
}

func ReverseSlice[T any](s []T) {
  first := 0
  last := len(s) - 1
  for first < last {
    s[first], s[last] = s[last], s[first]
    first++
    last--
  }
}

func parseData(input string) (current crateState, moves moveList, err error) {
  tempSplit := strings.Split(input, "\n\n")
  crateStateString, moveListString := tempSplit[0], tempSplit[1]
  
  crateStateList := strings.Split(crateStateString, "\n")

  var initialState crateState
  initialState.columns = make([][]rune, 0)
  for column := 1; column < len(crateStateList[len(crateStateList)-1]); column+=4 {
    initialState.columns = append(initialState.columns, make([]rune, 0))
    for line := range crateStateList {
      if line == len(crateStateList) - 1 {
        break
      }
      if rune(crateStateList[line][column]) == ' ' {
        continue
      }
      initialState.columns[len(initialState.columns) -1 ] = append(initialState.columns[len(initialState.columns) - 1], rune(crateStateList[line][column]))
    }
  }
  allMoves := moveList{}
  re := regexp.MustCompile(`move (.*) from (.*) to (.*)`)
  for _, line := range strings.Split(moveListString, "\n") {
    moveMatches := re.FindAllStringSubmatch(line, -1)
    if len(moveMatches) > 0 {
      quantity, source, dest := moveMatches[0][1], moveMatches[0][2], moveMatches[0][3]
      _, _, _ = quantity, source, dest
      quantityi, _ := strconv.Atoi(quantity)
      sourcei, _:= strconv.Atoi(source)
      desti, _ := strconv.Atoi(dest)
      allMoves.moves = append(allMoves.moves, move{
        quantity: quantityi,
        source: sourcei,
        dest: desti,
      })
      
    }
  }

  return initialState, allMoves, nil
}


func main() {
	data, err := os.ReadFile("2022/day5/day5.txt")
	if err != nil {
		log.Fatal(err)
	}
	data_s := string(data)

  state, moves, _ := parseData(data_s)

  for _, move := range moves.moves {
    pop := make([]rune, 0)
    pop = append(pop, state.columns[move.source-1][0:move.quantity]...)
    // ReverseSlice(pop)
    leave := make([]rune, 0)
    leave = append(leave,state.columns[move.source-1][move.quantity:]...)
    state.columns[move.source-1] = leave
    pop = append(pop, state.columns[move.dest-1]...)
    state.columns[move.dest-1] = pop
    _ = pop
  }

  for col := range state.columns {
    fmt.Printf("%v", string(state.columns[col][0]))
  }
}
