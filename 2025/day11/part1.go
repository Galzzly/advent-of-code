package main

import "strings"

func doPartOne(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	// Build the adjacency list graph
	devices := make(map[string][]string)
	for _, line := range lines {
		s := strings.Fields(line)
		device := strings.TrimSuffix(s[0], ":")
		devices[device] = s[1:]
	}

	// Count all paths from "you" to "out"
	pathCount := countPaths(devices, "you", "out", make(map[string]bool))

	return pathCount
}

// countPaths uses DFS to count all paths from start to end
func countPaths(graph map[string][]string, current, target string, visited map[string]bool) int {
	// If we reached the target, we found a path
	if current == target {
		return 1
	}

	// Mark current node as visited
	visited[current] = true
	defer func() { visited[current] = false }() // Unmark when backtracking

	totalPaths := 0

	// Explore all neighbors
	for _, neighbor := range graph[current] {
		// Skip if already visited (avoid cycles)
		if visited[neighbor] {
			continue
		}

		// Recursively count paths from neighbor to target
		totalPaths += countPaths(graph, neighbor, target, visited)
	}

	return totalPaths
}
