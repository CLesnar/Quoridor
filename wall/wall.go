package wall

import (
	"encoding/json"
	"errors"
	"fmt"
	"quoridor/point"
)

type WallDef struct {
	Separator string
	Symbol    string
}

type Wall struct {
	P1 point.Point `json:"point1"`
	P2 point.Point `json:"point2"`
}

func (w Wall) String() string {
	if wBytes, err := json.Marshal(w); err == nil {
		return string(wBytes)
	}
	return ""
}

func (w *Wall) Move(p, q point.Point) error {
	if w == nil {
		return errors.New("wall cannot be nil")
	}
	w.P1 = p
	w.P2 = q
	return nil
}

func PointPairIsValid(p, q point.Point) error {
	distance := p.Distance(q)
	if distance == 1 && (p.X == q.X || p.Y == q.Y) {
		return nil
	}
	return errors.New("wall points must be horizontal or vertical and must be 1 unit away")
}

func (w Wall) IsValid() error {
	return PointPairIsValid(w.P1, w.P2)
}

func (w Wall) IsHorizontal() bool {
	return w.P1.X == w.P2.X
}

func (w Wall) IsVertical() bool {
	return w.P1.Y == w.P2.Y
}

func (w Wall) Midpoint() point.Point {
	return w.P1.Add(w.P2).Divide(2)
}

func (w Wall) Overlaps(x Wall) error {
	midX, midW := x.Midpoint(), w.Midpoint()
	if midW.IsEqual(midX) {
		return fmt.Errorf("wall overlaps %v", x)
	}
	if w.IsHorizontal() && x.IsHorizontal() {
		midXXPlus1, midXXMinus1 := midX.Copy(), midX.Copy()
		midXXPlus1.X += 1
		midXXMinus1.X -= 1
		if !midW.IsEqual(midXXMinus1) && !midW.IsEqual(midXXPlus1) {
			return nil
		}
	}
	if w.IsVertical() && x.IsVertical() {
		midXYPlus1, midXYMinus1 := midX.Copy(), midX.Copy()
		midXYPlus1.Y += 1
		midXYMinus1.Y -= 1
		if !midW.IsEqual(midXYMinus1) && !midW.IsEqual(midXYPlus1) {
			return nil
		}
	}
	return fmt.Errorf("wall %v is invalid", x)
}
