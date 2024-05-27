package main

import "image/color"

type Piece interface {
	Color() color.Color
	RotateClockwise()
	RotateCounterClockwise()
	BlockCoordinates() [4][2]int
	Drop()
	MoveLeft()
	MoveRight()
}

type TPiece struct {
	center   [2]int
	rotation int
}

func (t *TPiece) Color() color.Color {
	return color.RGBA{R: 128, G: 0, B: 128, A: 255}
}

func (t *TPiece) RotateClockwise() {
	t.rotation = (t.rotation + 1) % 4
}

func (t *TPiece) RotateCounterClockwise() {
	t.rotation = (t.rotation + 3) % 4
}
func (t *TPiece) Drop() {
	t.center[1]--
}

func (t *TPiece) BlockCoordinates() [4][2]int {
	coordinates := [4][2]int{}
	switch t.rotation {
	case 0:
		coordinates = [4][2]int{
			{t.center[0] - 1, t.center[1]},
			{t.center[0], t.center[1]},
			{t.center[0] + 1, t.center[1]},
			{t.center[0], t.center[1] + 1},
		}
	case 1:
		coordinates = [4][2]int{
			{t.center[0], t.center[1] - 1},
			{t.center[0], t.center[1]},
			{t.center[0], t.center[1] + 1},
			{t.center[0] + 1, t.center[1]},
		}
	case 2:
		coordinates = [4][2]int{
			{t.center[0] - 1, t.center[1]},
			{t.center[0], t.center[1]},
			{t.center[0] + 1, t.center[1]},
			{t.center[0], t.center[1] - 1},
		}
	case 3:
		coordinates = [4][2]int{
			{t.center[0], t.center[1] - 1},
			{t.center[0], t.center[1]},
			{t.center[0], t.center[1] + 1},
			{t.center[0] - 1, t.center[1]},
		}
	}
	return coordinates
}
func (t *TPiece) MoveLeft() {
	t.center[0]--
}
func (t *TPiece) MoveRight() {
	t.center[0]++
}
func NewTPiece() *TPiece {
	return &TPiece{center: [2]int{gridWidth/2 - 1, gridHeight}, rotation: 0}
}

type IPiece struct {
	center   [2]int
	rotation int
}

func (i *IPiece) Color() color.Color {
	return color.RGBA{R: 0, G: 128, B: 128, A: 255}
}
func (i *IPiece) RotateClockwise() {
	i.rotation = (i.rotation + 1) % 4
}
func (i *IPiece) RotateCounterClockwise() {
	i.rotation = (i.rotation + 3) % 4
}
func (i *IPiece) Drop() {
	i.center[1]--
}
func (i *IPiece) BlockCoordinates() [4][2]int {
	coordinates := [4][2]int{}
	switch i.rotation {
	case 0:
		coordinates = [4][2]int{
			{i.center[0] - 2, i.center[1] + 1},
			{i.center[0] - 1, i.center[1] + 1},
			{i.center[0], i.center[1] + 1},
			{i.center[0] + 1, i.center[1] + 1},
		}
	case 1:
		coordinates = [4][2]int{
			{i.center[0], i.center[1] - 1},
			{i.center[0], i.center[1]},
			{i.center[0], i.center[1] + 1},
			{i.center[0], i.center[1] + 2},
		}
	case 2:
		coordinates = [4][2]int{
			{i.center[0] - 2, i.center[1]},
			{i.center[0] - 1, i.center[1]},
			{i.center[0], i.center[1]},
			{i.center[0] + 1, i.center[1]},
		}
	case 3:
		coordinates = [4][2]int{
			{i.center[0] - 1, i.center[1] - 1},
			{i.center[0] - 1, i.center[1]},
			{i.center[0] - 1, i.center[1] + 1},
			{i.center[0] - 1, i.center[1] + 2},
		}

	}
	return coordinates
}
func (i *IPiece) MoveLeft() {
	i.center[0]--
}
func (i *IPiece) MoveRight() {
	i.center[0]++
}
func NewIPiece() *IPiece {
	return &IPiece{center: [2]int{gridWidth / 2, gridHeight}, rotation: 0}
}

type OPiece struct {
	center [2]int
}

