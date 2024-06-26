package request

import (
	"github.com/go-playground/validator/v10"
)

type RequestPost struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}

func (input *RequestPost) Validate() error {
	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		return err
	}
	return nil
}
