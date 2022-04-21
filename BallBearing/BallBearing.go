package BallBearing

import "math/rand"

type Baller interface {
	Advance() (int, int)
}

type BallBearing struct {
	x int
	y int
}

func (b BallBearing) Advance() (int, int) {
	dir := (rand.Int() % 2) == 0
	if dir == true {
		b.x -= 1
	} else {
		b.x += 1
	}
	b.y++
	return b.x, b.y
}
