package clockface

import (
	"fmt"
	"io"
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
func SecondHand(w io.Writer, t time.Time) {
	p := SecondHandPoint(t)
	// scale
	p = Point{p.X * secondHandLength, p.Y * secondHandLength}
	// flip
	p = Point{p.X, -p.Y}
	// translate
	p = Point{p.X + clockCenterX, p.Y + clockCenterY}

	// write point to writer
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
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
