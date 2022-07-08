package models

import "testing"

func TestCheckIsMutant(t *testing.T) {
	dnaMutant := []string{"ATGCGA", "CAGTGC", "TTATTT", "AGACGG", "GCGTCA", "TCACTG"}

	t.Run("When dna is human, should return false", func(t *testing.T) {
		var matrix Matrix
		matrix.Create(dnaMutant)

		var sequence Sequence
		sequence.CheckIsMutant(matrix)
		assertEqual(t, sequence.IsMutant, false)
	})

	dnaMutant = []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}

	t.Run("When dna is mutant, should return true", func(t *testing.T) {
		var matrix Matrix
		matrix.Create(dnaMutant)

		var sequence Sequence
		sequence.CheckIsMutant(matrix)
		assertEqual(t, sequence.IsMutant, true)
	})

	dnaMutant = []string{"CTGCTA", "CTGTGC", "TTATGT", "AGAAGG", "TCCCAA", "TCACTA"}

	t.Run("When exist only one sequence, should return false", func(t *testing.T) {
		var matrix Matrix
		matrix.Create(dnaMutant)

		var sequence Sequence
		sequence.CheckIsMutant(matrix)
		assertEqual(t, sequence.IsMutant, false)
	})
}

func TestValidationCountItems(t *testing.T) {
	t.Run("When matrix dimension is greater than 3, should return true", func(t *testing.T) {
		var sequence Sequence
		got := sequence.validationCountItems(4)
		assertEqual(t, got, true)
	})

	t.Run("When dimension is less or equals than 3, should return false", func(t *testing.T) {
		var sequence Sequence
		got := sequence.validationCountItems(3)
		assertEqual(t, got, false)
	})
}

func TestCheckHorizontalSequence(t *testing.T) {
	dnaMutant := []string{"AAAA", "ATCG", "TATG", "TTTT"}

	t.Run("When exist a horizontal sequence of mutant, should return true", func(t *testing.T) {
		var matrix Matrix
		matrix.Create(dnaMutant)

		var sequence Sequence
		sequence.checkHorizontalSequence(matrix)
		assertEqual(t, sequence.IsMutant, true)
	})
}

func TestCheckVerticalSequence(t *testing.T) {
	dnaMutant := []string{"TACG", "TTCG", "TACG", "TACT"}

	t.Run("When exist a vertical sequence of mutant, should return true", func(t *testing.T) {
		var matrix Matrix
		matrix.Create(dnaMutant)

		var sequence Sequence
		sequence.checkVerticalSequence(matrix)
		assertEqual(t, sequence.IsMutant, true)
	})
}

func TestCheckSequenceOblique(t *testing.T) {
	dnaMutant := []string{"TACC", "ATCG", "TCTG", "CACT"}

	t.Run("When exist a oblique sequence of mutant, should return true", func(t *testing.T) {
		var matrix Matrix
		matrix.Create(dnaMutant)

		var sequence Sequence
		sequence.checkObliqueSequence(matrix, false)
		assertEqual(t, sequence.CountSequences, 1)
	})

	dnaMutant = []string{"GACT", "ACTG", "ATCG", "TACG"}

	t.Run("When exist a inverse oblique sequence of mutant, should return true", func(t *testing.T) {
		var matrix Matrix
		matrix.Create(dnaMutant)

		var sequence Sequence
		sequence.checkObliqueSequence(matrix, true)
		assertEqual(t, sequence.CountSequences, 1)
	})
}

func TestDifferentDimensions(t *testing.T) {

	t.Run("Matrix 4x4 with dna mutant, should return true", func(t *testing.T) {
		dnaMutant := []string{"TTTT", "ATCG", "TATG", "TACT"}
		var matrix Matrix
		matrix.Create(dnaMutant)

		var sequence Sequence
		sequence.CheckIsMutant(matrix)
		assertEqual(t, sequence.IsMutant, true)
	})

	t.Run("Matrix 4x4 with dna human, should return false", func(t *testing.T) {
		dnaHuman := []string{"TACG", "ACTG", "TATG", "TACT"}
		var matrix Matrix
		matrix.Create(dnaHuman)

		var sequence Sequence
		sequence.CheckIsMutant(matrix)
		assertEqual(t, sequence.IsMutant, false)
	})

	t.Run("Matrix 5x5 odd with dna mutant, should return true", func(t *testing.T) {
		dnaMutant := []string{"ATGCG", "CAGTG", "TTATG", "AGAAG", "CCCCT"}
		var matrix Matrix
		matrix.Create(dnaMutant)

		var sequence Sequence
		sequence.CheckIsMutant(matrix)
		assertEqual(t, sequence.IsMutant, true)
	})

	t.Run("Matrix 5x5 odd with dna human, should return false", func(t *testing.T) {
		dnaHuman := []string{"ATGCG", "CAGTG", "TTATT", "AGACG", "GCGTC"}
		var matrix Matrix
		matrix.Create(dnaHuman)

		var sequence Sequence
		sequence.CheckIsMutant(matrix)
		assertEqual(t, sequence.IsMutant, false)
	})

	t.Run("Matrix  6x6 with dna mutant, should return true", func(t *testing.T) {
		dnaMutant := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}
		var matrix Matrix
		matrix.Create(dnaMutant)

		var sequence Sequence
		sequence.CheckIsMutant(matrix)
		assertEqual(t, sequence.IsMutant, true)
	})

	t.Run("Matrix 6x6 with dna human, should return false", func(t *testing.T) {
		dnaHuman := []string{"ATGCGA", "CAGTGC", "TTATTT", "AGACGG", "GCGTCA", "TCACTG"}
		var matrix Matrix
		matrix.Create(dnaHuman)

		var sequence Sequence
		sequence.CheckIsMutant(matrix)
		assertEqual(t, sequence.IsMutant, false)
	})

	t.Run("Matrix 10x10 with dna mutant, should return true", func(t *testing.T) {
		dnaMutant := []string{"ATGCGATGCG", "CAGTGCAGTG", "TTATGTTATG", "AGAAGAGAAG", "CCCCTCCCCT", "ATGCGATGCG", "CAGTGCAGTG", "TTATGTTATG", "AGAAGAGAAG", "CCCCTCCCCT"}
		var matrix Matrix
		matrix.Create(dnaMutant)

		var sequence Sequence
		sequence.CheckIsMutant(matrix)
		assertEqual(t, sequence.IsMutant, true)
	})

	t.Run("Matrix 10x10 with dna human, should return false", func(t *testing.T) {
		dnaHuman := []string{"ATGCGATGCG", "CAGTGCAGTG", "TTATTGTATT", "AGACGAGACG", "GCGTCGCGTC", "ATGCGATGCG", "CAGTGCAGTG", "TTATTGTATT", "AGACGATACG", "GCGTCGCGTC"}
		var matrix Matrix
		matrix.Create(dnaHuman)

		var sequence Sequence
		sequence.CheckIsMutant(matrix)
		assertEqual(t, sequence.IsMutant, false)
	})
}
