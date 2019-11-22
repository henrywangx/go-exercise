package main

import (
	"container/list"
	"fmt"
)

const (
	rowSum = 6
	colSum = 7
)

type Board struct {
	players           map[string]*Player
	currentPlayerName string
	rowSum            int
	colSum            int
	placement         [rowSum][colSum]int
	//record history for debugging
	history *list.List
}

type action struct {
	playerName string
	row        int
	col        int
	success    bool
}

type Player struct {
	// black or white
	name string
	// black: 1, white:2
	value int
	// black: "x", white: "o"
	present string
}

func main() {
	fmt.Println("Let's Play Chess ^ ^!")
	fmt.Println("---------------------------------------")
	b1 := new(FirstGame)
	b1.Placement()
	b2 := new(SecondGame)
	b2.Placement()
	b3 := new(ThirdGame)
	b3.Placement()
}

// Init() will init players and action history
func (b *Board) Init(firstPlayerName string) {
	// 1.init players
	b.players = make(map[string]*Player)
	black := &Player{
		name:    "Black",
		value:   1,
		present: "x",
	}
	white := &Player{
		name:    "White",
		value:   2,
		present: "o",
	}
	b.players["Black"] = black
	b.players["White"] = white
	b.currentPlayerName = firstPlayerName
	if firstPlayerName != "Black" && firstPlayerName != "White" {
		b.currentPlayerName = "Black"
	}
	// 2.init history
	b.history = list.New()
}

func (b *Board) Print() {
	showMap := map[int]string{
		0: ".",
		1: "x",
		2: "o",
	}
	// row:6, col:7
	fmt.Println("Row | Placement")
	for row := 0; row < rowSum; row++ {
		fmt.Printf("%d: ", row)
		for col := 0; col < colSum; col++ {
			fmt.Printf(" %s", showMap[b.placement[row][col]])
		}
		fmt.Println("")
	}
}

func (b *Board) ChangePlayer() {
	if b.currentPlayerName == "Black" {
		b.currentPlayerName = "White"
		return
	}
	b.currentPlayerName = "Black"
}

func (b *Board) CurrentPlayerName() string {
	return b.currentPlayerName
}

func (b *Board) NewChess(playerName string, row int, col int) {
	var success bool
	player := b.players[playerName]

	if b.placement[row][col] == 0 {
		success = true
		b.placement[row][col] = player.value
	}
	// Save this log in history
	newAction := &action{
		playerName: player.name,
		row:        row,
		col:        col,
		success:    success,
	}
	b.history.PushBack(newAction)
}

func (b *Board) ShowHistory() {
	fmt.Println("Action History:")
	round := 1
	for element := b.history.Front(); element != nil; element = element.Next() {
		action := element.Value.(*action)
		fmt.Printf("round:%d, player:%s, row:%d, col:%d, success:%t\n", round, action.playerName, action.row, action.col, action.success)
		round++
	}
	fmt.Println("Show history end")
	fmt.Println("---------------------------------------")
}
