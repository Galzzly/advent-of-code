package main

import "strings"

func doPartTwo(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	// Build the adjacency list graph
	devices := make(map[string][]string)
	for _, line := range lines {
		s := strings.Fields(line)
		device := strings.TrimSuffix(s[0], ":")
		devices[device] = s[1:]
	}

	// Count paths from "svr" to "out" that visit both "dac" and "fft"
	memo := make(map[string]int)
	pathCount := countPathsWithRequiredMemo(devices, "svr", "out", []string{"dac", "fft"}, make(map[string]bool), make(map[string]bool), memo)

	return pathCount
}

// countPathsWithRequiredMemo counts paths with memoization
func countPathsWithRequiredMemo(graph map[string][]string, current, target string, required []string, visited map[string]bool, foundRequired map[string]bool, memo map[string]int) int {
	// If we reached the target
	if current == target {
		// Check if all required nodes were visited
		for _, req := range required {
			if !foundRequired[req] {
				return 0
			}
		}
		return 1
	}

	// Create memo key from current node and which required nodes have been found
	memoKey := current
	for _, req := range required {
		if foundRequired[req] {
			memoKey += "+" + req
		}
	}

	// Check memo
	if val, exists := memo[memoKey]; exists {
		return val
	}

	// Mark current node as visited
	visited[current] = true
	defer func() { visited[current] = false }()

	// Check if current node is a required node
	wasFound := make(map[string]bool)
	for _, req := range required {
		if current == req && !foundRequired[req] {
			foundRequired[req] = true
			wasFound[req] = true
		}
	}
	defer func() {
		for req := range wasFound {
			foundRequired[req] = false
		}
	}()

	totalPaths := 0

	// Explore all neighbors
	for _, neighbor := range graph[current] {
		if visited[neighbor] {
			continue
		}

		totalPaths += countPathsWithRequiredMemo(graph, neighbor, target, required, visited, foundRequired, memo)
	}

	memo[memoKey] = totalPaths
	return totalPaths
}
