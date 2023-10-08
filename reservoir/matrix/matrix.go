package matrix

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

func (m *Matrix) Get(row int, col int) string {
	return m.array[row][col]
}
