package matrix

type SquareMatrix [][]int

func NewSquareMatrix(m [][]int) Matrix {
	if !isSquareMatrix(m) {
		panic("m is not a square matrix")
	}
	return SquareMatrix(m)
}

func (m SquareMatrix) Size() (rowSize, colSize int) {
	return len(m), len(m)
}

func (m SquareMatrix) BaseSize() (rowSize, colSize int) {
	return len(m), len(m)
}

func (m SquareMatrix) Get(row, col int) int {
	return m[row][col]
}

func (m SquareMatrix) Set(row, col, value int) {
	m[row][col] = value
}

func (m SquareMatrix) RowMinimum(row int) int {
	return rowMinimum(m, row)
}

func (m SquareMatrix) ColMinimum(col int) int {
	return colMinimum(m, col)
}

func (m SquareMatrix) OmitRow(row int) Matrix {
	return omitRow(m, row)
}

func (m SquareMatrix) OmitCol(col int) Matrix {
	return omitCol(m, col)
}

func (m SquareMatrix) OmitRowCol(row, col int) Matrix {
	return omitRowCol(m, row, col)
}

func (m SquareMatrix) SetTemporarily(row, col, value int) Matrix {
	return setTemporarily(m, row, col, value)
}

func isSquareMatrix(matrix [][]int) bool {
	size := len(matrix)
	for _, row := range matrix {
		if len(row) != size {
			return false
		}
	}
	return true
}
