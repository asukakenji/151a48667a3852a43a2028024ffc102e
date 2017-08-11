package matrix

type OverrideMatrix struct {
	parent          Matrix
	overriddenRow   int
	overriddenCol   int
	overriddenValue int
}

func (m OverrideMatrix) Size() (rowSize, colSize int) {
	return m.parent.Size()
}

func (m OverrideMatrix) BaseSize() (rowSize, colSize int) {
	return m.parent.Size()
}

func (m OverrideMatrix) Get(row, col int) int {
	if row == m.overriddenRow && col == m.overriddenCol {
		return m.overriddenValue
	}
	return m.parent.Get(row, col)
}

func (m OverrideMatrix) Set(row, col, value int) {
	if row == m.overriddenRow && col == m.overriddenCol {
		m.overriddenValue = value
	}
	m.parent.Set(row, col, value)
}

func (m OverrideMatrix) RowMinimum(row int) int {
	return rowMinimum(m, row)
}

func (m OverrideMatrix) ColMinimum(col int) int {
	return colMinimum(m, col)
}

func (m OverrideMatrix) OmitRow(row int) Matrix {
	return omitRow(m, row)
}

func (m OverrideMatrix) OmitCol(col int) Matrix {
	return omitCol(m, col)
}

func (m OverrideMatrix) OmitRowCol(row, col int) Matrix {
	return omitRowCol(m, row, col)
}

func (m OverrideMatrix) SetTemporarily(row, col, value int) Matrix {
	return setTemporarily(m, row, col, value)
}
