package main

import (
	"math/rand/v2"
)

type SevenBag []*Piece

func (bag *SevenBag) getPiece() *Piece {
	if len(*bag) == 0 {
		*bag = *newSevenBag()
	}
	n := rand.Int() % len(*bag)
	piece := (*bag)[n]
	*bag = append((*bag)[:n], (*bag)[n+1:]...)
	return piece
}
func newSevenBag() *SevenBag {
	return &SevenBag{
		NewJPiece(),
		NewLPiece(),
		NewOPiece(),
		NewSPiece(),
		NewTPiece(),
		NewZPiece(),
		NewIPiece(),
	}
}
