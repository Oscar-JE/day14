package particle

import (
	"day14/common"
)

type Sand struct {
	position common.Point
}

func Init(p common.Point) Sand {
	return Sand{position: p}
}

func (s Sand) FallingPatern() []common.Point {
	down := common.InitPoint(0, 1)
	leftDown := common.InitPoint(-1, 1)
	rightDown := common.InitPoint(1, 1)
	return []common.Point{s.position.Add(down), s.position.Add(leftDown), s.position.Add(rightDown)}
}

func (s *Sand) SetPosition(p common.Point) {
	s.position = p
}

func (s Sand) GetPosition() common.Point {
	return s.position
}
