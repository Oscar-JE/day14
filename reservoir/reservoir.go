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
	offset := common.InitPoint(minCol, 0)
	m := matrix.Init(maxRow+1, maxCol-minCol+1)
	reservoir := Reservoir{offset: offset, matrix: m}
	for intervalIndex := range intervals {
		wallPoints := intervals[intervalIndex].ContainedPoints()
		for wallindex := range wallPoints {
			point := wallPoints[wallindex] // här är det lite konstigt
			reservoir.SetWall(point)
		}
	}

	return reservoir
}

func (r *Reservoir) SetWall(point common.Point) {
	point = point.Subtract(r.offset)
	point = point.Transpose()
	r.matrix.SetWall(point.GetRow(), point.GetCol())
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

func (r Reservoir) OnField(point common.Point) bool {
	internalPoint := point.Subtract(r.offset).Transpose()
	return r.matrix.InMatrix(internalPoint.GetRow(), internalPoint.GetCol())
}

func (r Reservoir) IsBlocked(point common.Point) bool {
	internalPoint := point.Subtract(r.offset).Transpose()
	return r.matrix.IsBlocked(internalPoint.GetRow(), internalPoint.GetCol())
}