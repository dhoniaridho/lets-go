package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ValidateBody[T any](validate *validator.Validate, data T) (Response, error) {
	validationErrors := []ErrorResponse{}

	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			// In this case data object is actually holding the User struct
			var elem ErrorResponse

			elem.FailedField = err.Field() // Export struct field name
			elem.Tag = err.Tag()           // Export struct tag
			elem.Value = err.Value()       // Export field value
			elem.Error = true

			validationErrors = append(validationErrors, elem)
		}

		return Response{
			Data:    nil,
			Message: errs.Error(),
			Status:  fiber.StatusBadRequest,
			Errors:  validationErrors,
		}, errs
	}

	return Response{
		Data:    data,
		Message: "success",
		Status:  fiber.StatusOK,
		Errors:  validationErrors,
	}, nil
}