func (o *OPiece) Color() color.Color {
	return color.RGBA{R: 128, G: 128, B: 0, A: 255}
}
func (o *OPiece) RotateClockwise()        {}
func (o *OPiece) RotateCounterClockwise() {}
func (o *OPiece) Drop() {
	o.center[1]--
}
func (o *OPiece) BlockCoordinates() [4][2]int {
	return [4][2]int{
		{o.center[0], o.center[1]},
		{o.center[0] + 1, o.center[1]},
		{o.center[0], o.center[1] + 1},
		{o.center[0] + 1, o.center[1] + 1},
	}
}
func (o *OPiece) MoveLeft() {
	o.center[0]--
}
func (o *OPiece) MoveRight() {
	o.center[0]++
}
func NewOPiece() *OPiece {
	return &OPiece{center: [2]int{gridWidth/2 - 1, gridHeight}}
}

type SPiece struct {
	center   [2]int
	rotation int
}

func (s *SPiece) Color() color.Color {
	return color.RGBA{R: 0, G: 128, B: 0, A: 255}
}
func (s *SPiece) RotateClockwise() {
	s.rotation = (s.rotation + 1) % 4
}
func (s *SPiece) RotateCounterClockwise() {
	s.rotation = (s.rotation + 3) % 4
}
func (s *SPiece) Drop() {
	s.center[1]--
}
func (s *SPiece) BlockCoordinates() [4][2]int {
	coordinates := [4][2]int{}
	switch s.rotation {
	case 0:
		coordinates = [4][2]int{
			{s.center[0] - 1, s.center[1]},
			{s.center[0], s.center[1]},
			{s.center[0], s.center[1] + 1},
			{s.center[0] + 1, s.center[1] + 1},
		}
	case 1:
		coordinates = [4][2]int{
			{s.center[0], s.center[1] + 1},
			{s.center[0], s.center[1]},
			{s.center[0] + 1, s.center[1]},
			{s.center[0] + 1, s.center[1] - 1},
		}
	case 2:
		coordinates = [4][2]int{
			{s.center[0] - 1, s.center[1] - 1},
			{s.center[0], s.center[1] - 1},
			{s.center[0], s.center[1]},
			{s.center[0] + 1, s.center[1]},
		}
	case 3:
		coordinates = [4][2]int{
			{s.center[0] - 1, s.center[1] + 1},
			{s.center[0] - 1, s.center[1]},
			{s.center[0], s.center[1]},
			{s.center[0], s.center[1] - 1},
		}
	}
	return coordinates
}
func (s *SPiece) MoveLeft() {
	s.center[0]--
}
func (s *SPiece) MoveRight() {
	s.center[0]++
}
func NewSPiece() *SPiece {
	return &SPiece{center: [2]int{gridWidth/2 - 1, gridHeight}, rotation: 0}
}

type ZPiece struct {
	center   [2]int
	rotation int
}

func (z *ZPiece) Color() color.Color {
	return color.RGBA{R: 128, G: 0, B: 0, A: 255}
}
func (z *ZPiece) RotateClockwise() {
	z.rotation = (z.rotation + 1) % 4
}
func (z *ZPiece) RotateCounterClockwise() {
	z.rotation = (z.rotation + 3) % 4
}
func (z *ZPiece) Drop() {
	z.center[1]--
}
func (z *ZPiece) BlockCoordinates() [4][2]int {
	coordinates := [4][2]int{}
	switch z.rotation {
	case 0:
		coordinates = [4][2]int{
			{z.center[0] - 1, z.center[1] + 1},
			{z.center[0], z.center[1] + 1},
			{z.center[0], z.center[1]},
			{z.center[0] + 1, z.center[1]},
		}
	case 1:
		coordinates = [4][2]int{
			{z.center[0] + 1, z.center[1] + 1},
			{z.center[0] + 1, z.center[1]},
			{z.center[0], z.center[1]},
			{z.center[0], z.center[1] - 1},
		}
	case 2:
		coordinates = [4][2]int{
			{z.center[0] - 1, z.center[1] - 1},
			{z.center[0], z.center[1] - 1},
			{z.center[0], z.center[1]},
			{z.center[0] + 1, z.center[1]},
		}
	case 3:
		coordinates = [4][2]int{
			{z.center[0] - 1, z.center[1] - 1},
			{z.center[0] - 1, z.center[1]},
			{z.center[0], z.center[1]},
			{z.center[0], z.center[1] + 1},
		}
	}
	return coordinates
}
func (z *ZPiece) MoveLeft() {
	z.center[0]--
}
func (z *ZPiece) MoveRight() {
	z.center[0]++
}
func NewZPiece() *ZPiece {
	return &ZPiece{center: [2]int{gridWidth/2 - 1, gridHeight}, rotation: 0}
}

