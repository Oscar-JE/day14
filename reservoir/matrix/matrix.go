package matrix

import "fmt"

type Matrix struct {
	array [][]string
}

func Init(nrRows int, nrCols int) Matrix {
	arr := [][]string{}
	for i := 0; i < nrRows; i++ {
		row := []string{}
		for j := 0; j < nrCols; j++ {
			row = append(row, ".")
		}
		arr = append(arr, row)
	}
	return Matrix{array: arr}
}

func (m *Matrix) Set(row int, col int, r string) {
	m.array[row][col] = r
}

func (m *Matrix) SetWall(row int, col int) {
	m.array[row][col] = "#"
}

func (m *Matrix) SetSand(row int, col int) {
	m.array[row][col] = "o"
}

func (m *Matrix) Get(row int, col int) string {
	return m.array[row][col]
}

func (m Matrix) String() string {
	retString := ""
	for index := range m.array {
		retString += fmt.Sprintln(m.array[index])
	}
	return retString
}

func (m Matrix) InMatrix(row int, col int) bool {
	inRow := 0 <= row && row < len(m.array)
	inCol := 0 <= col && col < len(m.array[0])
	return inRow && inCol
}
