package matrix

import "github.com/asukakenji/151a48667a3852a43a2028024ffc102e/constant"

type SubMatrix struct {
	parent     Matrix
	omittedRow int
	omittedCol int
}

func (m SubMatrix) Size() (rowSize, colSize int) {
	rowSize, colSize = m.parent.Size()
	if m.omittedRow != -1 {
		rowSize--
	}
	if m.omittedCol != -1 {
		colSize--
	}
	return rowSize, colSize
}

func (m SubMatrix) BaseSize() (rowSize, colSize int) {
	return m.parent.Size()
}

func (m SubMatrix) Get(row, col int) int {
	if row == m.omittedRow {
		return constant.Infinity
	}
	if col == m.omittedCol {
		return constant.Infinity
	}
	return m.parent.Get(row, col)
}

func (m SubMatrix) Set(row, col, value int) {
	if row == m.omittedRow {
		panic("Cannot set omitted row")
	}
	if col == m.omittedCol {
		panic("Cannot set ommited col")
	}
	m.parent.Set(row, col, value)
}

func (m SubMatrix) RowMinimum(row int) int {
	return rowMinimum(m, row)
}

func (m SubMatrix) ColMinimum(col int) int {
	return colMinimum(m, col)
}

func (m SubMatrix) OmitRow(row int) Matrix {
	return omitRow(m, row)
}

func (m SubMatrix) OmitCol(col int) Matrix {
	return omitCol(m, col)
}

func (m SubMatrix) OmitRowCol(row, col int) Matrix {
	return omitRowCol(m, row, col)
}

func (m SubMatrix) SetTemporarily(row, col, value int) Matrix {
	return setTemporarily(m, row, col, value)
}
