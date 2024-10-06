package main

type boid struct {
	x, y                 int
	xVelocity, yVelocity int
	maxX, maxY           int
}

func (b *boid) update() {
	b.x += b.xVelocity
	b.y += b.yVelocity

	b.x = (b.x + b.maxX) % b.maxX
	b.y = (b.y + b.maxY) % b.maxY

}
