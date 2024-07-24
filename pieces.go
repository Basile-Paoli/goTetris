package main

import "image/color"

type Piece struct {
	color                       color.Color
	center                      [2]int
	rotation                    int
	clockWiseRotationMap        [4][][2]int
	counterClockWiseRotationMap [4][][2]int
	blockPositionsFromRotation  [4][][2]int
}

func (p *Piece) Color() color.Color {
	return p.color
}
func (p *Piece) Drop() {
	p.center[1]--
}
func (p *Piece) MoveLeft() {
	p.center[0]--
}
func (p *Piece) MoveRight() {
	p.center[0]++
}
func (p *Piece) MoveToTop() {
	p.center = [2]int{gridWidth/2 - 1, gridHeight - 3}
	p.rotation = 0
}
func (p *Piece) Copy() *Piece {
	return &Piece{
		color:                       p.color,
		center:                      p.center,
		rotation:                    p.rotation,
		clockWiseRotationMap:        p.clockWiseRotationMap,
		counterClockWiseRotationMap: p.counterClockWiseRotationMap,
		blockPositionsFromRotation:  p.blockPositionsFromRotation,
	}
}
func (p *Piece) RotateClockwise(grid Grid) {
	p.rotation = (p.rotation + 1) % 4

rotationLoop:
	for _, vector := range p.clockWiseRotationMap[p.rotation] {

		for _, block := range p.BlockCoordinates() {
			x, y := block[0]+vector[0], block[1]+vector[1]

			if x < 0 || x >= gridWidth || y < 0 || y >= gridHeight || grid[x][y] != nil {
				continue rotationLoop
			}
		}
		p.center[0] += vector[0]
		p.center[1] += vector[1]
		return
	}
	p.rotation = (p.rotation + 3) % 4
}

func (p *Piece) RotateCounterClockwise(grid Grid) {
	p.rotation = (p.rotation + 3) % 4

rotationLoop:
	for _, translationVector := range p.counterClockWiseRotationMap[p.rotation] {
		for _, block := range p.BlockCoordinates() {
			x, y := block[0]+translationVector[0], block[1]+translationVector[1]
			if x < 0 || x >= gridWidth || y < 0 || y >= gridHeight || grid[x][y] != nil {
				continue rotationLoop
			}
		}
		p.center[0] += translationVector[0]
		p.center[1] += translationVector[1]
		return
	}
	p.rotation = (p.rotation + 1) % 4
}

func (p *Piece) BlockCoordinates() [4][2]int {
	var blocks [4][2]int
	for i, position := range p.blockPositionsFromRotation[p.rotation] {
		blocks[i][0] = position[0] + p.center[0]
		blocks[i][1] = position[1] + p.center[1]
	}
	return blocks
}

func JLTSZClockwiseRotationMap() [4][][2]int {
	return [4][][2]int{
		0: {{0, 0}, {-1, 0}, {-1, -1}, {0, 2}, {-1, 2}},
		1: {{0, 0}, {-1, 0}, {-1, 1}, {0, -2}, {-1, -2}},
		2: {{0, 0}, {1, 0}, {1, -1}, {0, 2}, {1, 2}},
		3: {{0, 0}, {1, 0}, {1, 1}, {0, -2}, {1, -2}},
	}
}

func JLTSZCounterClockwiseRotationMap() [4][][2]int {
	return [4][][2]int{
		0: {{0, 0}, {1, 0}, {1, -1}, {0, 2}, {1, 2}},
		1: {{0, 0}, {-1, 0}, {-1, 1}, {0, -2}, {-1, -2}},
		2: {{0, 0}, {-1, 0}, {-1, -1}, {0, 2}, {-1, 2}},
		3: {{0, 0}, {1, 0}, {1, 1}, {0, -2}, {1, -2}},
	}
}

//Pieces

func TPieceBlockPositionMap() [4][][2]int {
	return [4][][2]int{
		0: {{-1, 0}, {0, 0}, {1, 0}, {0, 1}},
		1: {{0, -1}, {0, 0}, {0, 1}, {+1, 0}},
		2: {{-1, 0}, {0, 0}, {1, 0}, {0, -1}},
		3: {{0, -1}, {0, 0}, {0, 1}, {-1, 0}},
	}
}
func NewTPiece() *Piece {
	return &Piece{
		color:                       color.RGBA{R: 153, G: 0, B: 255, A: 255},
		center:                      [2]int{},
		rotation:                    0,
		clockWiseRotationMap:        JLTSZClockwiseRotationMap(),
		counterClockWiseRotationMap: JLTSZCounterClockwiseRotationMap(),
		blockPositionsFromRotation:  TPieceBlockPositionMap(),
	}
}

func IClockwiseRotationMap() [4][][2]int {
	return [4][][2]int{
		0: {{0, 0}, {1, 0}, {-2, 0}, {1, -2}, {-2, 1}},
		1: {{0, 0}, {-2, 0}, {1, 0}, {-2, -1}, {1, 2}},
		2: {{0, 0}, {-1, 0}, {2, 0}, {-1, 2}, {2, -1}},
		3: {{0, 0}, {2, 0}, {-1, 0}, {2, 1}, {-1, -2}},
	}
}

func ICounterClockwiseRotationMap() [4][][2]int {
	return [4][][2]int{
		0: {{0, 0}, {2, 0}, {-1, 0}, {2, 1}, {-1, -2}},
		1: {{0, 0}, {1, 0}, {-2, 0}, {1, -2}, {-2, 1}},
		2: {{0, 0}, {-2, 0}, {1, 0}, {-2, -1}, {1, 2}},
		3: {{0, 0}, {-1, 0}, {2, 0}, {-1, 2}, {2, -1}},
	}
}

