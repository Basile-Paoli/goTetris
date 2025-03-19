package grid

import "image/color"

type Square color.Color

const (
	Width  = 10
	Height = 23
)

var EmptySquare Square = nil

type Grid [Width][Height]Square

func (g *Grid) isValidSquare(x, y int) bool {
	return 0 <= x && x < Width && 0 <= y && y < Height
}

func (g *Grid) IsEmptySquare(x, y int) bool {
	return g.isValidSquare(x, y) && g[x][y] == EmptySquare
}

func (g *Grid) IsOccupiedSquare(x, y int) bool {
	return g.isValidSquare(x, y) && g[x][y] != EmptySquare
}

func (g *Grid) IsLineFull(y int) bool {
	for x := range Width {
		if g[x][y] == EmptySquare {
			return false
		}
	}
	return true
}

func (g *Grid) DestroyLine(line int) {
	for x := range Width {
		g[x][line] = EmptySquare
	}

	for y := line; y < Height-1; y++ {
		for x := range Width {
			g[x][y] = g[x][y+1]
		}
	}
}
