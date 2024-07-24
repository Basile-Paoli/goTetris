package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	g.drawLayout(screen)

	g.drawGrid(screen)

	if g.ghostPiece != nil {
		g.drawGhostPiece(screen)
	}

	if g.currentPiece != nil {
		g.drawCurrentPiece(screen)

	}
	g.drawNext(screen)
	g.drawHold(screen)
}

func (g *Game) drawSquare(square Square, column, row int, screen *ebiten.Image) {
	x := 170 + column*squareSize
	y := screenHeight - (row+2)*squareSize
	for i := 1; i < squareSize-1; i++ {
		for j := 1; j < squareSize-1; j++ {
			screen.Set(x+i, y+j, square.Color)
		}

	}
}

func (g *Game) drawLayout(screen *ebiten.Image) {
	for x := 0; x < gridWidth; x++ {
		g.drawSquare(Square{Color: color.White}, x, -1, screen)
	}
	for y := -1; y < gridHeight+1; y++ {
		g.drawSquare(Square{Color: color.White}, -1, y, screen)
		g.drawSquare(Square{Color: color.White}, gridWidth, y, screen)
	}
}

func (g *Game) drawGrid(screen *ebiten.Image) {
	for x := 0; x < gridWidth; x++ {
		for y := 0; y < gridHeight; y++ {
			if g.Grid[x][y] != nil {
				g.drawSquare(*g.Grid[x][y], x, y, screen)
			}
		}
	}
}

func (g *Game) drawGhostPiece(screen *ebiten.Image) {
	R, G, B, _ := g.currentPiece.Color().RGBA()
	ghostColor := color.RGBA{R: uint8(R) / 3, G: uint8(G) / 3, B: uint8(B) / 3, A: 100}
	for _, block := range g.ghostPiece.BlockCoordinates() {
		g.drawSquare(Square{Color: ghostColor}, block[0], block[1], screen)
	}
}
func (g *Game) drawCurrentPiece(screen *ebiten.Image) {
	for _, block := range g.currentPiece.BlockCoordinates() {
		g.drawSquare(Square{Color: g.currentPiece.Color()}, block[0], block[1], screen)
	}
}
func (g *Game) drawNext(screen *ebiten.Image) {
	for i, piece := range g.pieceQueue {
		for _, block := range piece.BlockCoordinates() {
			for x := 1; x < nextPieceSquareSize-1; x++ {
				for y := 1; y < nextPieceSquareSize-1; y++ {
					screen.Set(520+x+block[0]*nextPieceSquareSize, 100+i*nextPieceSquareSize*4-block[1]*nextPieceSquareSize+y, piece.Color())
				}

			}
		}
	}

}
func (g *Game) drawHold(screen *ebiten.Image) {
	if g.holdPiece == nil {
		return
	}
	baseX := 70 - gridWidth/2*squareSize
	baseY := 70 + gridHeight*squareSize
	for _, block := range g.holdPiece.BlockCoordinates() {
		for x := 1; x < squareSize-1; x++ {
			for y := 1; y < squareSize-1; y++ {
				screen.Set(baseX+x+block[0]*squareSize, baseY+y-block[1]*squareSize, g.holdPiece.Color())
			}
		}
	}
}
