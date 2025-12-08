package union

type UnionFind struct {
	parent []int
}

func NewUnionFind(size int) *UnionFind {
	uf := &UnionFind{
		parent: make([]int, size),
	}
	for i := range uf.parent {
		uf.parent[i] = i
	}
	return uf
}

func (uf *UnionFind) Find(x int) int {
	if x == uf.parent[x] {
		return x
	}
	uf.parent[x] = uf.Find(uf.parent[x])
	return uf.parent[x]
}

// Typically you'd also add Union:
func (uf *UnionFind) Union(x, y int) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)
	if rootX != rootY {
		uf.parent[rootX] = rootY
	}
}
