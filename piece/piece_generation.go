package piece

import (
	"math/rand/v2"
)

const bagSize = 7

type SevenBag struct {
	pieces [bagSize]Piece
	len    int
}

func (bag *SevenBag) getPiece() Piece {
	if bag.len == 0 {
		*bag = newSevenBag()
	}

	bag.len--
	return bag.pieces[bag.len]
}

func newSevenBag() SevenBag {
	bag := [bagSize]Piece{
		newJPiece(),
		newLPiece(),
		newOPiece(),
		newSPiece(),
		newTPiece(),
		newZPiece(),
		newIPiece(),
	}

	rand.Shuffle(bagSize, func(i, j int) {
		bag[i], bag[j] = bag[j], bag[i]
	})

	return SevenBag{
		pieces: bag,
		len:    bagSize,
	}
}
