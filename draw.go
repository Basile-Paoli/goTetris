package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"tetris/grid"
	"tetris/piece"
)

const (
	screenWidth           = 600
	screenHeight          = 600
	gridLeftPadding       = 170
	NextPieceRightPadding = 80
	NextPieceTopPadding   = 100
	squareSize            = 26
	nextPieceSquareSize   = 18
	nextPieceSize         = 4 * nextPieceSquareSize
)

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	g.drawLayout(screen)

	g.drawGrid(screen)

	g.drawGhostPiece(screen)

	g.drawCurrentPiece(screen)

	g.drawNext(screen)
	g.drawHold(screen)
}

func (g *Game) drawLayout(screen *ebiten.Image) {
	for x := range grid.Width {
		drawSquare(color.White, x, -1, screen)
	}
	for y := -1; y < grid.Height-2; y++ {
		drawSquare(color.White, -1, y, screen)
		drawSquare(color.White, grid.Width, y, screen)
	}
}

func (g *Game) drawGrid(screen *ebiten.Image) {
	for x := range grid.Width {
		for y := range grid.Height {
			if g.grid[x][y] != grid.EmptySquare {
				drawSquare(g.grid[x][y], x, y, screen)
			}
		}
	}
}

func (g *Game) ghostPiece() piece.Piece {
	ghostPiece := g.currentPiece
	for ghostPiece.MoveDown(g.grid) {
	}

	return ghostPiece
}

func (g *Game) drawGhostPiece(screen *ebiten.Image) {
	R, G, B, _ := g.currentPiece.Color.RGBA()
	ghostColor := color.RGBA{R: uint8(R) / 3, G: uint8(G) / 3, B: uint8(B) / 3, A: 255}
	ghostPiece := g.ghostPiece()
	for _, block := range ghostPiece.CurrentBlockPositions() {
		drawSquare(ghostColor, block.X, block.Y, screen)
	}
}

func (g *Game) drawCurrentPiece(screen *ebiten.Image) {
	for _, block := range g.currentPiece.CurrentBlockPositions() {
		drawSquare(g.currentPiece.Color, block.X, block.Y, screen)
	}
}

func (g *Game) drawNext(screen *ebiten.Image) {
	for i, p := range g.pieceQueue.Pieces() {
		for _, block := range p.DefaultBlockPositions() {
			baseX := screenWidth - NextPieceRightPadding + block.X*nextPieceSquareSize
			baseY := NextPieceTopPadding + i*nextPieceSize - block.Y*nextPieceSquareSize
			drawRectangle(p.Color, baseX, baseY, nextPieceSquareSize-2, nextPieceSquareSize-2, screen)
		}
	}
}

func (g *Game) drawHold(screen *ebiten.Image) {
	if g.holdPiece.IsNull() {
		return
	}
	baseX := gridLeftPadding - int(4.5*squareSize)
	baseY := 4 * squareSize

	for _, block := range g.holdPiece.DefaultBlockPositions() {
		x := baseX + block.X*squareSize + 1
		y := baseY - block.Y*squareSize + 1
		drawRectangle(g.holdPiece.Color, x, y, squareSize-2, squareSize-2, screen)
	}
}

func drawSquare(square grid.Square, column, row int, screen *ebiten.Image) {
	baseX := gridLeftPadding + column*squareSize + 1
	baseY := screenHeight - (row+2)*squareSize + 1
	drawRectangle(square, baseX, baseY, squareSize-2, squareSize-2, screen)
}

func drawRectangle(color color.Color, x, y, width, height int, screen *ebiten.Image) {
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			screen.Set(x+i, y+j, color)
		}
	}
}
