package structs

import (
	"math/rand"
	"time"
)

type bead struct {
	x int
	y int
}

func (b *bead) Advance() (int, int) {
	rand.Seed(time.Now().UnixNano())
	dir := rand.Int() % 3
	if dir == 0 {
		b.x--
	} else if dir == 1 {
		b.x++
	}
	b.y--

	return b.x, b.y
}

func (b *bead) Reset(size, layers int) (int, int) {
	(*b).x = size / 2
	(*b).y = layers
	return (*b).x, (*b).y
}

func (b bead) ReadX() int {
	return b.x
}
