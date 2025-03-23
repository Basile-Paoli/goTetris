package piece

import (
	"image/color"
	"tetris/grid"
)

type vector struct {
	x, y int
}
type Position struct {
	X, Y int
}

type Piece struct {
	Color                       color.Color
	center                      Position
	rotation                    int
	clockWiseRotationMap        [4][]vector
	counterClockWiseRotationMap [4][]vector
	blockPositionsFromRotation  [4][]vector
}

var NullPiece = Piece{}

func (p *Piece) IsNull() bool {
	return p.Color == nil
}

func (p *Piece) MoveDown(g grid.Grid) bool {
	for _, block := range p.CurrentBlockPositions() {
		if !g.IsEmptySquare(block.X, block.Y-1) {
			return false
		}
	}
	p.center.Y--
	return true
}

func (p *Piece) MoveLeft(grid grid.Grid) {
	for _, block := range p.CurrentBlockPositions() {
		if !grid.IsEmptySquare(block.X-1, block.Y) {
			return
		}
	}
	p.center.X--
}

func (p *Piece) MoveRight(grid grid.Grid) {
	for _, block := range p.CurrentBlockPositions() {
		if !grid.IsEmptySquare(block.X+1, block.Y) {
			return
		}
	}
	p.center.X++
}

func (p *Piece) MoveToTop() {
	p.center = Position{grid.Width/2 - 1, grid.Height - 3}
	p.rotation = 0
}

func (p *Piece) isValidMovement(grid grid.Grid, translationVector vector) bool {
	for _, block := range p.CurrentBlockPositions() {
		x, y := block.X+translationVector.x, block.Y+translationVector.y
		if !grid.IsEmptySquare(x, y) {
			return false
		}
	}
	return true
}

func (p *Piece) CurrentBlockPositions() (blocks []Position) {
	for _, vec := range p.blockPositionsFromRotation[p.rotation] {
		blocks = append(blocks, Position{p.center.X + vec.x, p.center.Y + vec.y})
	}
	return blocks
}

func (p *Piece) DefaultBlockPositions() (positions []Position) {
	for _, vec := range p.blockPositionsFromRotation[0] {
		positions = append(positions, Position{vec.x, vec.y})
	}
	return positions
}

var tPieceBlockPositionMap = [4][]vector{
	0: {{-1, 0}, {0, 0}, {1, 0}, {0, 1}},
	1: {{0, -1}, {0, 0}, {0, 1}, {+1, 0}},
	2: {{-1, 0}, {0, 0}, {1, 0}, {0, -1}},
	3: {{0, -1}, {0, 0}, {0, 1}, {-1, 0}},
}

func newTPiece() Piece {
	return Piece{
		Color:                       color.RGBA{R: 153, G: 0, B: 255, A: 255},
		center:                      Position{0, 0},
		rotation:                    0,
		clockWiseRotationMap:        jltszClockwiseRotationMap,
		counterClockWiseRotationMap: jltszCounterClockwiseRotationMap,
		blockPositionsFromRotation:  tPieceBlockPositionMap,
	}
}

var IPieceBlockPositions = [4][]vector{
	0: {{-1, 1}, {0, 1}, {1, 1}, {2, 1}},
	1: {{1, -1}, {1, 0}, {1, 1}, {1, 2}},
	2: {{-1, 0}, {0, 0}, {1, 0}, {2, 0}},
	3: {{0, -1}, {0, 0}, {0, 1}, {0, 2}},
}

func newIPiece() Piece {
	return Piece{
		rotation:                    0,
		center:                      Position{},
		Color:                       color.RGBA{R: 0, G: 255, B: 255, A: 255},
		blockPositionsFromRotation:  IPieceBlockPositions,
		clockWiseRotationMap:        iClockwiseRotationMap,
		counterClockWiseRotationMap: iCounterClockwiseRotationMap,
	}
}

