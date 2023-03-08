package quoridor

// Old Project:
// C:\DEV\old-pc-devenv\DEVENV\DEV\src\quoridor\api\include\quoridor_api.h

/*

                      |  |  |  |  |  |  |  |  |  |
                      |  |  |  |  |  |  |  |  |  |
                      |  |  |  |  |  |  |  |  |  |

             +----+----+----+----+----+----+----+----+----+
             | 8       | w,c1,8,7 -P2-                    |
             +    +    |    +    +    +    +    +    +    +
             | 7       |                                  |
  ---------  +    +    +    +    +    +    +    +    +    +  ---------
  ---------  | 6                                          |  ---------
  ---------  +    +    +    +    +    +    +    +    +    +  ---------
  ---------  | 5                                          |  ---------
  ---------  +    +    +    +    +    +    +    +    +    +  ---------
  ---------  | 4                                          |  ---------
  ---------  +    +    +    +    +    +    +    +    +    +  ---------
  ---------  | 3                                          |  ---------
  ---------  +    +    +    +    +    +    +    +    +    +  ---------
  ---------  | 2                                          |  ---------
             +    +    +    +    +    +    +    +    +    +
             | 1                                          |
             +    +    +    +    +    +    +    +    +    +
             | 0                  -P1-                    |
             +----+----+----+----+----+----+----+----+----+

                      |  |  |  |  |  |  |  |  |  |
                      |  |  |  |  |  |  |  |  |  |
                      |  |  |  |  |  |  |  |  |  |

Walls: 0 - 7c; 1 - 8r;
"p1,p,8,4;p2,p,0,4;p1,p,7,4;p2,p,1,4;p1,w,r4,4,5;p2,w,r4,2,3;p1,p,6,4;p2,p,2,4;p1,w,c2,1,2;p2,w,c4,5,6;"
p#,p,row,column; "p1,p,8,4;"
p#,w,<r or c>#,#,#; "p1,w,r4,4,5;" "p1,w,c2,1,2;"

*/
/*
* Quoridor Board
* (row, Column) == (x, y)
*
*
*     { 0  } P4 Winning Column
*     y
*     |                     P3 Winning Column { 8  }
*     +----+----+----+----+----+----+----+----+----+
*  8  |                    -P2-                  8 |   { 8 } P1 Winning Row
*     +    +    +    +    +    +    +    +    +    +
*  7  |                                     7      |
*     +    +    +    +    +    +    +    +    +    +
*  6  |                                6           |
*     +    +    +    +    +    +    +    +    +    +
*  5  |                          5                 |
*     +    +    +    +    +    +    +    +    +    +
*  4  |-P3-                  4                 -P4-|
*     +    +    +    +    +    +    +    +    +    +
*  3  |                3                           |
*     +    +    +    +    +    +    +    +    +    +
*  2  |            2                               |
*     +    +    +    +    +    +    +    +    +    +
*  1  |      1                                     |
*     +    +    +    +    +    +    +    +    +    +
*  0  | 0                  -P1-                    |    { 0 } P2 Winning Row
*     +----+----+----+----+----+----+----+----+----+---x
*       0    1    2    3    4    5    6    7    8
*
*  Player | Start  | Win/End
*-------------------------------
*    P1   | (4, 0) |  (_, 8)
*    P2   | (4, 8) |  (_, 0)
*    P3   | (0, 4) |  (8, _)
*    P4   | (8, 4) |  (0, _)
 */
