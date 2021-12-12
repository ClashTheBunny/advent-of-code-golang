package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Graph struct {
	verticies   map[string]int
	names       map[int]string
	adjacencies [][]int
}

func buildGraph(graphStrings []string) Graph {

	graph := Graph{make(map[string]int, 0), make(map[int]string, 0), make([][]int, 0)}

	for _, vertex := range graphStrings {
		vertex := strings.Split(vertex, "-")
		for _, vert := range vertex {
			_, seen_vertex := graph.verticies[vert]
			if !seen_vertex {
				current_len := len(graph.verticies)
				graph.verticies[vert] = current_len
				graph.names[current_len] = vert
				graph.adjacencies = append(graph.adjacencies, make([]int, 0))
			}
		}
		graph.adjacencies[graph.verticies[vertex[0]]] = append(graph.adjacencies[graph.verticies[vertex[0]]], graph.verticies[vertex[1]])
		graph.adjacencies[graph.verticies[vertex[1]]] = append(graph.adjacencies[graph.verticies[vertex[1]]], graph.verticies[vertex[0]])
	}
	return graph
}

func countPaths(graph Graph, start int, end int, visited []bool, burned_small bool, path []int) int64 {
	local_visited := make([]bool, len(visited))
	copy(local_visited, visited)
	pathCount := int64(0)
	if graph.names[start] != strings.ToUpper(graph.names[start]) {
		local_visited[start] = true
	}
	if start == end {
		pathCount++
	} else {
		for _, vertex := range graph.adjacencies[start] {
			if !local_visited[vertex] {
				path = append(path, vertex)
				pathCount += countPaths(graph, vertex, end, local_visited, burned_small, path)
			} else if !burned_small && "start" != graph.names[vertex] {
				path = append(path, vertex)
				pathCount += countPaths(graph, vertex, end, local_visited, true, path)
			}
		}
	}
	return pathCount
}

func main() {

	data, err := os.ReadFile("day12.txt")
	if err != nil {
		log.Fatal(err)
	}

	graphStrings := strings.Split(strings.TrimSpace(string(data)), "\n")
	fmt.Println(graphStrings)

	graph := buildGraph(graphStrings)
	fmt.Println(graph)
	visited := make([]bool, len(graph.verticies))
	pathCount1 := countPaths(graph, graph.verticies["start"], graph.verticies["end"], visited, true, []int{})
	pathCount2 := countPaths(graph, graph.verticies["start"], graph.verticies["end"], visited, false, []int{})

	fmt.Println(graphStrings)
	fmt.Println(graph)
	fmt.Println(pathCount1, pathCount2)
}
