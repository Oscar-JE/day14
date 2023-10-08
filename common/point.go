package common

type Point struct {
	row int
	col int
}

func InitPoint(row int, col int) Point {
	return Point{row, col}
}

func (p Point) Add(other Point) Point {
	newRow := p.row + other.row
	newCol := p.col + other.col
	return Point{newRow, newCol}
}

func (p Point) negate() Point {
	newRow := -p.row
	newCol := -p.col
	return Point{newRow, newCol}
}

func (p Point) Subtract(other Point) Point {
	return p.Add(other.negate())
}

func (p Point) Eq(other Point) bool {
	return p.row == other.row && p.col == other.col
}

func (p Point) Above(other Point) bool {
	return p.col == other.col && p.row <= other.row
}

func (p Point) Below(other Point) bool {
	return p.col == other.col && p.row >= other.row
}
func (p Point) ToTheLeftOf(other Point) bool {
	return p.row == other.row && p.col <= other.col
}
func (p Point) ToTheRightOf(other Point) bool {
	return p.row == other.row && p.col >= other.col
}
