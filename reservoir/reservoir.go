package reservoir

import (
	"day14/common"
	"day14/reservoir/matrix"
)

type Reservoir struct {
	offset common.Point
	matrix matrix.Matrix
}

func InitFromIntervalls(intervals []common.ClosedInterval) Reservoir {
	maxRow, minCol, maxCol := findMaxAndMin(intervals)
	offset := common.InitPoint(0, minCol)
	m := matrix.Init(maxRow, maxCol-minCol)
	return Reservoir{offset: offset, matrix: m}
}

func findMaxAndMin(intervals []common.ClosedInterval) (int, int, int) {
	maxRow := intervals[0].MaxRow()
	minCol := intervals[0].MinCol()
	maxCol := intervals[0].MaxCol()
	for i := 1; i < len(intervals); i++ {
		maxRow = Max(maxRow, intervals[i].MaxRow())
		minCol = Min(minCol, intervals[i].MinCol())
		maxCol = Max(maxCol, intervals[i].MaxCol())
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
