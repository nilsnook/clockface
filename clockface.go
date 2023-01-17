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

// SecondHand is the unit vector of the second hand of an analogue clock
// at time `t` represented as a Point
func SecondHand(t time.Time) Point {
	return Point{150, 60}
}

func SecondsInRadians(t time.Time) float64 {
	// return float64(t.Second()) * (math.Pi / 30)
	return (math.Pi / (30 / float64(t.Second())))
}
