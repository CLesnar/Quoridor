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
	"quoridor/square"
	"quoridor/wall"
)

type GameDef struct {
	Pawn           pawn.PawnDef
	Board          board.BoardDef
	HorizontalWall wall.WallDef
	VerticalWall   wall.WallDef
}

type Game struct {
	Title   string           `json:"title"`
	Def     *GameDef         `json:"definition"`
	Players []*player.Player `json:"players"`
	Board   board.Board      `json:"board"`
	History move.History     `json:"history"`
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

func (g Game) MovePlayer(p *player.Player, moveInput move.Move) error {
	if p == nil {
		return errors.New("player cannot be nil")
	}
	switch {
	case moveInput.Pawn != nil:
		if err := p.MovePawn(moveInput.Pawn.Point); err != nil {
			return fmt.Errorf("failed to move player %s: %v", p.String(), err)
		}
		g.Board.MovePawn(p.Prev, p.Point)
		return nil
	case moveInput.Wall != nil:
		if err := p.MoveWall(moveInput.Wall.P1, moveInput.Wall.P2); err != nil {
			return fmt.Errorf("failed to place player %s wall: %v", p.String(), err)
		}
		g.Board.PlaceWall(*moveInput.Wall)
		return nil
	default:
		return fmt.Errorf("failed to move player %s: invalid move: %v", p.PlayerStr(), moveInput)
	}
}

func (g *Game) GetAllWalls() []wall.Wall {
	walls := []wall.Wall{}
	for _, p := range g.Players {
		walls = append(walls, p.Walls...)
	}
	return walls
}

func (g Game) GetPlayer(player int) *player.Player {
	for _, p := range g.Players {
		if p.Num == player {
			return p
		}
	}
	return nil
}

func (g Game) PointIsOccupied(p point.Point) error {
	for _, p := range g.Players {
		if p.IsEqual(p.Point) {
			return fmt.Errorf("point is occupied by player %s", p.PlayerStr())
		}
	}
	return nil
}

func (g *Game) IsValidPawnMove(player player.Player, pawnMove pawn.Pawn) error {
	if err := g.Board.IsValidPoint(pawnMove.Point); err != nil {
		return fmt.Errorf("pawn has invalid point: %v", err)
	}
	if err := g.PointIsOccupied(pawnMove.Point); err != nil {
		return err
	}
	subtractPoint := player.Point.Subtract(pawnMove.Point).Abs()
	if !subtractPoint.Equals(point.Point{X: 1, Y: 0}, point.Point{X: 0, Y: 1}) {
		return errors.New("pawn cannot move more than 1 space")
	}
	var mid1, mid2 point.Point
	switch {
	case player.IsMoveUp(pawnMove.Point):
		mid1, mid2 = point.Point{X: player.X, Y: player.Y}, point.Point{X: player.X + 1, Y: player.Y}
	case player.IsMoveDown(pawnMove.Point):
		mid1, mid2 = point.Point{X: player.X, Y: player.Y - 1}, point.Point{X: player.X + 1, Y: player.Y - 1}
	case player.IsMoveRight(pawnMove.Point):
		mid1, mid2 = point.Point{X: player.X, Y: player.Y}, point.Point{X: player.X, Y: player.Y + 1}
	case player.IsMoveLeft(pawnMove.Point):
		mid1, mid2 = point.Point{X: player.X - 1, Y: player.Y}, point.Point{X: player.X - 1, Y: player.Y + 1}
	}
	walls := g.GetAllWalls()
	for _, w := range walls {
		if err := w.IsValid(); err != nil {
			continue
		}
		midW := w.Midpoint()
		if midW.Equals(mid1, mid2) {
			return fmt.Errorf("wall %v is blocking", w)
		}
	}
	return nil
}

func (g *Game) IsValidMove(moveInput move.Move) error {
	err := fmt.Errorf("invalid move: %v", moveInput)
	if moveInput.Pawn != nil {
		if err := g.Board.IsValidPoint(moveInput.Pawn.Point); err != nil {
			return fmt.Errorf("invalid move %v: %v", moveInput, err)
		}
		if err := g.PointIsOccupied(moveInput.Pawn.Point); err != nil {
			return err
		}
		g.IsValidPawnMove(*g.GetPlayer(moveInput.Player), *moveInput.Pawn)
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

func (g Game) GetValidPawnMovesInDirection(player int, p point.Point, direction board.Direction) []move.Move {
	moves := []move.Move{}
	var squareToCheck, toSquare *square.Square
	var wallIsBlocking bool
	switch direction {
	case board.North:
		squareToCheck = g.Board.GetSquare(p)
		if squareToCheck == nil {
			return moves
		}
		wallIsBlocking = squareToCheck.HasWall == square.HasWallHorizontal
		toSquare = g.Board.GetSquareNorth(p)
		if toSquare == nil {
			return moves
		}
	case board.South:
		squareToCheck = g.Board.GetSquareSouth(p)
		if squareToCheck == nil {
			return moves
		}
		wallIsBlocking = squareToCheck.HasWall == square.HasWallHorizontal
		toSquare = squareToCheck
	case board.East:
		squareToCheck = g.Board.GetSquare(p)
		if squareToCheck == nil {
			return moves
		}
		wallIsBlocking = squareToCheck.HasWall == square.HasWallVertical
		toSquare = g.Board.GetSquareEast(p)
		if toSquare == nil {
			return moves
		}
	case board.West:
		squareToCheck = g.Board.GetSquareWest(p)
		if squareToCheck == nil {
			return moves
		}
		wallIsBlocking = squareToCheck.HasWall == square.HasWallVertical
		toSquare = squareToCheck
	default:
		return moves
	}

	if wallIsBlocking {
		return moves
	}

	if toSquare.IsOccupied {
		// Special cases for jumping
		threeSquares := g.Board.GetThreeSquaresDirection(*toSquare, direction)
		sq := g.Board.GetSquareDirection(*toSquare, direction)
		if sq != nil {
			if DirectionPerpendiculartoWall(direction, *sq) {
				delete(threeSquares, sq.String())
				var dir board.Direction
				for _, v := range threeSquares {
					switch direction {
					case board.North, board.South:
						if v.X < sq.X {
							dir = board.East
						} else if v.X > sq.X {
							dir = board.West
						}
					case board.East, board.West:
						if v.Y < sq.Y {
							dir = board.South
						} else if v.Y > sq.Y {
							dir = board.North
						}
					default:
						continue
					}
					moves = append(moves, g.GetValidPawnMovesInDirection(player, v.Point, dir)...)
				}
			} else {
				moves = append(moves, g.GetValidPawnMovesInDirection(player, sq.Point, direction)...)
			}
		}
	} else {
		moves = append(moves, move.Move{Player: player, Pawn: &pawn.Pawn{Point: point.Point{X: toSquare.X, Y: toSquare.Y}}})
	}

	return moves
}

func DirectionPerpendiculartoWall(dir board.Direction, sq square.Square) bool {
	if (sq.HasWall == square.HasWallHorizontal && (dir == board.North || dir == board.South)) ||
		(sq.HasWall == square.HasWallVertical && (dir == board.East || dir == board.West)) {
		return true
	}
	return false
}
