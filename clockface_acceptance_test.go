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

			want := c.line

			for _, got := range svg.Line {
				if got == want {
					return
				}
			}

			if !containsLine(want, svg.Line) {
				t.Errorf("In SVG output with Lines %+v, want second hand with Line %+v", svg.Line, want)
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
