package move

import (
	"encoding/json"
	"fmt"
	"quoridor/board"
	"quoridor/pawn"
	"quoridor/point"
	"quoridor/wall"
)

// Format
/*
Walls: 0 - 7c; 1 - 8r;
"p1,p,8,4;p2,p,0,4;p1,p,7,4;p2,p,1,4;p1,w,r4,4,5;p2,w,r4,2,3;p1,p,6,4;p2,p,2,4;p1,w,c2,1,2;p2,w,c4,5,6;"
p#,p,row,column; "p1,p,8,4;"
p#,w,<r or c>#,#,#; "p1,w,r4,4,5;" "p1,w,c2,1,2;"
*/

type MovePlayer struct {
	PlayerNum     int    `json:"player"`
	Name          string `json:"name"`
	WinningRow    *int   `json:"winningRow,omitempty"`
	WinningColumn *int   `json:"winningColumn,omitempty"`
}

type History struct {
	Moves   []Move       `json:"moves"`
	Board   board.Board  `json:"board"`
	Players []MovePlayer `json:"players"`
}

type Move struct {
	Player int        `json:"player"`
	Pawn   *pawn.Pawn `json:"pawn,omitempty"`
	Wall   *wall.Wall `json:"wall,omitempty"`
}

func (m Move) String() string {
	if bytes, err := json.Marshal(m); err == nil {
		return fmt.Sprint(bytes)
	}
	return ""
}

func (h History) String() string {
	if bytes, err := json.Marshal(h); err == nil {
		return fmt.Sprint(bytes)
	}
	return ""
}

func GetWallMap(boardPosition []Move) map[string]wall.Wall {
	wallMap := map[string]wall.Wall{}
	for _, position := range boardPosition {
		wallMove := position.Wall
		if wallMove != nil {
			wallMap[wallMove.String()] = *wallMove
		}
	}
	return wallMap
}

func FindAllWallMoves(boardPosition History, player int) []wall.Wall {
	wallMoves := []wall.Wall{}
	wallRows, wallColumns := boardPosition.Board.Rows-1, boardPosition.Board.Columns-1
	pointsMap := point.CreatePointMap(point.Point{X: 0, Y: 0}, point.Point{X: wallRows, Y: wallColumns})

	wallMap := GetWallMap(boardPosition.Moves)
	for _, w := range wallMap {
		delete(pointsMap, w.P1.String())
		delete(pointsMap, w.P2.String())
	}

	for _, point := range pointsMap {
		fmt.Println(point) // todo:
	}

	return wallMoves
}

func FindAllPawnMoves(boardPosition History, player int) []pawn.Pawn {
	pawnMoves := []pawn.Pawn{}

	pointsMap := point.CreatePointMap(point.Point{X: 0, Y: 0}, point.Point{X: wallRows, Y: wallColumns}) // TODO
	return pawnMoves
}
