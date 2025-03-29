package matrix

import "errors"

var ErrInvalidDimensions = errors.New("invalid matrix dimensions")

type Matrix struct {
	s [][]int
}

func New(s [][]int) (*Matrix, error) {
	for _, row := range s {
		if len(row) != len(s[0]) {
			return nil, ErrInvalidDimensions
		}
	}
	if len(s) == 0 {
		return nil, ErrInvalidDimensions
	}
	return &Matrix{s}, nil
}

func (m *Matrix) RowsLen() int {
	return len(m.s)
}

func (m *Matrix) ColsLen() int {
	return len(m.s[0])
}

func (m *Matrix) Add(m2 *Matrix) (*Matrix, error) {
	if m.RowsLen() != m2.RowsLen() || m.ColsLen() != m2.ColsLen() {
		return nil, ErrInvalidDimensions
	}
	added := make([][]int, len(m.s))
	for i := range added {
		added[i] = make([]int, len(m.s[i]))
		for j := range added[i] {
			added[i][j] = m.s[i][j] + m2.s[i][j]
		}
	}
	return New(added)
}

func Multiply(m1, m2 *Matrix) (*Matrix, error) {
	if m1.ColsLen() != m2.RowsLen() {
		return nil, ErrInvalidDimensions
	}
	multiplied := make([][]int, m1.RowsLen())
	for i := range multiplied {
		multiplied[i] = make([]int, m2.ColsLen())
		for j := range multiplied[i] {
			for k := range m1.s[i] {
				multiplied[i][j] += m1.s[i][k] * m2.s[k][j]
			}
		}
	}
	return New(multiplied)
}
