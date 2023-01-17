package clockface_test

import (
	"clockface"
	"math"
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
//
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

func TestSecondsInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 0, 30), math.Pi},
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 0, 45), (math.Pi / 2) * 3},
		{simpleTime(0, 0, 7), (math.Pi / 30) * 7},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := clockface.SecondsInRadians(c.time)
			want := c.angle
			if got != want {
				t.Fatalf("got %v, want %v", got, want)
			}
		})
	}
}

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(312, time.October, 28, hours, minutes, seconds, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}