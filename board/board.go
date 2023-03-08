package board

import (
	"errors"
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

func (b *Board) IsValidPoint(p point.Point) error {
	if b == nil {
		return errors.New("board cannot be nil")
	}
	if p.X < b.Rows || p.Y < b.Columns || p.X > 0 || p.Y > 0 {
		return nil
	}
	return fmt.Errorf("point %v is outside board boundary", p)
}
