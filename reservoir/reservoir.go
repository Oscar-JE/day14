package reservoir

import (
	"day14/common"
	"day14/reservoir/matrix"
	"fmt"
)

type Reservoir struct {
	offset common.Point
	matrix matrix.Matrix
}

func InitFromIntervalls(intervals []common.ClosedInterval) Reservoir {
	maxRow, minCol, maxCol := findMaxAndMin(intervals)
	offset := common.InitPoint(0, minCol)
	m := matrix.Init(maxRow, maxCol-minCol+1)
	return Reservoir{offset: offset, matrix: m}
}

func findMaxAndMin(intervals []common.ClosedInterval) (int, int, int) {
	maxRow := intervals[0].MaxCol()
	minCol := intervals[0].MinRow()
	maxCol := intervals[0].MaxRow()
	for i := 1; i < len(intervals); i++ {
		maxRow = Max(maxRow, intervals[i].MaxCol())
		minCol = Min(minCol, intervals[i].MinRow())
		maxCol = Max(maxCol, intervals[i].MaxRow())
	}
	return maxRow, minCol, maxCol
}

func Max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func (r Reservoir) String() string {
	return fmt.Sprint(r.matrix)
}
