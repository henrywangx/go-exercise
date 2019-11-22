package main

import (
	"fmt"
	"math/rand"
	"time"
)

type SecondGame struct {
	Board
	// count for used chess of a col
	used [colSum]int
	// bit to represent whether current col is available
	availableCol map[int]bool
	winner       string
}

func (b *SecondGame) Placement() {
	rand.Seed(time.Now().Unix())
	fmt.Println("Placement of SecondGame")
	b.InitAvailableCol()
	b.Init("Black")
	fmt.Println("Placment before run:")
	b.Print()
	b.Run()
	//b.ShowHistory()
	fmt.Println("Placment after run:")
	b.Print()
	fmt.Printf("The winner is: %s\n", b.winner)
	fmt.Println("--------------Second End---------------")
}

func (b *SecondGame) InitAvailableCol() {
	b.availableCol = make(map[int]bool)
	for col := 0; col < colSum; col++ {
		b.availableCol[col] = true
	}
}

func (b *SecondGame) Run() {
	for {
		b.NewTry()
		if b.IsFinished() {
			break
		}
		b.ChangePlayer()
	}
	b.winner = b.CurrentPlayerName()
}

func (b *SecondGame) IsFinished() bool {
	if len(b.availableCol) == 0 {
		return true
	}
	return false
}

func (b *SecondGame) NewTry() {
	randomCol := rand.Intn(colSum)
	_, ok := b.availableCol[randomCol]
	if !ok {
		// NewChess will only record it
		b.NewChess(b.CurrentPlayerName(), 0, randomCol)
		return
	}
	row := rowSum - b.used[randomCol] - 1
	b.NewChess(b.CurrentPlayerName(), row, randomCol)
	b.used[randomCol]++
	if row == 0 {
		delete(b.availableCol, randomCol)
	}
}
