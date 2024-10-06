package main

import "math"

type boid struct {
	pos                  Point
	xVelocity, yVelocity float64
	maxX, maxY           float64
	view                 *View
	behindX, behindY     int
	nearby               []boid
}

var radius = 10
var maxSpeed = 2.0
var adjustRate = 0.015

func (b *boid) update() {

	sepX, sepY := b.calcSeperation(b.nearby)
	b.xVelocity += sepX
	b.yVelocity += sepY

	b.xVelocity = math.Max(math.Min(b.xVelocity, maxSpeed), -maxSpeed)
	b.yVelocity = math.Max(math.Min(b.yVelocity, maxSpeed), -maxSpeed)

	b.pos.x += b.xVelocity
	b.pos.y += b.yVelocity

	b.pos.x = math.Mod(b.pos.x+b.maxX, b.maxX)
	b.pos.y = math.Mod(b.pos.y+b.maxY, b.maxY)

	b.calcBehind()
}

func (b *boid) calcNearby(boids []boid) {
	b.view.Compute(b.viewLimit, int(b.pos.x), int(b.pos.y), radius)

	var nearby []boid
	for _, other := range boids {
		if other.pos.x == b.pos.x && other.pos.y == b.pos.y {
			continue
		}
		if _, ok := b.view.Visible[point{int(other.pos.x), int(other.pos.y)}]; ok {
			nearby = append(nearby, other)
		}
	}
	b.nearby = nearby
}

func (b *boid) calcBehind() {
	max := b.xVelocity
	if b.yVelocity > max {
		max = b.yVelocity
	}

	if max == 0 {
		return
	}

	dirX := b.xVelocity * -1 / max
	dirY := b.yVelocity * -1 / max

	b.behindX = int(b.pos.x + dirX)
	b.behindY = int(b.pos.y + dirY)
}

func (b *boid) viewLimit(x, y int) bool {
	return b.behindX == x && b.behindY == y
}

func (b *boid) calcSeperation(boids []boid) (x, y float64) {
	var xSum, ySum float64
	for _, other := range boids {
		if other.pos.x == b.pos.x && other.pos.y == b.pos.y {
			continue
		}
		xSum += (b.pos.x - other.pos.x) / b.distanceTo(other) * adjustRate
		ySum += (b.pos.y - other.pos.y) / b.distanceTo(other) * adjustRate
	}
	return xSum, ySum
}

func (b *boid) distanceTo(other boid) float64 {
	return math.Sqrt(math.Pow(float64(b.pos.x-other.pos.x), 2) + math.Pow(float64(b.pos.y-other.pos.y), 2))
}
