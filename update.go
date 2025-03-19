package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"tetris/grid"
)

func (g *Game) Update() error {
	if g.gameOver && !inpututil.IsKeyJustPressed(ebiten.KeySemicolon) {
		return nil
	}

	g.processInputs()

	if g.nextTick > 0 {
		g.nextTick--
		return nil
	}

	if !g.dropCurrentPiece() {
		g.ticksSinceLastDrop++

		if g.ticksSinceLastDrop >= g.config.ticksBeforeLock {
			g.handleLock()
		}
	}

	g.nextTick = g.config.tickDuration - 1
	return nil
}

func (g *Game) handleLock() {
	g.freezePiece()
	g.lineDestruction()
	g.nextPiece()
	g.ticksSinceLastDrop = 0
	g.hasHeld = false
}

func (g *Game) processInputs() {
	if inpututil.IsKeyJustPressed(ebiten.KeySemicolon) {
		*g = *newGame()
		return
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.switchPiece()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		g.currentPiece.MoveLeft(g.grid)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		g.currentPiece.MoveRight(g.grid)
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyD) {
		g.performDas()
	} else {
		g.das = g.config.das
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyH) {
		g.currentPiece.RotateCounterClockwise(g.grid)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyJ) {
		g.currentPiece.RotateClockwise(g.grid)
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.softDrop()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		g.hardDrop()
	}
}

func (g *Game) switchPiece() {

	if g.holdPiece.IsNull() {
		g.holdPiece = g.currentPiece
		g.nextPiece()
	} else if !g.hasHeld {
		g.currentPiece, g.holdPiece = g.holdPiece, g.currentPiece
		g.currentPiece.MoveToTop()
	}
	g.hasHeld = true
}

func (g *Game) shouldMoveFromDas() bool {
	g.das--
	if g.das > 0 {
		return false
	}
	g.arr--
	if g.arr > 0 {
		return false
	}

	g.arr = g.config.arr
	return true
}

func (g *Game) performDas() {
	if !g.shouldMoveFromDas() {
		return
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.currentPiece.MoveLeft(g.grid)
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.currentPiece.MoveRight(g.grid)
	}
}

func (g *Game) softDrop() {

	if g.config.softDropRate == 0 {
		for g.dropCurrentPiece() {
		}
		return
	}

	if g.softDropDelay <= 0 {
		g.dropCurrentPiece()
		g.softDropDelay = g.config.softDropRate
	}
	g.softDropDelay--
}

func (g *Game) hardDrop() {
	for g.dropCurrentPiece() {
	}
	g.handleLock()
}

func (g *Game) dropCurrentPiece() (dropped bool) {
	return g.currentPiece.MoveDown(g.grid)
}

func (g *Game) freezePiece() {
	for _, block := range g.currentPiece.BlockCoordinates() {
		g.grid[block.X][block.Y] = g.currentPiece.Color
	}
}

func (g *Game) nextPiece() {
	g.currentPiece = g.pieceQueue.Next()
	g.currentPiece.MoveToTop()
	for _, block := range g.currentPiece.BlockCoordinates() {
		if g.grid.IsOccupiedSquare(block.X, block.Y) {
			g.gameOver = true
		}
	}
}

func (g *Game) lineDestruction() {
	for y := range grid.Height {
		for g.grid.IsLineFull(y) {
			g.grid.DestroyLine(y)
		}
	}
}
