package common

import (
	"testing"
)

func TestAddPoint(t *testing.T) {
	p1 := Point{1, 2}
	p2 := Point{3, 4}
	p3 := p1.Add(p2)
	if p3.row != 4 || p3.col != 6 {
		t.Errorf("Wrong behaviour in point add")
	}
}

func TestSubtractPoint(t *testing.T) {
	p1 := Point{1, 2}
	p2 := Point{3, 4}
	p3 := p2.Subtract(p1)
	if p3.row != 2 && p3.col != 2 {
		t.Errorf("Wrong behaviour in point subtract")
	}
}

func TestAbove(t *testing.T) {
	p1 := Point{1, 2}
	p2 := Point{3, 4}
	if p2.Above(p1) {
		t.Errorf("wtf1")
	}
	p1 = Point{1, 0}
	p2 = Point{3, 0}
	if !p1.Above(p2) {
		t.Errorf("p1 should be over p2")
	}
} // resterande får vi testa senare med debugg
