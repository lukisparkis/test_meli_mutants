package mutants

import (
	"github.com/gofiber/fiber/v2"
	"meli_test/internal/api/validators"
	"meli_test/internal/pkg/services/mutants"
)

// Check godoc
// @Summary Check if a dna is from a mutant.
// @Description Check a given dna.
// @ID CheckMutant
// @Accept  json
// @Produce  json
// @Tags Mutants
// @Param data body validators.DnaValidator true "Check a dna"
// @Success 200 {object} map[string]string "Message successfully!"
// @Failure 403 {string} error
// @Failure 500 {string} error
// @Router /mutants [post]
func Check(c *fiber.Ctx) error {
	var dnaValidator validators.DnaValidator
	if err := c.BodyParser(&dnaValidator); err != nil {
		data := validators.Response{Message: "Error trying to parse body data. Error: " + err.Error()}
		return c.Status(fiber.StatusBadRequest).JSON(data)
	}
	if err := dnaValidator.Validate(); err != nil {
		data := validators.Response{Message: "Error validating data. Error: " + err.Error()}
		return c.Status(fiber.StatusUnprocessableEntity).JSON(data)
	}
	isMutant, err := mutants.Check(&dnaValidator)
	if err != nil {
		data := validators.Response{Message: "Error processing request. Error: " + err.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(data)
	}
	if isMutant {
		data := validators.Response{Message: "The given DNA belongs to a Mutant."}
		return c.Status(fiber.StatusOK).JSON(data)
	}
	data := validators.Response{Message: "The given DNA belongs to a Human."}
	return c.Status(fiber.StatusOK).JSON(data)
}
