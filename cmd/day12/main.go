package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/relarson/advent-of-code-2021.git/pkg"
)

const (
	NOT_VISITABLE int = iota
	REPEAT
	VISITABLE
)

func main() {
	fmt.Println("Problem 1: " + strconv.Itoa(problem1()))
	fmt.Println("Problem 2: " + strconv.Itoa(problem2()))
}

func problem1() int {
	lines, err := pkg.ReadLines("cmd/day12/input.txt")
	if err != nil {
		println(err.Error())
		return -1
	}

	adj := buildAdjList(lines)

	paths := allPathsDFS("start", adj, nil, 0)

	return len(paths)
}

func problem2() int {
	lines, err := pkg.ReadLines("cmd/day12/input.txt")
	if err != nil {
		println(err.Error())
		return -1
	}

	adj := buildAdjList(lines)

	paths := allPathsDFS("start", adj, nil, 1)

	return len(paths)
}

func allPathsDFS(start string, adj map[string][]string, visitedSmallCaves []string, repeatsAllowed int) [][]string {
	var allPaths [][]string

	if start == "end" {
		return nil
	}

	for _, neighbor := range adj[start] {
		if neighbor == "end" {
			// start to end is the path to add
			allPaths = append(allPaths, []string{start, neighbor})
			continue
		}
		visitibility := isVisitable(neighbor, visitedSmallCaves)
		if visitibility == REPEAT && repeatsAllowed == 0 {
			continue
		}
		if visitibility != NOT_VISITABLE {
			newRepeats := repeatsAllowed
			if visitibility == REPEAT {
				newRepeats--
			}
			newVisited := append(visitedSmallCaves, neighbor)
			routesFromNeighbor := allPathsDFS(neighbor, adj, newVisited, newRepeats)
			for _, route := range routesFromNeighbor {
				path := []string{start}
				allPaths = append(allPaths, append(path, route...))
			}
		}
	}

	return allPaths
}

func isVisitable(node string, visitedSmallCaves []string) int {
	if visitedSmallCaves == nil {
		return VISITABLE
	}

	if node == "start" || node == "end" {
		return NOT_VISITABLE
	}

	lcVersion := strings.ToLower(node)

	if lcVersion != node {
		return VISITABLE
	}

	for _, visited := range visitedSmallCaves {
		if visited == node {
			return REPEAT
		}
	}

	return VISITABLE
}

func buildAdjList(edges []string) map[string][]string {
	adj := make(map[string][]string, 0)

	for _, edge := range edges {
		nodes := strings.Split(edge, "-")
		for i := 0; i < 2; i++ {
			current, ok := adj[nodes[i]]
			if ok {
				adj[nodes[i]] = append(current, nodes[1-i])
			} else {
				adj[nodes[i]] = []string{nodes[1-i]}
			}
		}
	}

	return adj
}
