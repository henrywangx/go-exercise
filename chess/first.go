package main

import (
	"fmt"
)

type FirstGame struct {
	Board
}

func (b *FirstGame) Placement() {
	fmt.Println("Placement of FirstGame")
	b.Init("Black")
	b.NewChess("Black", 0, 3)
	b.NewChess("Black", 3, 2)
	b.NewChess("White", 2, 3)
	b.NewChess("White", 3, 5)
	//b.ShowHistory()
	b.Print()
	fmt.Println("--------------First End---------------")
}
