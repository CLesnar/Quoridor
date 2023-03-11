package board

import (
	"fmt"
	"quoridor/point"
)

type BoardDef struct {
	RowSeparator    string
	ColumnSeparator string
}

type Board struct {
	Rows    int
	Columns int
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
