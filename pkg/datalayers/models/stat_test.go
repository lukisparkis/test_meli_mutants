package models

import "testing"

func TestCalculateRatio(t *testing.T) {
	t.Run("When CountHumanDna is equal than 0, should return 0", func(t *testing.T) {
		var stat Stat
		stat.CountHumanDna = 0
		stat.calculateRatio()
		assertEqual(t, stat.Ratio, float64(0))
	})

	t.Run("When CountHumanDna is not equals than 0, should return CountMutantDna / CountHumanDna", func(t *testing.T) {
		var stat Stat
		stat.CountMutantDna = 40
		stat.CountHumanDna = 100
		stat.calculateRatio()
		assertResult(t, stat.Ratio, 0.4)
	})
}

func TestCountDna(t *testing.T) {
	t.Run("When array is empty, should return 0", func(t *testing.T) {
		var mutant []Mutant
		var stat Stat
		stat.countDna(mutant)
		assertEqual(t, stat.CountMutantDna, 0)
		assertEqual(t, stat.CountHumanDna, 0)
	})

	t.Run("When array have 3 mutants, should return 3", func(t *testing.T) {
		mutant := []Mutant{
			{IsMutant: true},
			{IsMutant: true},
			{IsMutant: true},
		}
		var stat Stat
		stat.countDna(mutant)
		assertEqual(t, stat.CountMutantDna, 3)
		assertEqual(t, stat.CountHumanDna, 0)
	})

	t.Run("When array have 3 humans, should return 3", func(t *testing.T) {
		mutant := []Mutant{
			{IsMutant: false},
			{IsMutant: false},
			{IsMutant: false},
		}
		var stat Stat
		stat.countDna(mutant)
		assertEqual(t, stat.CountMutantDna, 0)
		assertEqual(t, stat.CountHumanDna, 3)
	})

	t.Run("When array have 3 humans, should return 3", func(t *testing.T) {
		mutant := []Mutant{
			{IsMutant: false},
			{IsMutant: false},
			{IsMutant: false},
		}
		var stat Stat
		stat.countDna(mutant)
		assertEqual(t, stat.CountMutantDna, 0)
		assertEqual(t, stat.CountHumanDna, 3)
	})

	t.Run("When array have 3 humans and 2 mutants, should return 3 and 2", func(t *testing.T) {
		mutant := []Mutant{
			{IsMutant: false},
			{IsMutant: false},
			{IsMutant: false},
			{IsMutant: true},
			{IsMutant: true},
		}
		var stat Stat
		stat.countDna(mutant)
		assertEqual(t, stat.CountHumanDna, 3)
		assertEqual(t, stat.CountMutantDna, 2)
	})
}