func getOPieceBlockPositions() (res [4][]vector) {
	for i := range 4 {
		res[i] = []vector{
			{0, 0},
			{1, 0},
			{0, 1},
			{1, 1},
		}
	}
	return res
}

var oPieceBlockPositions = getOPieceBlockPositions()

func newOPiece() Piece {
	return Piece{
		Color:                       color.RGBA{R: 255, G: 255, B: 0, A: 255},
		center:                      Position{},
		rotation:                    0,
		blockPositionsFromRotation:  oPieceBlockPositions,
		clockWiseRotationMap:        oPieceClockwiseRotationMap,
		counterClockWiseRotationMap: oPieceCounterClockwiseRotationMap,
	}
}

var sPieceBlockPositions = [4][]vector{
	0: {{-1, 0}, {0, 0}, {0, 1}, {1, 1}},
	1: {{0, 1}, {0, 0}, {1, 0}, {1, -1}},
	2: {{-1, -1}, {0, -1}, {0, 0}, {1, 0}},
	3: {{-1, 1}, {-1, 0}, {0, 0}, {0, -1}},
}

func newSPiece() Piece {
	return Piece{
		Color:                       color.RGBA{R: 0, G: 255, B: 0, A: 255},
		rotation:                    0,
		center:                      Position{},
		clockWiseRotationMap:        jltszClockwiseRotationMap,
		counterClockWiseRotationMap: jltszCounterClockwiseRotationMap,
		blockPositionsFromRotation:  sPieceBlockPositions,
	}
}

var zPieceBlockPositions = [4][]vector{
	0: {{-1, 1}, {0, +1}, {0, 0}, {1, 0}},
	1: {{1, 1}, {1, 0}, {0, 0}, {0, -1}},
	2: {{-1, 0}, {0, 0}, {0, -1}, {1, -1}},
	3: {{-1, -1}, {-1, 0}, {0, 0}, {0, 1}},
}

func newZPiece() Piece {
	return Piece{
		Color:                       color.RGBA{R: 255, G: 0, B: 0, A: 255},
		center:                      Position{},
		rotation:                    0,
		blockPositionsFromRotation:  zPieceBlockPositions,
		clockWiseRotationMap:        jltszClockwiseRotationMap,
		counterClockWiseRotationMap: jltszCounterClockwiseRotationMap,
	}
}

var jPieceBlockPositions = [4][]vector{
	0: {{-1, 1}, {-1, 0}, {0, 0}, {1, 0}},
	1: {{1, 1}, {0, 1}, {0, 0}, {0, -1}},
	2: {{-1, 0}, {0, 0}, {1, 0}, {1, -1}},
	3: {{-1, -1}, {0, -1}, {0, 0}, {0, 1}},
}

func newJPiece() Piece {
	return Piece{
		Color:                       color.RGBA{R: 0, G: 0, B: 255, A: 255},
		rotation:                    0,
		center:                      Position{},
		blockPositionsFromRotation:  jPieceBlockPositions,
		clockWiseRotationMap:        jltszClockwiseRotationMap,
		counterClockWiseRotationMap: jltszCounterClockwiseRotationMap,
	}
}

var lPieceBlockPositions = [4][]vector{
	0: {{-1, 0}, {0, 0}, {1, 0}, {1, 1}},
	1: {{0, 1}, {0, 0}, {0, -1}, {1, -1}},
	2: {{-1, 0}, {0, 0}, {1, 0}, {-1, -1}},
	3: {{0, 1}, {0, 0}, {0, -1}, {-1, 1}},
}

func newLPiece() Piece {
	return Piece{
		Color:                       color.RGBA{R: 255, G: 170, B: 0, A: 255},
		rotation:                    0,
		center:                      Position{},
		blockPositionsFromRotation:  lPieceBlockPositions,
		clockWiseRotationMap:        jltszClockwiseRotationMap,
		counterClockWiseRotationMap: jltszCounterClockwiseRotationMap,
	}
}
