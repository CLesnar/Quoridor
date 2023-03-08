package game

import (
	"errors"
	"fmt"
	"log"
	"os"
	"quoridor/board"
	"quoridor/move"
	"quoridor/pawn"
	"quoridor/player"
	"quoridor/point"
	"quoridor/wall"
)

type GameDef struct {
	Pawn           pawn.PawnDef
	Board          board.BoardDef
	HorizontalWall wall.WallDef
	VerticalWall   wall.WallDef
}

type Game struct {
	Title   string
	Def     *GameDef
	Players []*player.Player
	Board   board.Board
}

func (g *GameDef) CreateGame(title string, boardRowsColumns point.Point, players ...player.PlayerDef) (*Game, error) {

	var allPlayers []*player.Player

	for _, player := range players {
		p, err := player.CreatePlayer()
		if err != nil {
			return nil, fmt.Errorf("failed to create game: %v", err)
		}
		allPlayers = append(allPlayers, p)
	}

	game := &Game{
		Title: title,
		Board: board.Board{
			Rows:    boardRowsColumns.X,
			Columns: boardRowsColumns.Y,
		},
		Players: allPlayers,
	}

	return game, nil
}

func (g *Game) Start() {
	for _, p := range g.Players {
		for {
			log.Printf("waiting for player %s\n", p.PlayerStr())
			moveInput := g.GetPlayerMoveInput()
			moveInput.Player = p.Num
			if err := g.IsValidMove(moveInput); err != nil {
				log.Printf("player %v has invalid move", p)
				continue
			}
			g.MovePlayer(p, moveInput)
			if p.HasWon() {
				g.EndGame(p)
			}
		}
	}
}

func (g *Game) GetAllWalls() []wall.Wall {
	walls := []wall.Wall{}
	for _, p := range g.Players {
		walls = append(walls, p.Walls...)
	}
	return walls
}

func (g Game) PointIsOccupied(p point.Point) error {
	for _, p := range g.Players {
		if p.IsEqual(p.Point) {
			return fmt.Errorf("point is occupied by player %s", p.PlayerStr())
		}
	}
	return nil
}

func (g Game) MovePlayer(p *player.Player, moveInput move.Move) error {
	if p == nil {
		return errors.New("player cannot be nil")
	}
	if moveInput.Pawn != nil {
		return p.MovePawn(moveInput.Pawn.Point)
	} else if moveInput.Wall != nil {
		return p.MoveWall(moveInput.Wall.P1, moveInput.Wall.P2)
	} else {
		return fmt.Errorf("failed to move player %s: invalid move: %v", p.PlayerStr(), moveInput)
	}
}

func (g *Game) IsValidMove(moveInput move.Move) error {
	err := fmt.Errorf("invalid move: %v", moveInput)
	if moveInput.Pawn != nil {
		if err := g.Board.IsValidPoint(moveInput.Pawn.Point); err != nil {
			return fmt.Errorf("pawn has invalid point: %v", err)
		}
		if err := g.PointIsOccupied(moveInput.Pawn.Point); err != nil {
			return err
		}
		// TODO: p.IsValidMove() // pawn can only move 1 square at a time. Pawn cannot move through walls.
		return nil
	}
	if moveInput.Wall != nil {
		if err := wall.PointPairIsValid(moveInput.Wall.P1, moveInput.Wall.P2); err != nil {
			return fmt.Errorf("point pair is not valid for a wall: %v", err)
		}
		if err := g.Board.IsValidPoint(moveInput.Wall.P1); err != nil {
			return fmt.Errorf("wall has invalid point: %v", err)
		} else if err := g.Board.IsValidPoint(moveInput.Wall.P2); err != nil {
			return fmt.Errorf("wall has invalid point: %v", err)
		} else if err := moveInput.Wall.IsValid(); err != nil {
			return fmt.Errorf("wall is invalid: %v", err)
		}
		walls := g.GetAllWalls()
		for _, w := range walls {
			if err := moveInput.Wall.Overlaps(w); err != nil {
				return err
			}
		}
		// TODO: g.IsValidWallMove() // verify all pawns can reach their goal/winning row/column - no boxing in
		return nil
	}
	return err
}

func (g *Game) GetPlayerMoveInput() move.Move {
	m := move.Move{}
	return m
}

func (g *Game) EndGame(p *player.Player) {
	log.Printf("*** Player %s Wins! ***", p.PlayerStr())
	os.Exit(0) // or restart
}

func (g *Game) IsPointOccupied(p point.Point) (bool, error) {
	if g == nil {
		return false, errors.New("game cannot be nil")
	}
	for _, player := range g.Players {
		if player.IsEqual(p) {
			return true, nil
		}
	}
	return false, nil
}
