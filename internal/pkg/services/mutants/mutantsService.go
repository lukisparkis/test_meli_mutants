package mutants

import (
	"meli_test/internal/api/validators"
	"meli_test/pkg/datalayers/models"
)

// Check Services
func Check(dnaValidator *validators.DnaValidator) (bool, error) {
	isMutant := models.IsMutant(dnaValidator.Dna)
	return isMutant, nil
}
