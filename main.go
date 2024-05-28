package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image/color"
)

const (
	screenWidth         = 600
	screenHeight        = 600
	gridWidth           = 10
	gridHeight          = 20
	squareSize          = 26
	nextPieceSquareSize = 18
	tickDuration        = 10
	ticksBeforeLock     = 10
)

type Square struct {
	Color color.Color
}

type Game struct {
	Grid               [gridWidth][gridHeight + 2]*Square
	currentPiece       Piece
	pieceQueue         []Piece
	nextTick           int
	ticksSinceLastDrop int
	bag                SevenBag
	holdPiece          Piece
	ghostPiece         Piece
	gameOver           bool
}

func newGame() *Game {
	game := &Game{}
	game.bag = *newSevenBag()
	game.currentPiece = game.bag.getPiece()
	game.pieceQueue = make([]Piece, 5)
	for i := 0; i < 5; i++ {
		game.pieceQueue[i] = game.bag.getPiece()
	}
	game.currentPiece.MoveToTop()
	return game
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
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

func (g *Game) Update() error {
	if g.gameOver {
		return nil
	}
	g.processInputs()
	if g.nextTick == 0 {
		if !g.dropPiece() {
			g.ticksSinceLastDrop++
			if g.ticksSinceLastDrop >= ticksBeforeLock {
				for _, block := range g.currentPiece.BlockCoordinates() {
					g.Grid[block[0]][block[1]] = &Square{Color: g.currentPiece.Color()}
				}
				g.nextPiece()
				g.ticksSinceLastDrop = 0
			}
		}
		g.nextTick = tickDuration
		g.lineDestruction()
	}
	g.setGhostPiece()
	g.nextTick--
	return nil
}

func (g *Game) processInputs() {
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
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		for g.dropPiece() {

		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		if g.holdPiece == nil {
			g.currentPiece.MoveToTop()
			g.holdPiece = g.currentPiece
			g.nextPiece()
		} else {
			g.currentPiece.MoveToTop()
			g.currentPiece, g.holdPiece = g.holdPiece, g.currentPiece
			g.currentPiece.MoveToTop()
		}
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
	if g.ghostPiece != nil {
		g.drawGhostPiece(screen)
	}
	if g.currentPiece != nil {
		for _, block := range g.currentPiece.BlockCoordinates() {
			g.drawSquare(Square{Color: g.currentPiece.Color()}, block[0], block[1], screen)
		}

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
func (g *Game) drawGhostPiece(screen *ebiten.Image) {
	R, G, B, _ := g.currentPiece.Color().RGBA()
	ghostColor := color.RGBA{R: uint8(R) / 3, G: uint8(G) / 3, B: uint8(B) / 3, A: 100}
	for _, block := range g.ghostPiece.BlockCoordinates() {
		g.drawSquare(Square{Color: ghostColor}, block[0], block[1], screen)
	}

}
func (g *Game) drawHold(screen *ebiten.Image) {
	if g.holdPiece == nil {
		return
	}
	baseX := 70 - gridWidth/2*nextPieceSquareSize
	baseY := 70 + gridHeight*nextPieceSquareSize
	for _, block := range g.holdPiece.BlockCoordinates() {
		for x := 1; x < nextPieceSquareSize-1; x++ {
			for y := 1; y < nextPieceSquareSize-1; y++ {
				screen.Set(baseX+x+block[0]*nextPieceSquareSize, baseY+y-block[1]*nextPieceSquareSize, g.holdPiece.Color())
			}
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
