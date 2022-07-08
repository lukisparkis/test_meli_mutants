package validators

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"regexp"
)

// DnaValidator this type allow to define the request DNA.
type DnaValidator struct {
	Dna []string `json:"dna"`
}

// Response this type allow to define the response.
type Response struct {
	Message string `json:"message"`
}

// Validate if request is correct
func (m DnaValidator) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Dna, validation.Required.Error("Field is required")),
		validation.Field(&m.Dna, validation.Each(is.UpperCase, is.Alpha, validation.Match(regexp.MustCompile("^[CAGT]+$")).Error("It is only allowed to use the words [ACGT]"))),
	)
}
