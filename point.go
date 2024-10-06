package main

import "math"

type Point struct {
	x float64
	y float64
}

// Adds x and y of two vectors
func (v1 Point) Add(v2 Point) Point {
	return Point{x: v1.x + v2.x, y: v1.y + v2.y}
}

// Subtracts x and y of two vectors
func (v1 Point) Subtract(v2 Point) Point {
	return Point{x: v1.x - v2.x, y: v1.y - v2.y}
}

// Multiplies x and y of two vectors
func (v1 Point) Multiply(v2 Point) Point {
	return Point{x: v1.x * v2.x, y: v1.y * v2.y}
}

// Adds a value v to a Vector
func (vect Point) AddV(val float64) Point {
	return Point{x: vect.x + val, y: vect.y + val}
}

// Multiplies a value v to a Vector
func (vect Point) MultiplyV(val float64) Point {
	return Point{x: vect.x * val, y: vect.y * val}
}

// Divides a value v by a Vector
func (vect Point) DivideV(val float64) Point {
	return Point{x: vect.x / val, y: vect.y / val}
}

// Limits a vector to the passed lower and upper bounds
func (vect Point) Limit(lower, upper float64) Point {
	return Point{x: math.Min(math.Max(vect.x, lower), upper), y: math.Min(math.Max(vect.y, lower), upper)}
}

// Calculates the distances between two vectors
func (v1 Point) Distance(v2 Point) float64 {
	return math.Sqrt(math.Pow(v1.x-v2.x, 2) + math.Pow(v1.y-v2.y, 2))
}
