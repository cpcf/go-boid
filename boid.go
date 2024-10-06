package main

type boid struct {
	x, y                 int
	xVelocity, yVelocity int
}

func (b *boid) update() {
	b.x += b.xVelocity
	b.y += b.yVelocity
}
