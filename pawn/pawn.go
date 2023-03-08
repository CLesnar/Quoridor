package pawn

import (
	"errors"
	"quoridor/point"
)

type PawnDef struct {
	SymbolFmt string
}

func (p *PawnDef) CreatePawn(playerNum int) (*Pawn, error) {
	if p == nil {
		return nil, errors.New("pawndef cannot be nil")
	}
	return &Pawn{}, nil
}

type Pawn struct {
	point.Point `json:"point"`
	Prev        point.Point `json:"-"`
}

func (p *Pawn) Move(q point.Point) {
	p.Prev.X, p.Prev.Y = p.X, p.Y
	p.X, p.Y = q.X, q.Y
}
