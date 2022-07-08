package models

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	setConfigDatabase()
	os.Exit(m.Run())
}

func TestIsMutant(t *testing.T) {
	t.Run("When the dna is empty, should return false", func(t *testing.T) {
		var dna []string
		got := IsMutant(dna)
		assertResult(t, got, false)
	})

	t.Run("When the dimension dna matrix is less or equal than 3, should return false", func(t *testing.T) {
		dna := []string{"AAA", "TTT", "CCC"}
		got := IsMutant(dna)
		assertResult(t, got, false)
	})

	t.Run("When the array dna is equal than 4 and the same letter, should return true", func(t *testing.T) {
		dna := []string{"AAAA", "TTTT", "CCCC", "GGGG"}
		got := IsMutant(dna)
		assertResult(t, got, true)
	})

	t.Run("When exist a Horizontal sequence, should return true", func(t *testing.T) {
		dna := []string{"CTGCCA", "CAGTGC", "TTTTGG", "AGAAGG", "CCCCTA", "TCACTG"}
		got := IsMutant(dna)
		assertResult(t, got, true)
	})

	t.Run("When exist a Vertical sequence, should return true", func(t *testing.T) {
		dna := []string{"CTGCGA", "CAGTGC", "TTGTGT", "AGGAGG", "GCCCTA", "TCACTG"}
		got := IsMutant(dna)
		assertResult(t, got, true)
	})

	t.Run("When exist a Oblique sequence, should return true", func(t *testing.T) {
		dna := []string{"ATGCCA", "CAGTGC", "TTATGG", "AGAAGG", "GCCGTA", "TCGCTG"}
		got := IsMutant(dna)
		assertResult(t, got, true)
	})

	t.Run("When send a human dna, should return false", func(t *testing.T) {
		dna := []string{"ATGCGA", "CAGTGC", "TTATTT", "AGACGG", "GCGTCA", "TCACTG"}
		got := IsMutant(dna)
		assertResult(t, got, false)
	})
}
