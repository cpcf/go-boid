package main

import "math/rand/v2"

type boid struct {
	pos           Point
	nextPos       Point
	vel           Point
	forward       Point
	maxX, maxY    float64
	bounce        bool
	clampMinSpeed bool
}

func initBoidsOnScreenSize(screenWidth, screenHeight int) []boid {
	count := screenWidth*screenHeight/125 + 1
	return initRandomBoids(count, screenWidth, screenHeight)
}

func initRandomBoids(count int, screenWidth, screenHeight int) []boid {
	boids := make([]boid, count)
	for i := range boids {
		boids[i] = boid{
			pos: Point{
				x: rand.Float64() * float64(screenWidth),
				y: rand.Float64() * float64(screenHeight),
			},
			vel: Point{
				x: rand.Float64()*maxSpeed*2 - maxSpeed,
				y: rand.Float64()*maxSpeed*2 - maxSpeed,
			},
			maxX:          float64(screenWidth),
			maxY:          float64(screenHeight),
			bounce:        bouncey,
			clampMinSpeed: clampMinSpeed,
		}
	}
	return boids
}

func (b *boid) update(boids []boid) {

	accel := b.calcAcceleration(boids)

	b.vel = b.vel.Add(accel).Limit(-maxSpeed, maxSpeed)

	if b.clampMinSpeed {
		if b.vel.x >= 0 && b.vel.x < targetMinSpeed {
			b.vel.x = Lerp(b.vel.x, targetMinSpeed, 0.5)
		}
		if b.vel.x < 0 && b.vel.x > -targetMinSpeed {
			b.vel.x = Lerp(b.vel.x, -targetMinSpeed, 0.5)
		}
		if b.vel.y >= 0 && b.vel.y < targetMinSpeed {
			b.vel.y = Lerp(b.vel.y, targetMinSpeed, 0.5)
		}
		if b.vel.y < 0 && b.vel.y > -targetMinSpeed {
			b.vel.y = Lerp(b.vel.y, -targetMinSpeed, 0.5)
		}
	}

	b.nextPos = b.pos.Add(b.vel)
	if !b.bounce {
		b.wrapAroundScreen()
	}

	updateForward(b)

}

func (b *boid) wrapAroundScreen() {
	if b.nextPos.x < 0 {
		b.nextPos.x += b.maxX
	} else if b.nextPos.x > b.maxX {
		b.nextPos.x -= b.maxX
	}

	if b.nextPos.y < 0 {
		b.nextPos.y += b.maxY
	} else if b.nextPos.y > b.maxY {
		b.nextPos.y -= b.maxY
	}
}

func (b *boid) move() {
	b.pos = b.nextPos
}

func (b *boid) calcAcceleration(boids []boid) Point {
	accel := Point{}
	if b.bounce {
		accel = Point{bounce(b.pos.x, b.maxX), bounce(b.pos.y, b.maxY)}
	}

	var accelCohesion, accelSeparation, accelAlignment Point

	sep, avgPos, avgVel, count := b.measureNearby(boids)

	if count == 0 {
		return accel
	}
	avgPos = avgPos.DivideV(float64(count))
	avgVel = avgVel.DivideV(float64(count))

	accelAlignment = avgVel.Subtract(b.vel).MultiplyV(adjustRate).MultiplyV(alignmentRate)
	accelCohesion = avgPos.Subtract(b.pos).MultiplyV(adjustRate).MultiplyV(cohesionRate)
	accelSeparation = sep.MultiplyV(adjustRate).MultiplyV(separationRate)

	accel = accel.Add(accelAlignment).Add(accelCohesion).Add(accelSeparation)
	return accel
}

func (b *boid) measureNearby(boids []boid) (Point, Point, Point, int) {
	var sep, avgPos, avgVel Point
	count := 0

	for _, other := range boids {
		if other.pos.x == b.pos.x && other.pos.y == b.pos.y {
			continue
		}
		if b.pos.Distance(other.pos) < radius {
			count++
			avgVel = avgVel.Add(other.vel)
			avgPos = avgPos.Add(other.pos)
			sep = sep.Add(b.pos.Subtract(other.pos).DivideV(b.pos.Distance(other.pos) * 1.5))
		}
	}
	return sep, avgPos, avgVel, count
}

func bounce(pos, maxBorderPos float64) float64 {
	if pos < radius {
		return 1 / pos
	} else if pos > maxBorderPos-radius {
		return 1 / (pos - maxBorderPos)
	}
	return 0
}

func updateForward(b *boid) {
	b.forward = b.vel.Normalize()
}
