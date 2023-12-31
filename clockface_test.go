package clockface_test

import (
	"clockface"
	"math"
	"testing"
	"time"
)

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(2023, time.January, 28, hours, minutes, seconds, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}

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
			if !roughlyEqualFloat64(got, want) {
				t.Fatalf("got %v, want %v", got, want)
			}
		})
	}
}

func TestSecondHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point clockface.Point
	}{
		{simpleTime(0, 0, 30), clockface.Point{0, -1}},
		{simpleTime(0, 0, 45), clockface.Point{-1, 0}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := clockface.SecondHandPoint(c.time)
			want := c.point

			if !roughlyEqualPoint(got, want) {
				t.Fatalf("got %v, want %v", got, want)
			}
		})
	}
}

func TestMinuteHandInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 30, 0), math.Pi},
		{simpleTime(0, 0, 7), (math.Pi / (30 * 60)) * 7},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := clockface.MinutesInRadians(c.time)
			want := c.angle
			if !roughlyEqualFloat64(got, want) {
				t.Fatalf("got %v, want %v", got, want)
			}
		})
	}
}

func TestMinuteHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point clockface.Point
	}{
		{simpleTime(0, 30, 0), clockface.Point{0, -1}},
		{simpleTime(0, 45, 0), clockface.Point{-1, 0}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := clockface.MinuteHandPoint(c.time)
			want := c.point

			if !roughlyEqualPoint(got, want) {
				t.Fatalf("got %v, want %v", got, want)
			}
		})
	}
}

func TestHourHandInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(6, 0, 0), math.Pi},
		{simpleTime(0, 0, 0), 0},
		{simpleTime(21, 0, 0), math.Pi * 1.5},
		{simpleTime(0, 1, 30), math.Pi / ((6 * 60 * 60) / 90)},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := clockface.HoursInRadians(c.time)
			want := c.angle
			if !roughlyEqualFloat64(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

func TestHourHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point clockface.Point
	}{
		{simpleTime(6, 0, 0), clockface.Point{0, -1}},
		{simpleTime(21, 0, 0), clockface.Point{-1, 0}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := clockface.HourHandPoint(c.time)
			want := c.point
			if !roughlyEqualPoint(got, want) {
				t.Fatalf("got %v, want %v", got, want)
			}
		})
	}
}

func roughlyEqualFloat64(a, b float64) bool {
	const equalityThreshold = 1e-7
	return math.Abs(a-b) < equalityThreshold
}

func roughlyEqualPoint(a, b clockface.Point) bool {
	return roughlyEqualFloat64(a.X, b.X) && roughlyEqualFloat64(a.Y, b.Y)
}
