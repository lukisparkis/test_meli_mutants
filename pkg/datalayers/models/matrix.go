package models

import (
	"strings"
)

// Matrix struct
type Matrix struct {
	Data      [][]string
	Dimension int
}

// Create a matrix with a string array given
func (m *Matrix) Create(arrays []string) {
	if len(arrays) <= 0 {
		return
	}
	m.setDimension(arrays)
	m.fillMatrix(arrays)
}

func (m *Matrix) setDimension(arrays []string) {
	characters := strings.SplitAfter(arrays[0], "")
	m.Dimension = len(characters)
}

func (m *Matrix) fillMatrix(arrays []string) {
	matrix := make([][]string, 0)

	for _, item := range arrays {
		letters := strings.SplitAfter(item, "")
		matrix = append(matrix, letters)
	}

	m.Data = matrix
}
