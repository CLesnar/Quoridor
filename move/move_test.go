package move

import (
	"encoding/json"
	"fmt"
	"quoridor/pawn"
	"quoridor/point"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MoveString(t *testing.T) {
	m := Move{
		Player: 1,
		Pawn: &pawn.Pawn{
			Point: point.Point{
				X: 1,
				Y: 2,
			},
		},
	}
	b, err := json.Marshal(m)
	assert.Nil(t, err)
	fmt.Println(string(b))
}
