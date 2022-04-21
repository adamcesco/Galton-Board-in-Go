package structs

import (
	"math/rand"
	"time"
)

type bead struct {
	x    int
	y    int
	maxX int
}

func (b *bead) advance() (int, int) {
	rand.Seed(time.Now().UnixNano())
	dir := rand.Int() % 3
	if dir == 0 && b.x > 0 {
		b.x--
	} else if dir == 1 && b.x < b.maxX {
		b.x++
	}
	b.y--

	return b.x, b.y
}

func (b *bead) reset(size, layers int) (int, int) {
	b.maxX = size - 1
	(*b).x = size / 2
	(*b).y = layers
	return (*b).x, (*b).y
}
