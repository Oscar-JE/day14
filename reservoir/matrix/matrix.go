package matrix

type Matrix struct {
	array [][]rune
}

func Init(nrRows int, nrCols int) Matrix {
	arr := [][]rune{}
	for i := 0; i < nrRows; i++ {
		row := []rune{}
		for j := 0; j < nrCols; j++ {
			row = append(row, '.')
		}
		arr = append(arr, row)
	}
	return Matrix{array: arr}
}

func (m *Matrix) Set(row int, col int, r rune) {
	m.array[row][col] = r
}

func (m *Matrix) Get(row int, col int) rune {
	return m.array[row][col]
}
