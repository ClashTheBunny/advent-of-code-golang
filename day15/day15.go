package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	dijkstra "github.com/RyanCarrier/dijkstra"
)

type Point struct {
	i int
	j int
}

func nodeID(x, y, row_len int) int {
	return x + row_len*y
}

func createMapPart1(data []byte) (*dijkstra.Graph, int, int, int) {
	graph := dijkstra.NewGraph()

	squidmap_1 := strings.Split(strings.TrimSpace(string(data)), "\n")
	squiddict := make(map[Point][]int, 0)
	row_len := len(squidmap_1[0])

	for j, squids := range squidmap_1 {
		for i, squidvalue := range squids {
			thisnodeID := nodeID(i, j, row_len)
			graph.AddVertex(thisnodeID)
			squiddict[Point{i, j}] = []int{int(squidvalue - '0'), 0}
			// fmt.Printf("%d", squiddict[Point{i, j}])
		}
	}
	_ = graph
	for j, squids := range squidmap_1 {
		for i := range squids {
			thisnodeID := nodeID(i, j, row_len)
			for dx := -1; dx <= 1; dx++ {
				for dy := -1; dy <= 1; dy++ {
					if i+dx >= 0 && i+dx < len(squids) {
						if j+dy >= 0 && j+dy < len(squids) {
							if (dx == 0 || dy == 0) && dx != dy {
								otherID := nodeID(i+dx, j+dy, row_len)
								// println(i, j, dx, dy, thisnodeID, otherID, int64(squiddict[Point{i + dx, j + dy}][0]))
								graph.AddArc(thisnodeID, otherID, int64(squiddict[Point{i + dx, j + dy}][0]))
							}
						}
					}
				}
			}
		}
	}
	return graph, len(squidmap_1[0]) - 1, len(squidmap_1) - 1, len(squidmap_1[0])
}

func createMapPart2(data []byte) (*dijkstra.Graph, int, int, int) {
	graph := dijkstra.NewGraph()

	squidmap_1 := strings.Split(strings.TrimSpace(string(data)), "\n")
	squiddict := make(map[Point]int64, 0)
	mult := 5
	orig_xmax, orig_ymax, orig_row_len := len(squidmap_1[0])-1, len(squidmap_1)-1, len(squidmap_1[0])
	xmax, ymax, row_len := len(squidmap_1[0])*mult-1, len(squidmap_1)*mult-1, len(squidmap_1[0])*mult
	_, _ = orig_xmax, orig_ymax

	for j := 0; j <= ymax; j++ {
		for i := 0; i <= xmax; i++ {
			thisnodeID := nodeID(i, j, row_len)
			graph.AddVertex(thisnodeID)
			xmult := int64(i / orig_row_len)
			ymult := int64(j / orig_row_len)
			cost := (int64(squidmap_1[j%orig_row_len][i%orig_row_len]-'0') + xmult + ymult) % 9
			// fmt.Println(thisnodeID, i, j)
			// fmt.Println(xmult, ymult, i%orig_row_len, j%orig_row_len, cost)
			if cost == 0 {
				cost = 9
			}
			squiddict[Point{i, j}] = cost
		}
	}
	// fmt.Println(squiddict)
	for j := 0; j <= ymax; j++ {
		for i := 0; i <= xmax; i++ {
			thisnodeID := nodeID(i, j, row_len)
			for dx := -1; dx <= 1; dx++ {
				for dy := -1; dy <= 1; dy++ {
					if (i+dx >= 0) && (i+dx < row_len) {
						if (j+dy >= 0) && (j+dy < row_len) {
							if (dx == 0 || dy == 0) && dx != dy {
								otherID := nodeID(i+dx, j+dy, row_len)
								cost, found := squiddict[Point{i + dx, j + dy}]
								// println(i, j, dx, dy, row_len, thisnodeID, otherID, cost)
								if !found {
									log.Fatal("cost not found")
								}
								graph.AddArc(thisnodeID, otherID, cost)
							}
						}
					}
				}
			}
		}
	}
	return graph, xmax, ymax, row_len
}

func main() {
	// data, err := os.ReadFile("day15_ex1.txt")
	data, err := os.ReadFile("day15.txt")
	if err != nil {
		log.Fatal(err)
	}

	graph, maxx, maxy, row_len := createMapPart1(data)
	best, err := graph.Shortest(0, nodeID(maxx, maxy, row_len))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Part 1 shortest distance ", best.Distance, " following path ", best.Path)
	graph, maxx, maxy, row_len = createMapPart2(data)

	best, err = graph.Shortest(0, nodeID(maxx, maxy, row_len))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Part 2 shortest distance ", best.Distance, " following path ", best.Path)
}
