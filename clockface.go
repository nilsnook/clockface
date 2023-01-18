package clockface

import (
	"math"
	"time"
)

// A Point represents a two dimensional Cartesian coordiante
type Point struct {
	X float64
	Y float64
}

const secondHandLength = 90
const clockCenterX = 150
const clockCenterY = 150

// SecondHand is the unit vector of the second hand of an analogue clock
// at time `t` represented as a Point
func SecondHand(t time.Time) Point {
	p := SecondHandPoint(t)
	// scale
	p = Point{p.X * secondHandLength, p.Y * secondHandLength}
	// flip
	p = Point{p.X, -p.Y}
	// translate
	p = Point{p.X + clockCenterX, p.Y + clockCenterY}

	// return Point{150, 60}
	return p
}

func SecondsInRadians(t time.Time) float64 {
	// return float64(t.Second()) * (math.Pi / 30)
	return (math.Pi / (30 / float64(t.Second())))
}

func SecondHandPoint(t time.Time) Point {
	angle := SecondsInRadians(t)
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}
