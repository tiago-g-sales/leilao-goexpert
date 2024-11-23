package validation

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	validator_en "github.com/go-playground/validator/v10/translations/en"
	"github.com/tiago-g-sales/leilao-goexpert/configuration/rest_err"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if value, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		enTransl := ut.New(en, en)
		transl, _ = enTransl.GetTranslator("en")
		validator_en.RegisterDefaultTranslations(value, transl)
	}
}

func ValidatEr(validator_err error) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidator validator.ValidationErrors

	if errors.As(validator_err, &jsonErr) {
		return rest_err.NewNotFoundError("Invalid type error")
	} else if errors.As(validator_err, &jsonValidator) {
		errorsCauses := []rest_err.Causes{}

		for _, e := range validator_err.(validator.ValidationErrors) {
			errorsCauses = append(errorsCauses, rest_err.Causes{
				Field:   e.Field(),
				Message: e.Translate(transl),
			})
		}

		return rest_err.NewBadRequestError("Invalid field values", errorsCauses...)
	} else {
		return rest_err.NewBadRequestError("Invalid field values")
	} 

}
