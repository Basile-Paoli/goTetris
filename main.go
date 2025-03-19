package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"tetris/grid"
	"tetris/piece"
)

type Game struct {
	grid               grid.Grid
	currentPiece       piece.Piece
	pieceQueue         piece.Queue
	nextTick           int
	ticksSinceLastDrop int
	holdPiece          piece.Piece
	hasHeld            bool
	gameOver           bool
	das                int
	arr                int
	softDropDelay      int
	config             Config
}

func newGame() *Game {
	pieceQueue := piece.NewQueue()
	config := defaultConfig()
	game := &Game{
		grid:               grid.Grid{},
		currentPiece:       pieceQueue.Next(),
		pieceQueue:         pieceQueue,
		nextTick:           config.tickDuration,
		ticksSinceLastDrop: 0,
		holdPiece:          piece.NullPiece,
		hasHeld:            false,
		gameOver:           false,
		das:                config.das,
		arr:                config.arr,
		softDropDelay:      0,
		config:             config,
	}
	game.currentPiece.MoveToTop()
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
