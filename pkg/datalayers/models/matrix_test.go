package models

import (
	"testing"
)

var arrays = []string{
	"ABCD",
	"EFGH",
	"IJKL",
	"WXYZ",
}

func TestCreateMatrix(t *testing.T) {
	t.Run("Create Matrix with a array given, should return a matrix valid", func(t *testing.T) {
		var matrix Matrix
		matrix.Create(arrays)
		want := [][]string{{"A", "B", "C", "D"}, {"E", "F", "G", "H"}, {"I", "J", "K", "L"}, {"W", "X", "Y", "Z"}}
		assertEqual(t, matrix.Dimension, 4)
		assertResult(t, matrix.Data, want)
	})
}

func TestSetDimension(t *testing.T) {
	t.Run("Calculate dimension of a Matrix of 4 x 4, should return 4", func(t *testing.T) {
		var matrix Matrix
		matrix.setDimension(arrays)
		assertEqual(t, matrix.Dimension, 4)
	})
}

func TestFillMatrix(t *testing.T) {
	t.Run("Fill a Matrix with an array given, should be a matrix valid", func(t *testing.T) {
		var matrix Matrix
		matrix.fillMatrix(arrays)
		want := [][]string{{"A", "B", "C", "D"}, {"E", "F", "G", "H"}, {"I", "J", "K", "L"}, {"W", "X", "Y", "Z"}}
		assertResult(t, matrix.Data, want)
	})
}
