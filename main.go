package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

const (
	screenWidth         = 600
	screenHeight        = 600
	gridWidth           = 10
	gridHeight          = 23
	squareSize          = 26
	nextPieceSquareSize = 18
	tickDuration        = 10
	ticksBeforeLock     = 10
	defaultDAS          = 8
	defaultARR          = 3
)

type Square struct {
	Color color.Color
}

type Grid [gridWidth][gridHeight]*Square

type Game struct {
	Grid               Grid
	currentPiece       *Piece
	pieceQueue         []*Piece
	nextTick           int
	ticksSinceLastDrop int
	bag                *SevenBag
	holdPiece          *Piece
	hasHeld            bool
	ghostPiece         *Piece
	gameOver           bool
	das                int
	arr                int
}

func newGame() *Game {
	game := &Game{}

	game.bag = newSevenBag()

	game.currentPiece = game.bag.getPiece()
	game.pieceQueue = make([]*Piece, 5)
	for i := 0; i < 5; i++ {
		game.pieceQueue[i] = game.bag.getPiece()
	}
	game.currentPiece.MoveToTop()
	game.das = defaultDAS
	game.arr = defaultARR
	return game
}

func (g *Game) Layout(int, int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	game := newGame()
	ebiten.SetWindowSize(screenWidth, screenHeight)
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
