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
	for _, block := range ghostPiece.BlockCoordinates() {
		drawSquare(ghostColor, block.X, block.Y, screen)
	}
}

func (g *Game) drawCurrentPiece(screen *ebiten.Image) {
	for _, block := range g.currentPiece.BlockCoordinates() {
		drawSquare(g.currentPiece.Color, block.X, block.Y, screen)
	}
}

func (g *Game) drawNext(screen *ebiten.Image) {
	for i, p := range g.pieceQueue.Pieces() {
		for _, block := range p.DefaultBlockCoordinates() {
			baseX := screenWidth - NextPieceRightPadding + block.X*nextPieceSquareSize
			baseY := NextPieceTopPadding + i*nextPieceSize - block.Y*nextPieceSquareSize

			for x := 1; x < nextPieceSquareSize-1; x++ {
				for y := 1; y < nextPieceSquareSize-1; y++ {
					screen.Set(baseX+x, baseY+y, p.Color)
				}
			}
		}
	}
}

func (g *Game) drawHold(screen *ebiten.Image) {
	if g.holdPiece.IsNull() {
		return
	}
	baseX := gridLeftPadding - (4 * squareSize)
	baseY := 4 * squareSize

	for _, block := range g.holdPiece.DefaultBlockCoordinates() {
		for x := 1; x < squareSize-1; x++ {
			for y := 1; y < squareSize-1; y++ {
				screen.Set(baseX+x+block.X*squareSize, baseY+y-block.Y*squareSize, g.holdPiece.Color)
			}
		}
	}
}

func drawSquare(square grid.Square, column, row int, screen *ebiten.Image) {
	baseX := gridLeftPadding + column*squareSize
	baseY := screenHeight - (row+2)*squareSize
	for i := 1; i < squareSize-1; i++ {
		for j := 1; j < squareSize-1; j++ {
			screen.Set(baseX+i, baseY+j, square)
		}
	}
}
