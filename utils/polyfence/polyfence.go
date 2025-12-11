package polyfence

import "image"

type Polygon []image.Point

type Polyfence struct {
	p Polygon
}

func NewPolyfence(p Polygon) *Polyfence {
	polyfence := &Polyfence{
		p: p.Copy(),
	}
	if len(polyfence.p) > 0 &&
		polyfence.p[0] != polyfence.p[len(polyfence.p)-1] {
		polyfence.p = append(polyfence.p, polyfence.p[0])
	}
	return polyfence
}

func (p Polygon) Copy() Polygon {
	newPolygon := make(Polygon, len(p))
	copy(newPolygon, p)
	return newPolygon
}

func (pf *Polyfence) Inside(P image.Point) bool {
	// First check if point is on an edge
	if pf.OnEdge(P) {
		return true
	}
	return !(checkOutside(pf.p, P) == 0)
}

func (pf *Polyfence) OnEdge(P image.Point) bool {
	n := len(pf.p) - 1
	for i := 0; i < n; i++ {
		p1, p2 := pf.p[i], pf.p[i+1]

		// Check if P is on the line segment from p1 to p2
		// Point is on segment if it's collinear and between p1 and p2

		// Check if collinear using cross product
		cross := (p2.X-p1.X)*(P.Y-p1.Y) - (P.X-p1.X)*(p2.Y-p1.Y)
		if cross != 0 {
			continue
		}

		// Check if P is between p1 and p2
		minX, maxX := p1.X, p2.X
		if minX > maxX {
			minX, maxX = maxX, minX
		}
		minY, maxY := p1.Y, p2.Y
		if minY > maxY {
			minY, maxY = maxY, minY
		}

		if P.X >= minX && P.X <= maxX && P.Y >= minY && P.Y <= maxY {
			return true
		}
	}
	return false
}

func checkOutside(p Polygon, P image.Point) (res int) {
	if len(p) < 3 {
		return 0
	}
	n := len(p) - 1 // Since we append first point to close, check all edges
	for i := 0; i < n; i++ {
		if p[i].Y <= P.Y {
			if p[i+1].Y > P.Y {
				if isLeft(p[i], p[i+1], P) > 0 {
					res++
				}
			}
		} else {
			if p[i+1].Y <= P.Y {
				if isLeft(p[i], p[i+1], P) < 0 {
					res--
				}
			}
		}
	}
	return
}

func isLeft(A, B, P image.Point) int {
	return (B.X-A.X)*(P.Y-A.Y) - (P.X-A.X)*(B.Y-A.Y)
}
