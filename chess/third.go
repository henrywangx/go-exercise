package main

import (
	"fmt"
	"math/rand"
	"time"
)

type ThirdGame struct {
	Board
	// count for used chess of a col
	used         [colSum]int
	// bit to represent whether current col is available
	availableCol map[int]bool
	winner string
}

func (b *ThirdGame) Placement() {
	rand.Seed(time.Now().Unix())
	fmt.Println("Placement of ThirdGame")
	b.InitAvailableCol()
	b.Init("Black")
	fmt.Println("Placment before run:")
	b.Print()
	b.Run()
	//b.ShowHistory()
	fmt.Println("Placment after run:")
	b.Print()
	fmt.Printf("The winner is: %s\n", b.winner)
	fmt.Println("--------------Third End---------------")
}

func (b *ThirdGame) InitAvailableCol() {
	b.availableCol = make(map[int]bool)
	for col:=0; col<colSum; col++ {
		b.availableCol[col] = true
	}
}

func (b *ThirdGame) Run() {
	for {
		isWin := b.NewTry()
		if isWin {
			break
		}
		if b.IsFinished() {
			b.winner = "Nobody"
			break
		}
		b.ChangePlayer()
	}
}

func (b *ThirdGame) IsFinished() bool {
	if len(b.availableCol) == 0 {
		return true
	}
	return false
}

func (b *ThirdGame) JudgeWin(row, col int) bool {
	value := b.players[b.CurrentPlayerName()].value
	// 1.vertical
	var up, bottom int
	if row>=3 {
		up = row-3
		bottom = 5
	} 
	if 3>row {
		up = 0
		bottom = row+3
	}
	var cnt1 = 1
	for i:=row-1; i>=up; i-- {
		if b.placement[i][col] == value {
			cnt1++
		} else {
			break
		}
		if cnt1 == 4 {
			return true
		}
	}
	for j:=row+1; j<=bottom; j++ {
		if b.placement[j][col] == value {
			cnt1++
		} else {
			break
		}
		if cnt1 == 4 {
			return true
		}
	}
	// 2.horizontal
	var left, right int
	if col>=3 {
		left = col-3
		right = 6
	} 
	if 3>col {
		left = 0
		right = col + 3
	}
	var cnt2 = 1
	for i:=col-1; i>=left; i-- {
		if b.placement[row][i] == value {
			cnt2++
		} else {
			break
		}
		if cnt2 == 4 {
			return true
		}
	}
	for j:=col+1; j<=right; j++ {
		if b.placement[row][j] == value {
			cnt2++
		} else {
			break
		}
		if cnt2 == 4 {
			return true
		}	
	}
	// 3.diagonal
	var di = row-1
	var dj = col-1
	var cnt3 = 1
	for {
		if di < up || dj < left {
			break
		}
		if b.placement[di][dj] == value {
			cnt3++
		} else {
			break
		}
		if cnt3 == 4 {
			return true
		}
		di--
		dj--
	}
	di = row+1
	dj = col+1
	for {
		if di > bottom || dj > right {
			break
		}
		if b.placement[di][dj] == value {
			cnt3++
		} else {
			break
		}
		if cnt3 == 4 {
			return true
		}
		di++
		dj++
	}
	
	// 4.back-diagonal
	var bi = row-1
	var bj = col+1
	var cnt4 = 1
	for {
		if bi < up || bj > right {
			break
		}
		if b.placement[bi][bj] == value {
			cnt4++
		} else {
			break
		}
		if cnt4 == 4 {
			return true
		}
		bi--
		bj++
	}
	bi = row+1
	bj = col-1
	for {
		if bi > bottom || bj < left {
			break
		}
		if b.placement[bi][bj] == value {
			cnt4++
		} else {
			break
		}
		if cnt4 == 4 {
			return true
		}
		bi++
		bj--
	}
	return false
}

func (b *ThirdGame) NewTry() bool{
	randomCol := rand.Intn(colSum)
	_, ok := b.availableCol[randomCol]
	if !ok {
		// NewChess will only record it
		b.NewChess(b.CurrentPlayerName(), 0, randomCol)
		return false
	}
	row := rowSum - b.used[randomCol] - 1
	b.NewChess(b.CurrentPlayerName(), row, randomCol)
	if b.JudgeWin(row, randomCol) {
		b.winner = b.CurrentPlayerName()
		return true
	}
	b.used[randomCol]++
	if row == 0 {
		delete(b.availableCol, randomCol)
	}
	return false
}