func IPieceBlockPositions() [4][][2]int {
	return [4][][2]int{
		0: {{-1, 1}, {0, 1}, {1, 1}, {2, 1}},
		1: {{1, -1}, {1, 0}, {1, 1}, {1, 2}},
		2: {{-1, 0}, {0, 0}, {1, 0}, {2, 0}},
		3: {{0, -1}, {0, 0}, {0, 1}, {0, 2}},
	}
}

func NewIPiece() *Piece {
	return &Piece{
		rotation:                    0,
		center:                      [2]int{},
		color:                       color.RGBA{R: 0, G: 255, B: 255, A: 255},
		blockPositionsFromRotation:  IPieceBlockPositions(),
		clockWiseRotationMap:        IClockwiseRotationMap(),
		counterClockWiseRotationMap: ICounterClockwiseRotationMap(),
	}
}

type OPiece struct{}

func OPieceClockwiseRotationMap() [4][][2]int {
	return [4][][2]int{
		0: {},
		1: {},
		2: {},
		3: {},
	}
}

func OPieceCounterClockwiseRotationMap() [4][][2]int {
	return [4][][2]int{
		0: {},
		1: {},
		2: {},
		3: {},
	}
}

func OPieceBlockPositions() (res [4][][2]int) {
	for i := 0; i < 4; i++ {
		res[i] = [][2]int{
			{0, 0},
			{1, 0},
			{0, 1},
			{1, 1},
		}
	}
	return
}

func NewOPiece() *Piece {
	return &Piece{
		color:                       color.RGBA{R: 255, G: 255, B: 0, A: 255},
		center:                      [2]int{},
		rotation:                    0,
		blockPositionsFromRotation:  OPieceBlockPositions(),
		clockWiseRotationMap:        OPieceClockwiseRotationMap(),
		counterClockWiseRotationMap: OPieceCounterClockwiseRotationMap(),
	}
}

func SPieceBlockPositions() [4][][2]int {
	return [4][][2]int{
		0: {{-1, 0}, {0, 0}, {0, 1}, {1, 1}},
		1: {{0, 1}, {0, 0}, {1, 0}, {1, -1}},
		2: {{-1, -1}, {0, -1}, {0, 0}, {1, 0}},
		3: {{-1, 1}, {-1, 0}, {0, 0}, {0, -1}},
	}
}
func NewSPiece() *Piece {
	return &Piece{
		color:                       color.RGBA{R: 0, G: 255, B: 0, A: 255},
		rotation:                    0,
		center:                      [2]int{},
		clockWiseRotationMap:        JLTSZClockwiseRotationMap(),
		counterClockWiseRotationMap: JLTSZCounterClockwiseRotationMap(),
		blockPositionsFromRotation:  SPieceBlockPositions(),
	}
}

func ZPieceBlockPositions() [4][][2]int {
	return [4][][2]int{
		0: {{-1, 1}, {0, +1}, {0, 0}, {1, 0}},
		1: {{1, 1}, {1, 0}, {0, 0}, {0, -1}},
		2: {{-1, 0}, {0, 0}, {0, -1}, {1, -1}},
		3: {{-1, -1}, {-1, 0}, {0, 0}, {0, 1}},
	}
}
func NewZPiece() *Piece {
	return &Piece{
		color:                       color.RGBA{R: 255, G: 0, B: 0, A: 255},
		center:                      [2]int{},
		rotation:                    0,
		blockPositionsFromRotation:  ZPieceBlockPositions(),
		clockWiseRotationMap:        JLTSZClockwiseRotationMap(),
		counterClockWiseRotationMap: JLTSZCounterClockwiseRotationMap(),
	}
}

func JPieceBlockPositions() [4][][2]int {
	return [4][][2]int{
		0: {{-1, 1}, {-1, 0}, {0, 0}, {1, 0}},
		1: {{1, 1}, {0, 1}, {0, 0}, {0, -1}},
		2: {{-1, 0}, {0, 0}, {1, 0}, {1, -1}},
		3: {{-1, -1}, {0, -1}, {0, 0}, {0, 1}},
	}
}

func NewJPiece() *Piece {
	return &Piece{
		color:                       color.RGBA{R: 0, G: 0, B: 255, A: 255},
		rotation:                    0,
		center:                      [2]int{},
		blockPositionsFromRotation:  JPieceBlockPositions(),
		clockWiseRotationMap:        JLTSZClockwiseRotationMap(),
		counterClockWiseRotationMap: JLTSZCounterClockwiseRotationMap(),
	}
}

func LPieceBlockPositions() [4][][2]int {
	return [4][][2]int{
		0: {{-1, 0}, {0, 0}, {1, 0}, {1, 1}},
		1: {{0, 1}, {0, 0}, {0, -1}, {1, -1}},
		2: {{-1, 0}, {0, 0}, {1, 0}, {-1, -1}},
		3: {{0, 1}, {0, 0}, {0, -1}, {-1, 1}},
	}
}
func NewLPiece() *Piece {
	return &Piece{
		color:                       color.RGBA{R: 255, G: 170, B: 0, A: 255},
		rotation:                    0,
		center:                      [2]int{},
		blockPositionsFromRotation:  LPieceBlockPositions(),
		clockWiseRotationMap:        JLTSZClockwiseRotationMap(),
		counterClockWiseRotationMap: JLTSZCounterClockwiseRotationMap(),
	}
}
