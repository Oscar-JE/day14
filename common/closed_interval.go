package common

type ClosedInterval struct {
	start    Point
	end      Point
	standing bool
}

func Init(p1 Point, p2 Point) ClosedInterval {
	if p1.Above(p2) {
		return ClosedInterval{p1, p2, true}
	}
	if p1.ToTheLeftOf(p2) {
		return ClosedInterval{p1, p2, false}
	}
	if p2.Above(p1) {
		return ClosedInterval{p2, p1, true}
	}
	if p2.ToTheLeftOf(p1) {
		return ClosedInterval{p2, p1, false}
	}
	panic("Not an accaptable intervall")
}

func (c ClosedInterval) Has(p Point) bool {
	if c.standing {
		return c.standingIn(p)
	} else {
		return c.layingIn(p)
	}
}

func (c ClosedInterval) ContainedPoints() []Point {
	var stepingDirection Point
	if c.standing {
		stepingDirection = Point{1, 0}
	} else {
		stepingDirection = Point{0, 1}
	}
	points := []Point{}
	currentPoint := c.start
	atEnd := currentPoint.Eq(c.end)
	for !atEnd {
		points = append(points, currentPoint)
		currentPoint = currentPoint.Add(stepingDirection)
		atEnd = currentPoint.Eq(c.end)
	}
	points = append(points, c.end)
	return points
}

func (c ClosedInterval) standingIn(p Point) bool {
	return c.start.Above(p) && c.end.Below(p)
}

func (c ClosedInterval) layingIn(p Point) bool {
	return c.start.ToTheLeftOf(p) && c.end.ToTheRightOf(p)
}
