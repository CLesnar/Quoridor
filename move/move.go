package move

import (
	"encoding/json"
	"fmt"
	"quoridor/pawn"
	"quoridor/wall"
)

// Format
/*
Walls: 0 - 7c; 1 - 8r;
"p1,p,8,4;p2,p,0,4;p1,p,7,4;p2,p,1,4;p1,w,r4,4,5;p2,w,r4,2,3;p1,p,6,4;p2,p,2,4;p1,w,c2,1,2;p2,w,c4,5,6;"
p#,p,row,column; "p1,p,8,4;"
p#,w,<r or c>#,#,#; "p1,w,r4,4,5;" "p1,w,c2,1,2;"
*/

type Move struct {
	Player int        `json:"player"`
	Pawn   *pawn.Pawn `json:"pawn,omitempty"`
	Wall   *wall.Wall `json:"wall,omitempty"`
}

func (m Move) String() string {
	if m.Wall != nil && m.Pawn != nil || m.Wall == nil && m.Pawn == nil {
		return ""
	}
	if bytes, err := json.Marshal(m); err != nil {
		return ""
	} else {
		return fmt.Sprint(bytes)
	}
}
