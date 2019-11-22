# Play Chess
## How to run it
```bash
cd chess
go build .
./chess
```
The example output:

```bash
➜  chess ./chess
Let's Play Chess ^ ^!
---------------------------------------
Placement of FirstGame
Row | Placement
0:  . . . x . . .
1:  . . . . . . .
2:  . . . o . . .
3:  . . x . . o .
4:  . . . . . . .
5:  . . . . . . .
--------------First End---------------
Placement of SecondGame
Placment before run:
Row | Placement
0:  . . . . . . .
1:  . . . . . . .
2:  . . . . . . .
3:  . . . . . . .
4:  . . . . . . .
5:  . . . . . . .
Placment after run:
Row | Placement
0:  x x x o o o o
1:  x x o o o x o
2:  o x x o x o x
3:  x o o o o x o
4:  o x o o o x o
5:  o x x x x o x
The winner is: White
--------------Second End---------------
Placement of ThirdGame
Placment before run:
Row | Placement
0:  . . . . . . .
1:  . . . . . . .
2:  . . . . . . .
3:  . . . . . . .
4:  . . . . . . .
5:  . . . . . . .
Placment after run:
Row | Placement
0:  . . . . . . .
1:  . . . . . . .
2:  . . . . . o x
3:  . . . . . x o
4:  o x o o . x o
5:  o x x x x o x
The winner is: Black
--------------Third End---------------
```