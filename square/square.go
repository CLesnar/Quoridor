package square

import (
	"encoding/json"
	"quoridor/point"
)

type SquareHasWall int

const (
	HasWallNone SquareHasWall = iota
	HasWallVertical
	HasWallHorizontal
)

type Square struct {
	point.Point `json:"point"`
	HasWall     SquareHasWall `json:"haswall"`
	IsOccupied  bool          `json:"isoccupied"`
}

func (s Square) String() string {
	if sBytes, err := json.Marshal(s); err == nil {
		return string(sBytes)
	}
	return ""
}

func (s *Square) UpdateWall(hasWall SquareHasWall) {
	s.HasWall = hasWall
}
func (s *Square) UpdateOccupied(isOccupied bool) {
	s.IsOccupied = isOccupied
}
