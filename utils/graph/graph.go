package graph

import (
	"container/heap"
	"image"
	"math"
)

type Queue[T any] []T

type Graph interface {
	Neighbours(p image.Point) []image.Point
}

type Grid[T any] struct {
	x, y      int
	state     map[image.Point]T
	movements []image.Point
}

func NewGrid[T any](x, y int, movements []image.Point) *Grid[T] {
	state := make(map[image.Point]T)
	return &Grid[T]{
		x:         x,
		y:         y,
		state:     state,
		movements: movements,
	}
}

func (g *Grid[T]) IsValid(x, y int) bool {
	switch {
	case x < 0, x >= g.x, y < 0, y >= g.y:
		return false
	default:
		return true
	}
}

func (g *Grid[T]) SetState(x, y int, state T) {
	if g.IsValid(x, y) {
		g.state[image.Point{x, y}] = state
	}
}

func (g *Grid[T]) GetState(p image.Point) T {
	return g.state[p]
}

func (g *Grid[T]) Neighbours(p image.Point) (res []image.Point) {
	for _, m := range g.movements {
		np := p.Add(m)
		if g.IsValid(np.X, np.Y) {
			res = append(res, np)
		}
	}
	return
}

func (q *Queue[T]) Put(x T) {
	*q = append(*q, x)
}

func (q *Queue[T]) Get() T {
	ret := (*q)[0]
	*q = (*q)[1:]
	return ret
}

func (q *Queue[T]) Empty() bool {
	return len(*q) == 0
}

func Search(g Graph, s, e image.Point) (res []image.Point) {
	var queue Queue[image.Point]
	queue.Put(s)

	from := map[image.Point]*image.Point{}
	from[s] = nil

	for !queue.Empty() {
		current := queue.Get()
		if current == e {
			break
		}
		for _, p := range g.Neighbours(current) {
			if _, ok := from[p]; !ok {
				queue.Put(p)
				from[p] = &current
			}
		}
	}
	res = []image.Point{e}
	for p := from[e]; p != nil; p = from[*p] {
		res = append(res, *p)
	}
	return
}

// PriorityQueueItem represents an item in the priority queue
type PriorityQueueItem struct {
	Node     string
	Distance int
	Index    int
}

// PriorityQueue implements heap.Interface for Dijkstra's algorithm
type PriorityQueue []*PriorityQueueItem

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Distance < pq[j].Distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*PriorityQueueItem)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.Index = -1
	*pq = old[0 : n-1]
	return item
}

// Dijkstra finds the shortest path from start to end node
// graph is a map of node -> list of connected nodes
// Returns the distance to the end node, or -1 if no path exists
func Dijkstra(graph map[string][]string, start, end string) int {
	distances := make(map[string]int)

	// Initialize distances for all nodes in the graph
	for node := range graph {
		distances[node] = math.MaxInt32
	}
	distances[start] = 0

	// Also ensure end node exists in distances even if it has no outgoing edges
	if _, exists := distances[end]; !exists {
		distances[end] = math.MaxInt32
	}

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &PriorityQueueItem{
		Node:     start,
		Distance: 0,
	})

	visited := make(map[string]bool)

	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*PriorityQueueItem)

		if current.Node == end {
			return current.Distance
		}

		if visited[current.Node] {
			continue
		}
		visited[current.Node] = true

		for _, neighbor := range graph[current.Node] {
			if visited[neighbor] {
				continue
			}

			newDist := current.Distance + 1 // Each edge has weight 1

			if newDist < distances[neighbor] {
				distances[neighbor] = newDist
				heap.Push(&pq, &PriorityQueueItem{
					Node:     neighbor,
					Distance: newDist,
				})
			}
		}
	}

	// No path found
	return -1
}

// DijkstraWithPath finds the shortest path and returns both distance and the path
func DijkstraWithPath(graph map[string][]string, start, end string) (int, []string) {
	distances := make(map[string]int)
	previous := make(map[string]string)

	// Initialize distances for all nodes in the graph
	for node := range graph {
		distances[node] = math.MaxInt32
	}
	distances[start] = 0

	// Also ensure end node exists in distances even if it has no outgoing edges
	if _, exists := distances[end]; !exists {
		distances[end] = math.MaxInt32
	}

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &PriorityQueueItem{
		Node:     start,
		Distance: 0,
	})

	visited := make(map[string]bool)

	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*PriorityQueueItem)

		if current.Node == end {
			// Reconstruct path
			path := []string{}
			node := end
			for node != "" {
				path = append([]string{node}, path...)
				prev, exists := previous[node]
				if !exists {
					break
				}
				node = prev
			}
			return current.Distance, path
		}

		if visited[current.Node] {
			continue
		}
		visited[current.Node] = true

		for _, neighbor := range graph[current.Node] {
			if visited[neighbor] {
				continue
			}

			newDist := current.Distance + 1

			if newDist < distances[neighbor] {
				distances[neighbor] = newDist
				previous[neighbor] = current.Node
				heap.Push(&pq, &PriorityQueueItem{
					Node:     neighbor,
					Distance: newDist,
				})
			}
		}
	}

	return -1, nil
}
