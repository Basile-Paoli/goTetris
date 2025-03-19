package piece

import "tetris/grid"

func (p *Piece) decrementRotation() {
	p.rotation = (p.rotation + 3) % 4
}

func (p *Piece) incrementRotation() {
	p.rotation = (p.rotation + 1) % 4
}

func (p *Piece) copy() Piece {
	return *p
}

func (p *Piece) RotateClockwise(grid grid.Grid) {
	rotatedPiece := p.copy()

	rotatedPiece.incrementRotation()
	for _, translationVector := range rotatedPiece.clockWiseRotationMap[rotatedPiece.rotation] {
		if rotatedPiece.isValidMovement(grid, translationVector) {
			p.rotation = rotatedPiece.rotation
			p.center.X += translationVector.x
			p.center.Y += translationVector.y
			return
		}
	}
}

func (p *Piece) RotateCounterClockwise(grid grid.Grid) {
	rotatedPiece := p.copy()
	rotatedPiece.decrementRotation()

	for _, translationVector := range p.counterClockWiseRotationMap[rotatedPiece.rotation] {
		if rotatedPiece.isValidMovement(grid, translationVector) {
			p.rotation = rotatedPiece.rotation
			p.center.X += translationVector.x
			p.center.Y += translationVector.y
			return
		}
	}
}

var jltszClockwiseRotationMap = [4][]vector{
	0: {{0, 0}, {-1, 0}, {-1, -1}, {0, 2}, {-1, 2}},
	1: {{0, 0}, {-1, 0}, {-1, 1}, {0, -2}, {-1, -2}},
	2: {{0, 0}, {1, 0}, {1, -1}, {0, 2}, {1, 2}},
	3: {{0, 0}, {1, 0}, {1, 1}, {0, -2}, {1, -2}},
}

var jltszCounterClockwiseRotationMap = [4][]vector{
	0: {{0, 0}, {1, 0}, {1, -1}, {0, 2}, {1, 2}},
	1: {{0, 0}, {-1, 0}, {-1, 1}, {0, -2}, {-1, -2}},
	2: {{0, 0}, {-1, 0}, {-1, -1}, {0, 2}, {-1, 2}},
	3: {{0, 0}, {1, 0}, {1, 1}, {0, -2}, {1, -2}},
}

var iClockwiseRotationMap = [4][]vector{
	0: {{0, 0}, {1, 0}, {-2, 0}, {1, -2}, {-2, 1}},
	1: {{0, 0}, {-2, 0}, {1, 0}, {-2, -1}, {1, 2}},
	2: {{0, 0}, {-1, 0}, {2, 0}, {-1, 2}, {2, -1}},
	3: {{0, 0}, {2, 0}, {-1, 0}, {2, 1}, {-1, -2}},
}

var iCounterClockwiseRotationMap = [4][]vector{
	0: {{0, 0}, {2, 0}, {-1, 0}, {2, 1}, {-1, -2}},
	1: {{0, 0}, {1, 0}, {-2, 0}, {1, -2}, {-2, 1}},
	2: {{0, 0}, {-2, 0}, {1, 0}, {-2, -1}, {1, 2}},
	3: {{0, 0}, {-1, 0}, {2, 0}, {-1, 2}, {2, -1}},
}

var oPieceClockwiseRotationMap = [4][]vector{}

var oPieceCounterClockwiseRotationMap = [4][]vector{}
