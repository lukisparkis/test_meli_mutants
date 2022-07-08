package models

const allowedDimension = 3
const limitToIgnoreInOblique = 3
const minimumSequence = 1

// Sequence struct
type Sequence struct {
	IsMutant       bool
	CountSequences int
}

func (s *Sequence) setIsMutant(value bool) {
	if value {
		s.setCountSequence()
	}
	if value && s.CountSequences > minimumSequence {
		s.IsMutant = true
	}
}

func (s *Sequence) setCountSequence() {
	s.CountSequences++
}

// CheckIsMutant validate if a matrix has mutant dna.
func (s *Sequence) CheckIsMutant(matrix Matrix) {
	isValid := s.validationCountItems(matrix.Dimension)
	if !isValid {
		return
	}

	s.checkHorizontalSequence(matrix)
	if s.IsMutant {
		return
	}

	s.checkVerticalSequence(matrix)
	if s.IsMutant {
		return
	}

	s.checkObliqueSequence(matrix, false)
	if s.IsMutant {
		return
	}

	s.checkObliqueSequence(matrix, true)
	if s.IsMutant {
		return
	}
}

func (s *Sequence) validationCountItems(dimension int) bool {
	return !(dimension <= allowedDimension)
}

func (s *Sequence) checkHorizontalSequence(matrix Matrix) {
	for i := 0; i < matrix.Dimension; i++ {
		s.checkSequence(matrix.Data[i])
		if s.IsMutant {
			return
		}
	}
}

func (s *Sequence) checkVerticalSequence(matrix Matrix) {
	for col := 0; col < matrix.Dimension; col++ {
		sequence := make([]string, 0)
		for row := 0; row < matrix.Dimension; row++ {
			sequence = append(sequence, matrix.Data[row][col])
		}
		s.checkSequence(sequence)
		if s.IsMutant {
			return
		}
	}
}

func (s *Sequence) checkObliqueSequence(matrix Matrix, isInverse bool) {
	isReturn := false
	y := 0
	x := getStartIndex(matrix.Dimension, isInverse)
	c := getCycles(matrix.Dimension)

	for i := 0; i < c; i++ {
		sequence := make([]string, 0)
		count := 0

		if isInverse {
			for row := x; row >= 0 && y == 0; row-- {
				sequence = append(sequence, matrix.Data[row][count])
				count++
			}
		} else {
			for row := x; row <= (matrix.Dimension-1) && y == 0; row++ {
				sequence = append(sequence, matrix.Data[row][count])
				count++
			}
		}

		if isInverse {
			count = matrix.Dimension - 1
			for col := y; col <= (matrix.Dimension-1) && y >= 1; col++ {
				sequence = append(sequence, matrix.Data[count][col])
				count--
			}
		} else {
			count = 0
			for col := y; col <= (matrix.Dimension-1) && y >= 1; col++ {
				sequence = append(sequence, matrix.Data[count][col])
				count++
			}
		}

		if isInverse {
			if x == matrix.Dimension-1 {
				isReturn = true
			}
		} else {
			if x == y {
				isReturn = true
			}
		}

		if isInverse {
			if x == 0 {
				x = matrix.Dimension - 1
			} else {
				x++
			}

			if y != 0 {
				x = matrix.Dimension - 1
			}
		} else {
			if x == 0 {
				x = 0
			} else {
				x--
			}

			if y == 0 {
				y = 0
			}
		}

		if isReturn {
			y++
		}

		s.checkSequence(sequence)
		if s.IsMutant {
			return
		}
	}
}

func (s *Sequence) checkSequence(sequence []string) bool {
	var lastLetter string
	count := 0
	attempts := 4
	max := len(sequence)

	for i, sq := range sequence {
		if lastLetter == "" || lastLetter == sq {
			count++
		} else {
			if (max - i) < attempts {
				return false
			}
			count = 1
		}
		lastLetter = sq

		if count == 4 {
			s.setIsMutant(true)
			if s.IsMutant {
				return true
			}
		}
	}

	return false
}

func getCycles(dimension int) int {
	return ((dimension - limitToIgnoreInOblique) * 2) - 1
}

func getStartIndex(dimension int, isInverse bool) int {
	if isInverse {
		return limitToIgnoreInOblique
	}
	return (dimension - limitToIgnoreInOblique) - 1
}
