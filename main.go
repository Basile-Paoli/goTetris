package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image/color"
)

const (
	screenWidth  = 600
	screenHeight = 600
	gridWidth    = 10
	gridHeight   = 20
	squareSize   = 26
	tickDuration = 10
)

type Square struct {
	Color color.Color
}

type Game struct {
	Grid         [gridWidth][gridHeight + 2]*Square
	currentPiece Piece
	nextTick     int
	bag          SevenBag
}

func newGame() *Game {
	bag := newSevenBag()
	piece := bag.getPiece()
	return &Game{
		bag:          *bag,
		currentPiece: piece,
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) dropPiece() bool {
	for _, block := range g.currentPiece.BlockCoordinates() {
		if block[1] == 0 {
			return false
		}
		if g.Grid[block[0]][block[1]-1] != nil {
			return false
		}
	}
	g.currentPiece.Drop()
	return true
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		for _, block := range g.currentPiece.BlockCoordinates() {
			if block[0] == 0 || g.Grid[block[0]-1][block[1]] != nil {
				return nil
			}
		}
		g.currentPiece.MoveLeft()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		for _, block := range g.currentPiece.BlockCoordinates() {
			if block[0] == gridWidth-1 || g.Grid[block[0]+1][block[1]] != nil {
				return nil
			}
		}
		g.currentPiece.MoveRight()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyH) {
		g.currentPiece.RotateCounterClockwise()
		for _, block := range g.currentPiece.BlockCoordinates() {
			if block[0] < 0 || block[0] >= gridWidth || block[1] < 0 || block[1] >= gridHeight {
				g.currentPiece.RotateClockwise()
				break
			}
			if g.Grid[block[0]][block[1]] != nil {
				g.currentPiece.RotateClockwise()
				break
			}

		}

	}
	if inpututil.IsKeyJustPressed(ebiten.KeyJ) {
		g.currentPiece.RotateClockwise()
		for _, block := range g.currentPiece.BlockCoordinates() {
			if block[0] < 0 || block[0] >= gridWidth || block[1] < 0 || block[1] >= gridHeight {
				g.currentPiece.RotateCounterClockwise()
				break
			}
			if g.Grid[block[0]][block[1]] != nil {
				g.currentPiece.RotateCounterClockwise()
				break
			}
		}
	}
	if g.nextTick == 0 {
		if !g.dropPiece() {
			for _, block := range g.currentPiece.BlockCoordinates() {
				g.Grid[block[0]][block[1]] = &Square{Color: g.currentPiece.Color()}
			}
			g.currentPiece = g.bag.getPiece()
		}
		g.nextTick = tickDuration
		g.lineDestruction()
	}
	g.nextTick--
	return nil
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

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	for x := 0; x < gridWidth; x++ {
		g.drawSquare(Square{Color: color.White}, x, -1, screen)
	}
	for y := -1; y < gridHeight+1; y++ {
		g.drawSquare(Square{Color: color.White}, -1, y, screen)
		g.drawSquare(Square{Color: color.White}, gridWidth, y, screen)
	}
	for x := 0; x < gridWidth; x++ {
		for y := 0; y < gridHeight; y++ {
			if g.Grid[x][y] != nil {
				g.drawSquare(*g.Grid[x][y], x, y, screen)
			}
		}
	}
	if g.currentPiece != nil {
		for _, block := range g.currentPiece.BlockCoordinates() {
			g.drawSquare(Square{Color: g.currentPiece.Color()}, block[0], block[1], screen)
		}

	}
}

func (g *Game) drawSquare(square Square, column, row int, screen *ebiten.Image) {
	x := (screenWidth-gridWidth*squareSize)/2 + column*squareSize
	y := screenHeight - (row+2)*squareSize
	for i := 1; i < squareSize-1; i++ {
		for j := 1; j < squareSize-1; j++ {
			screen.Set(x+i, y+j, square.Color)
		}

	}

}
func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	game := newGame()
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}

}
