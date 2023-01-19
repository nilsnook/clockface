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

func TestSVGWriterAtMidnight(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

	// var b strings.Builder
	// clockface.SVGWriter(&b, tm)
	// got := b.String()
	// want := `<line x1="150" y1="150" x2="150", y2="60">`
	//
	// if !strings.Contains(got, want) {
	// 	t.Errorf("Expected to find the second hand %v, in the SVG outout %v", want, got)
	// }

	b := bytes.Buffer{}
	clockface.WriteSVG(&b, tm)

	svg := clockface.SVG{}
	xml.Unmarshal(b.Bytes(), &svg)

	x2 := "150.000"
	y2 := "60.000"

	for _, line := range svg.Line {
		if line.X2 == x2 && line.Y2 == y2 {
			return
		}
		t.Logf("In SVG output, got x2=%+v and y2=%+v", line.X2, line.Y2)
	}

	t.Errorf("In SVG output, want second hand with x2=%+v and y2=%+v", x2, y2)
}
