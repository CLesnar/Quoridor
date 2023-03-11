package point

import (
	"fmt"
	"math"
)

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (p Point) String() string {
	return fmt.Sprintf("%d,%d", p.X, p.Y)
}

func (p Point) Distance(q Point) int {
	xi, yi := p.X-q.X, p.Y-q.Y
	x, y := float64(xi), float64(yi)
	dist := math.Sqrt(x*x + y*y)
	return int(dist)
}

func (p Point) IsEqual(q Point) bool {
	return p.X == q.X && p.Y == q.Y
}

func (p Point) Equals(points ...Point) bool {
	for _, point := range points {
		if !p.IsEqual(point) {
			return false
		}
	}
	return true
}

func Equal(points ...Point) bool {
	if len(points) < 2 {
		return true
	}
	q := points[0]
	for _, p := range points[1:] {
		if !q.IsEqual(p) {
			return false
		}
	}
	return false
}

func (p Point) Add(q Point) Point {
	return Point{
		X: p.X + q.X,
		Y: p.Y + q.Y,
	}
}

func (p Point) Subtract(q Point) Point {
	return Point{
		X: p.X - q.X,
		Y: p.Y - q.Y,
	}
}

func (p Point) Abs() Point {
	return Point{
		X: int(math.Abs(float64(p.X))),
		Y: int(math.Abs(float64(p.Y))),
	}
}

func (p Point) Divide(divisor int) Point {
	return Point{
		X: p.X / divisor,
		Y: p.Y / divisor,
	}
}

func (p Point) Copy() Point {
	return Point{
		X: p.X,
		Y: p.Y,
	}
}
