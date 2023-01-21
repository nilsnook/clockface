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

const (
	clockCenterX     = 150
	clockCenterY     = 150
	secondHandLength = 90
	minuteHandLength = 80
)

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}

func SecondsInRadians(t time.Time) float64 {
	// return float64(t.Second()) * (math.Pi / 30)
	return (math.Pi / (30 / float64(t.Second())))
}

func SecondHandPoint(t time.Time) Point {
	return angleToPoint(SecondsInRadians(t))
}

// SecondHand writes the SVG line for second hand of an analogue clock
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

func MinutesInRadians(t time.Time) float64 {
	return (SecondsInRadians(t) / 60) +
		(math.Pi / (30 / float64(t.Minute())))
}

func MinuteHandPoint(t time.Time) Point {
	return angleToPoint(MinutesInRadians(t))
}

// MinuteHand writes the SVG line for minute hand of an analogue clock
// at time `t` represented as a Point
func MinuteHand(w io.Writer, t time.Time) {
	p := MinuteHandPoint(t)
	// scale
	p = Point{p.X * minuteHandLength, p.Y * minuteHandLength}
	// flip
	p = Point{p.X, -p.Y}
	// translate
	p = Point{p.X + clockCenterX, p.Y + clockCenterY}

	// write point to writer
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:7px;"/>`, p.X, p.Y)
}
