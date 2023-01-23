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
	hourHandLength   = 50
)

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}

func makeHand(p Point, length float64) Point {
	// scale
	p = Point{p.X * length, p.Y * length}
	// flip
	p = Point{p.X, -p.Y}
	// translate
	p = Point{p.X + clockCenterX, p.Y + clockCenterY}

	return p
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
	// make second hand
	p := makeHand(SecondHandPoint(t), secondHandLength)
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
	// make minute hand
	p := makeHand(MinuteHandPoint(t), minuteHandLength)
	// write point to writer
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#222;stroke-width:5px;"/>`, p.X, p.Y)
}

func HoursInRadians(t time.Time) float64 {
	return (MinutesInRadians(t) / 12) + (math.Pi / (6 / float64(t.Hour()%12)))
}

func HourHandPoint(t time.Time) Point {
	return angleToPoint(HoursInRadians(t))
}

// HourHand writes the SVG line for the hour hand of an analogue clock
// at time `t` represented as a point
func HourHand(w io.Writer, t time.Time) {
	// make hour hand
	p := makeHand(HourHandPoint(t), hourHandLength)
	// write point to writer
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:7px;"/>`, p.X, p.Y)
}
