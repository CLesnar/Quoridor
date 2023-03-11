package player

import (
	"errors"
	"fmt"
	"quoridor/move"
	"quoridor/pawn"
	"quoridor/point"
	"quoridor/wall"
)

type PlayerDef struct {
	pawn.PawnDef `json:"-,"`
	Name         string      `json:"name"`
	Start        point.Point `json:"start"`
	Winning      point.Point `json:"winning"`
	Num          int         `json:"number"`
	Walls        int         `json:"walls"`
	Symbol       string      `json:"symbol"`
}

type Player struct {
	PlayerDef
	pawn.Pawn
	Walls []wall.Wall
}

func (p *PlayerDef) CreatePlayer() (*Player, error) {
	if p == nil {
		return nil, errors.New("playerdef cannot be nil")
	}
	pawn, err := p.CreatePawn(p.Num)
	if err != nil {
		return nil, fmt.Errorf("failed to create player: %v", err)
	}
	walls := []wall.Wall{}
	for i := 0; i < p.Walls; i++ {
		walls = append(walls, wall.Wall{})
	}
	player := &Player{
		Pawn:  *pawn,
		Walls: walls,
	}
	return player, nil
}

func (p Player) GetPawnMove() *move.Move {
	return &move.Move{
		Pawn:   &p.Pawn,
		Player: p.Num,
	}
}

func (p Player) GetWallMove(w *wall.Wall) *move.Move {
	return &move.Move{
		Wall:   w,
		Player: p.Num,
	}
}

func (p *Player) MovePawn(to point.Point) error {
	if p == nil {
		return errors.New("player cannot be nil")
	}
	p.Move(to)
	return nil
}

func (p *Player) GetNextAvailableWall() *wall.Wall {
	zero := point.Point{
		X: 0,
		Y: 0,
	}
	for _, w := range p.Walls {
		if point.Equal(zero, w.P1, w.P2) {
			return &w
		}
	}
	return nil
}

func (p *Player) MoveWall(p1, p2 point.Point) error {
	if p == nil {
		return errors.New("player cannot be nil")
	}
	wall := p.GetNextAvailableWall()
	if wall == nil {
		return errors.New("no more walls")
	}
	wall.P1, wall.P2 = p1, p2
	return nil
}

func (p Player) HasWon() bool {
	return (p.Winning.X > 0 && p.X == p.Winning.X) || (p.Winning.Y > 0 && p.Y == p.Winning.Y)
}

func (p Player) PlayerStr() string {
	return fmt.Sprintf("%d:%s", p.Num, p.Name)
}
