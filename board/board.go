package board

import (
	"fmt"
	"quoridor/point"
	"quoridor/square"
	"quoridor/wall"
)

type Direction int

const (
	North Direction = iota
	South
	East
	West
)

type BoardDef struct {
	RowSeparator    string
	ColumnSeparator string
}

type Board struct {
	Rows    int               `json:"rows"`
	Columns int               `json:"columns"`
	Squares [][]square.Square `json:"squares"`
}

func CreateBoard(rows int, columns int) Board {
	b := Board{
		Rows:    rows,
		Columns: columns,
	}
	for r := 0; r < b.Rows; r++ {
		for c := 0; c < b.Columns; c++ {
			b.Squares[r] = append(b.Squares[r], square.Square{Point: point.Point{X: r, Y: c}})
		}
	}
	return b
}

func (b Board) SetSquare(p point.Point, hasWall square.SquareHasWall, isOccupied bool) {
	square := b.Squares[p.X][p.Y]
	square.HasWall = hasWall
	square.IsOccupied = isOccupied
}

func (b Board) GetSquare(p point.Point) *square.Square {
	if b.IsValidPoint(p) == nil {
		return &b.Squares[p.X][p.Y]
	}
	return nil
}

func (b Board) GetSquareDirection(sq square.Square, dir Direction) *square.Square {
	switch dir {
	case North:
		return b.GetSquareNorth(sq.Point)
	case South:
		return b.GetSquareSouth(sq.Point)
	case East:
		return b.GetSquareEast(sq.Point)
	case West:
		return b.GetSquareWest(sq.Point)
	default:
		return nil
	}
}

func (b Board) GetThreeSquaresDirection(sq square.Square, dir Direction) map[string]square.Square {
	squares := map[string]square.Square{}
	s := sq
	addToMap := func(sq square.Square, dir Direction) {
		if a := b.GetSquareDirection(s, dir); a != nil {
			squares[a.String()] = *a
		}
	}
	switch dir {
	case North, South:
		addToMap(s, dir)
		s.Y--
		addToMap(s, dir)
		s.Y += 2
		addToMap(s, dir)
	case East, West:
		addToMap(s, dir)
		s.X--
		addToMap(s, dir)
		s.X += 2
		addToMap(s, dir)
	}
	return squares
}

func (b Board) GetSquareNorth(p point.Point) *square.Square {
	p.X++
	return b.GetSquare(p)
}

func (b Board) GetSquareSouth(p point.Point) *square.Square {
	p.X--
	return b.GetSquare(p)
}

func (b Board) GetSquareEast(p point.Point) *square.Square {
	p.Y++
	return b.GetSquare(p)
}

func (b Board) GetSquareWest(p point.Point) *square.Square {
	p.Y--
	return b.GetSquare(p)
}

func (b Board) IsValidPoint(p point.Point) error {
	if p.X < b.Rows || p.Y < b.Columns || p.X > 0 || p.Y > 0 {
		return nil
	}
	return fmt.Errorf("point %v is outside board boundary", p)
}

func (b Board) IsValidWallPoint(points ...point.Point) error {
	wallRow, wallColumn := b.Rows-1, b.Columns-1
	for _, p := range points {
		if p.X > wallRow || p.Y > wallColumn {
			return fmt.Errorf("point %v is an invalid wall point", p)
		}
	}
	return nil
}

func (b Board) MovePawn(from point.Point, to point.Point) {
	b.GetSquare(from).UpdateOccupied(false)
	b.GetSquare(to).UpdateOccupied(true)
}

func (b Board) PlaceWall(w wall.Wall) {
	wallUpdate := square.HasWallVertical
	if w.IsHorizontal() {
		wallUpdate = square.HasWallHorizontal
	}
	b.GetSquare(w.P1).UpdateWall(wallUpdate)
	b.GetSquare(w.P2).UpdateWall(wallUpdate)
}