type JPiece struct {
	center   [2]int
	rotation int
}

func (j *JPiece) Color() color.Color {
	return color.RGBA{R: 0, G: 0, B: 128, A: 255}
}
func (j *JPiece) RotateClockwise() {
	j.rotation = (j.rotation + 1) % 4
}
func (j *JPiece) RotateCounterClockwise() {
	j.rotation = (j.rotation + 3) % 4
}
func (j *JPiece) Drop() {
	j.center[1]--
}
func (j *JPiece) BlockCoordinates() [4][2]int {
	coordinates := [4][2]int{}
	switch j.rotation {
	case 0:
		coordinates = [4][2]int{
			{j.center[0] - 1, j.center[1] + 1},
			{j.center[0] - 1, j.center[1]},
			{j.center[0], j.center[1]},
			{j.center[0] + 1, j.center[1]},
		}
	case 1:
		coordinates = [4][2]int{
			{j.center[0] + 1, j.center[1] + 1},
			{j.center[0], j.center[1] + 1},
			{j.center[0], j.center[1]},
			{j.center[0], j.center[1] - 1},
		}
	case 2:
		coordinates = [4][2]int{
			{j.center[0] - 1, j.center[1]},
			{j.center[0], j.center[1]},
			{j.center[0] + 1, j.center[1]},
			{j.center[0] + 1, j.center[1] - 1},
		}
	case 3:
		coordinates = [4][2]int{
			{j.center[0] - 1, j.center[1] - 1},
			{j.center[0], j.center[1] - 1},
			{j.center[0], j.center[1]},
			{j.center[0], j.center[1] + 1},
		}
	}
	return coordinates
}
func (j *JPiece) MoveLeft() {
	j.center[0]--
}
func (j *JPiece) MoveRight() {
	j.center[0]++
}
func NewJPiece() *JPiece {
	return &JPiece{center: [2]int{gridWidth/2 - 1, gridHeight}, rotation: 0}
}

type LPiece struct {
	center   [2]int
	rotation int
}

func (l *LPiece) Color() color.Color {
	return color.RGBA{R: 255, G: 128, B: 0, A: 255}

}
func (l *LPiece) RotateClockwise() {
	l.rotation = (l.rotation + 1) % 4

}
func (l *LPiece) RotateCounterClockwise() {
	l.rotation = (l.rotation + 3) % 4
}
func (l *LPiece) Drop() {
	l.center[1]--
}
func (l *LPiece) BlockCoordinates() [4][2]int {
	coordinates := [4][2]int{}
	switch l.rotation {
	case 0:
		coordinates = [4][2]int{
			{l.center[0] - 1, l.center[1]},
			{l.center[0], l.center[1]},
			{l.center[0] + 1, l.center[1]},
			{l.center[0] + 1, l.center[1] + 1},
		}
	case 1:
		coordinates = [4][2]int{
			{l.center[0], l.center[1] + 1},
			{l.center[0], l.center[1]},
			{l.center[0], l.center[1] - 1},
			{l.center[0] + 1, l.center[1] - 1},
		}
	case 2:
		coordinates = [4][2]int{
			{l.center[0] - 1, l.center[1]},
			{l.center[0], l.center[1]},
			{l.center[0] + 1, l.center[1]},
			{l.center[0] - 1, l.center[1] - 1},
		}
	case 3:
		coordinates = [4][2]int{
			{l.center[0], l.center[1] + 1},
			{l.center[0], l.center[1]},
			{l.center[0], l.center[1] - 1},
			{l.center[0] - 1, l.center[1] + 1},
		}
	}
	return coordinates
}
func (l *LPiece) MoveLeft() {
	l.center[0]--
}
func (l *LPiece) MoveRight() {
	l.center[0]++
}
func NewLPiece() *LPiece {
	return &LPiece{center: [2]int{gridWidth/2 - 1, gridHeight}, rotation: 0}
}
