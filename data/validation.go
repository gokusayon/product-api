package dataimport

import (
	"fmt"
	"github.com/go-playground/validator"
	"regexp"
)
// ValidationError wraps the validators FieldError so we do not
// expose this to out code
type ValidationError struct {
	validator.FieldError
}


func (v *ValidationError) Error() string{
	return fmt.Sprintf(
		"Key: '%s' Error: Feild validation for '%s' failed on '%s' tag",
		v.Namespace(),
		v.Field(),
		v.Tag())
}

// ValidationErrors is a collection of ValidationError
type ValidationErrors []ValidationError

// Errors converts the slice into a string slice
func (v ValidationErrors) Errors() []string {
	errs := []string{}
	for _, err := range v {
		errs = append(errs, err.Error())
	}

	return errs
}


// Validation contains
type Validation struct {
	validate *validator.Validate
}

func NewValidation() *Validation{
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return &Validation{validate}
}

func (v *Validation) Validate(i interface{}) ValidationErrors {
	valErr := v.validate.Struct(i)

	if valErr == nil {
		return nil
	}

	errs := valErr.(validator.ValidationErrors)

	var returnErrs []ValidationError
	for _, err := range errs {
		// cast the FieldError into our ValidationError and append to the slice
		ve := ValidationError{err.(validator.FieldError)}
		returnErrs = append(returnErrs, ve)
	}

	return returnErrs
}

// fomat : abc-asdf-asdfs
func validateSKU(f validator.FieldLevel) bool{

	re := regexp.MustCompile(`[a-z+]-[a-z]+-[a-z]`)
	matches := re.FindAllString(f.Field().String(), -1)

	if len(matches) != 1{
		return false
	}

	return true
}
