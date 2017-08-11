package matrix

import "github.com/asukakenji/151a48667a3852a43a2028024ffc102e/constant"

type Matrix interface {
	Size() (rowSize, colSize int)
	BaseSize() (rowSize, colSize int)
	Get(row, col int) int
	Set(row, col, value int)
	RowMinimum(row int) int
	ColMinimum(col int) int
	OmitRow(row int) Matrix
	OmitCol(col int) Matrix
	OmitRowCol(row, col int) Matrix
	SetTemporarily(row, col, value int) Matrix
}

func rowMinimum(m Matrix, row int) int {
	min := constant.MaxInt
	_, colSize := m.BaseSize()
	for col := 0; col < colSize; col++ {
		if v := m.Get(row, col); v != constant.Infinity && v < min {
			min = v
		}
	}
	return min
}

func colMinimum(m Matrix, col int) int {
	min := constant.MaxInt
	rowSize, _ := m.BaseSize()
	for row := 0; row < rowSize; row++ {
		if v := m.Get(row, col); v != constant.Infinity && v < min {
			min = v
		}
	}
	return min
}

func omitRow(m Matrix, row int) Matrix {
	return SubMatrix{
		parent:     m,
		omittedRow: row,
		omittedCol: constant.Infinity,
	}
}

func omitCol(m Matrix, col int) Matrix {
	return SubMatrix{
		parent:     m,
		omittedRow: constant.Infinity,
		omittedCol: col,
	}
}

func omitRowCol(m Matrix, row, col int) Matrix {
	return SubMatrix{
		parent:     m,
		omittedRow: row,
		omittedCol: col,
	}
}

func setTemporarily(m Matrix, row, col, value int) Matrix {
	return OverrideMatrix{
		parent:          m,
		overriddenRow:   row,
		overriddenCol:   col,
		overriddenValue: value,
	}
}
