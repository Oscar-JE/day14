package common

import (
	"testing"
)

func TestInIntervall(t *testing.T) {
	p1 := Point{0, 4}
	p2 := Point{0, 4}
	interval := Init(p1, p2)
	if !interval.Has(p1) || interval.Has(p1.Add(p2)) {
		t.Errorf("weird behaviour of intervall")
	}
}

func TestContainedPointsSingular(t *testing.T) {
	p1 := Point{0, 4}
	p2 := Point{0, 4}
	interval := Init(p1, p2)
	points := interval.ContainedPoints()
	if len(points) != 1 || !points[0].Eq(p1) {
		t.Errorf("controll Cointained points when end equals start")
	}
}

func TestContainedPointsTwo(t *testing.T) {
	p1 := Point{0, 4}
	p2 := Point{0, 5}
	interval := Init(p1, p2)
	points := interval.ContainedPoints()
	if len(points) != 2 || !points[0].Eq(p1) || !points[1].Eq(p2) {
		t.Errorf("controll Cointained points with two points")
	}
}
