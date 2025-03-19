package piece

import (
	"iter"
)

type Queue struct {
	queue [5]Piece
	start int
	bag   SevenBag
}

func (pq *Queue) Next() Piece {
	piece := pq.queue[pq.start]
	pq.queue[pq.start] = pq.bag.getPiece()
	pq.start = (pq.start + 1) % 5
	return piece
}

func (pq *Queue) Pieces() iter.Seq2[int, Piece] {
	return func(yield func(int, Piece) bool) {
		for i := range 5 {
			if !yield(i, pq.queue[(pq.start+i)%5]) {
				break
			}
		}
	}
}

func NewQueue() Queue {
	bag := newSevenBag()
	queue := [5]Piece{}
	for i := range 5 {
		queue[i] = bag.getPiece()
	}

	return Queue{
		queue: queue,
		start: 0,
		bag:   bag,
	}
}
