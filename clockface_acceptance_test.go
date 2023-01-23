package clockface_test

import (
	"bytes"
	"clockface"
	"encoding/xml"
	"testing"
	"time"
)

// func TestSecondHandAtMidnight(t *testing.T) {
// 	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)
//
// 	want := clockface.Point{X: 150, Y: 150 - 90}
// 	got := clockface.SecondHand(tm)
//
// 	if got != want {
// 		t.Errorf("got %v, want %v", got, want)
// 	}
// }

// func TestSecondHandAt30Seconds(t *testing.T) {
// 	tm := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)
//
// 	want := clockface.Point{X: 150, Y: 150 + 90}
// 	got := clockface.SecondHand(tm)
//
// 	if got != want {
// 		t.Errorf("got %v, want %v", got, want)
// 	}
// }

// Test SVG for second hand
func TestSVGWriterSecondHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line clockface.Line
	}{
		{
			simpleTime(0, 0, 0),
			clockface.Line{150, 150, 150, 60},
		},
		{
			simpleTime(0, 0, 30),
			clockface.Line{150, 150, 150, 240},
		},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			b := bytes.Buffer{}
			clockface.WriteSVG(&b, c.time)

			svg := clockface.SVG{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !containsLine(c.line, svg.Line) {
				t.Errorf("In SVG output with Lines %+v, want second hand with Line %+v", svg.Line, c.line)
			}
		})
	}
}

// Test SVG for minute hand
func TestSVGWriterMinuteHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line clockface.Line
	}{
		{simpleTime(0, 0, 0), clockface.Line{150, 150, 150, 70}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			b := bytes.Buffer{}
			clockface.WriteSVG(&b, c.time)

			svg := clockface.SVG{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !containsLine(c.line, svg.Line) {
				t.Errorf("In the SVG output with Lines %+v, wanted a minute hand line %+v", svg.Line, c.line)
			}
		})
	}
}

// Test SVG for hour hand
func TestSVGWriterHourHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line clockface.Line
	}{
		{simpleTime(6, 0, 0), clockface.Line{150, 150, 150, 200}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			b := bytes.Buffer{}
			clockface.WriteSVG(&b, c.time)

			svg := clockface.SVG{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !containsLine(c.line, svg.Line) {
				t.Errorf("In the SVG output with Lines %+v, wanted a hour hand line %+v", svg.Line, c.line)
			}
		})
	}
}

func containsLine(l clockface.Line, lines []clockface.Line) bool {
	for _, line := range lines {
		if line == l {
			return true
		}
	}
	return false
}
