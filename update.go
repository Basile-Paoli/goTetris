package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) Update() error {
	if g.gameOver && !inpututil.IsKeyJustPressed(ebiten.KeySemicolon) {
		return nil
	}

	g.processInputs()

	if g.nextTick == 0 {
		if !g.dropPiece() {
			g.ticksSinceLastDrop++

			if g.ticksSinceLastDrop >= ticksBeforeLock {
				g.freezePiece()
				g.lineDestruction()
				g.nextPiece()
				g.ticksSinceLastDrop = 0
				g.hasHeld = false
			}
		}
		g.nextTick = tickDuration
	}
	g.setGhostPiece()
	g.nextTick--
	return nil
}

func (g *Game) processInputs() {
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.das--
		if g.das <= 0 {
			g.arr--
			if g.arr <= 0 {
				legal := true
				for _, block := range g.currentPiece.BlockCoordinates() {
					if block[0] == 0 || g.Grid[block[0]-1][block[1]] != nil {
						legal = false
					}
				}
				if legal {
					g.currentPiece.MoveLeft()
				}
				g.arr = defaultARR
			}
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.das--
		if g.das <= 0 {
			g.arr--
			if g.arr <= 0 {
				legal := true
				for _, block := range g.currentPiece.BlockCoordinates() {
					if block[0] == gridWidth-1 || g.Grid[block[0]+1][block[1]] != nil {
						legal = false
					}
				}
				if legal {
					g.currentPiece.MoveRight()
				}
				g.arr = defaultARR
			}
		}
	}
	if !ebiten.IsKeyPressed(ebiten.KeyA) && !ebiten.IsKeyPressed(ebiten.KeyD) {
		g.das = defaultDAS
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySemicolon) {
		*g = *newGame()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		for g.dropPiece() {
		}
		g.nextTick = 0
		g.ticksSinceLastDrop = ticksBeforeLock
		return
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		for g.dropPiece() {
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		if g.holdPiece == nil {
			g.currentPiece.MoveToTop()
			g.holdPiece = g.currentPiece
			g.nextPiece()
		} else if !g.hasHeld {
			g.currentPiece.MoveToTop()
			g.currentPiece, g.holdPiece = g.holdPiece, g.currentPiece
		}
		g.hasHeld = true
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		legal := true
		for _, block := range g.currentPiece.BlockCoordinates() {
			if block[0] == 0 || g.Grid[block[0]-1][block[1]] != nil {
				legal = false
			}
		}
		if legal {
			g.currentPiece.MoveLeft()
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		legal := true
		for _, block := range g.currentPiece.BlockCoordinates() {
			if block[0] == gridWidth-1 || g.Grid[block[0]+1][block[1]] != nil {
				legal = false
			}
		}
		if legal {
			g.currentPiece.MoveRight()
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyH) {
		g.currentPiece.RotateCounterClockwise(g.Grid)

	}
	if inpututil.IsKeyJustPressed(ebiten.KeyJ) {
		g.currentPiece.RotateClockwise(g.Grid)
	}
}

func (g *Game) dropPiece() bool {
	for _, block := range g.currentPiece.BlockCoordinates() {
		if block[1] == 0 || g.Grid[block[0]][block[1]-1] != nil {
			return false
		}
	}
	g.currentPiece.Drop()
	return true
}

func (g *Game) freezePiece() {
	for _, block := range g.currentPiece.BlockCoordinates() {
		g.Grid[block[0]][block[1]] = &Square{Color: g.currentPiece.Color()}
	}

}

func (g *Game) nextPiece() {
	g.currentPiece = g.pieceQueue[0]
	g.currentPiece.MoveToTop()
	g.pieceQueue = append(g.pieceQueue[1:], g.bag.getPiece())
	for _, block := range g.currentPiece.BlockCoordinates() {
		if g.Grid[block[0]][block[1]] != nil {
			g.gameOver = true
		}
	}
}

func (g *Game) setGhostPiece() {
	g.ghostPiece = g.currentPiece.Copy()
	for {
		for _, block := range g.ghostPiece.BlockCoordinates() {
			if block[1] == 0 || g.Grid[block[0]][block[1]-1] != nil {
				return
			}
		}
		g.ghostPiece.Drop()
	}
}

func (g *Game) lineDestruction() {
	for y := 0; y < gridHeight; y++ {
		full := true
		for x := 0; x < gridWidth; x++ {
			if g.Grid[x][y] == nil {
				full = false
				break
			}
		}
		if full {
			for x := 0; x < gridWidth; x++ {
				g.Grid[x][y] = nil
			}
			for y2 := y; y2 < gridHeight-1; y2++ {
				for x := 0; x < gridWidth; x++ {
					g.Grid[x][y2] = g.Grid[x][y2+1]
				}
			}
			y--
		}

	}
}
